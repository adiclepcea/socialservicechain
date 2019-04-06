package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	cbor "github.com/brianolson/cbor_go"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/signing"
	"gopkg.in/yaml.v2"
)

//SSCClientCli is the basic struct for calling a SocialServiceChain transaction on a validator
type SSCClientCli struct {
	url    string
	signer *signing.Signer
}

//NewSSCClientCli creates a SSCClientCli with the given
//url and key
func NewSSCClientCli(url string, keyfile string) (*SSCClientCli, error) {

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return nil, fmt.Errorf("Invalid url provided: %s, not a valid http url", url)
	}

	var privateKey signing.PrivateKey
	if keyfile != "" {
		privateKeyStr, err := ioutil.ReadFile(keyfile)
		if err != nil {
			return nil, fmt.Errorf("Failed to read private key: %s", err.Error())
		}
		privateKey = signing.NewSecp256k1PrivateKey(privateKeyStr)
	} else {
		privateKey = signing.NewSecp256k1Context().NewRandomPrivateKey()
	}
	cryptoFactory := signing.NewCryptoFactory(signing.NewSecp256k1Context())
	signer := cryptoFactory.NewSigner(privateKey)

	return &SSCClientCli{url, signer}, nil
}

func (sscClient *SSCClientCli) getStatus(batchID string, wait uint) (string, error) {
	path := fmt.Sprintf("%s?id=%s&wait=%d", BatchStatusAPI, batchID, wait)
	response, err := sscClient.getRequest(path)
	if err != nil {
		return "", err
	}

	responseMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal(response, &responseMap)
	if err != nil {
		return "", fmt.Errorf("Error reading response: %s", err.Error())
	}
	entry := responseMap["data"].([]interface{})[0].(map[interface{}]interface{})
	return fmt.Sprint(entry["status"]), nil
}

func (sscClient *SSCClientCli) getRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", sscClient.url, path)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to perform the GET request: %s\n", err.Error())
		return nil, err
	}

	return readResponse(response)
}

func (sscClient *SSCClientCli) postRequest(path string, data []byte, contentType string, name string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", sscClient.url, path)
	response, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Failed to perform the POST request: %s\n", err.Error())
		return nil, err
	}

	return readResponse(response)
}

func readResponse(response *http.Response) ([]byte, error) {

	defer func() {
		ioutil.ReadAll(response.Body)
		response.Body.Close()
	}()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %s", err.Error())
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Error %d: %s, %s", response.StatusCode, response.Status, responseBody)
	}

	return responseBody, nil
}

func (sscClient *SSCClientCli) sendTransaction(verb string, name string, value uint, wait uint) ([]byte, error) {

	// construct the payload information in CBOR format
	payloadData := make(map[string]interface{})
	payloadData["Verb"] = verb
	payloadData["Name"] = name
	payloadData["Value"] = value
	payload, err := cbor.Dumps(payloadData)
	if err != nil {
		return nil, fmt.Errorf("Failed to construct CBOR: %s", err.Error())
	}

	// construct the address
	address := sscClient.getAddress(name)

	// Construct TransactionHeader
	rawTransactionHeader := transaction_pb2.TransactionHeader{
		SignerPublicKey:  sscClient.signer.GetPublicKey().AsHex(),
		FamilyName:       FamilyName,
		FamilyVersion:    FamilyVersion,
		Dependencies:     []string{}, // empty dependency list
		Nonce:            strconv.Itoa(rand.Int()),
		BatcherPublicKey: sscClient.signer.GetPublicKey().AsHex(),
		Inputs:           []string{address},
		Outputs:          []string{address},
		PayloadSha512:    sha512HashValue(string(payload)),
	}
	transactionHeader, err := proto.Marshal(&rawTransactionHeader)
	if err != nil {
		return nil, fmt.Errorf("Unable to serialize transaction header: %s", err.Error())
	}

	// Signature of TransactionHeader
	transactionHeaderSignature := hex.EncodeToString(sscClient.signer.Sign(transactionHeader))

	// Construct Transaction
	transaction := transaction_pb2.Transaction{
		Header:          transactionHeader,
		HeaderSignature: transactionHeaderSignature,
		Payload:         []byte(payload),
	}

	// Get BatchList
	rawBatchList, err := sscClient.createBatchList([]*transaction_pb2.Transaction{&transaction})
	if err != nil {
		return nil, fmt.Errorf("Unable to construct batch list: %s", err.Error())
	}
	batchID := rawBatchList.Batches[0].HeaderSignature
	batchList, err := proto.Marshal(&rawBatchList)
	if err != nil {
		return nil, fmt.Errorf("Unable to serialize batch list: %s", err.Error())
	}

	if wait > 0 {
		waitTime := uint(0)
		startTime := time.Now()
		response, err := sscClient.postRequest(BatchSubmitAPI, batchList, ContentTypeOctetStream, name)
		if err != nil {
			return nil, err
		}
		for waitTime < wait {
			status, err := sscClient.getStatus(batchID, wait-waitTime)
			if err != nil {
				return nil, err
			}
			waitTime = uint(time.Now().Sub(startTime))
			if status != "PENDING" {
				return response, nil
			}
		}
		return response, nil
	}

	return sscClient.postRequest(BatchSubmitAPI, batchList, ContentTypeOctetStream, name)
}

func (sscClient *SSCClientCli) getPrefix() string {
	return sha512HashValue(FamilyName)[:FamilyNamespaceAddressLen]
}

func (sscClient *SSCClientCli) getAddress(name string) string {
	prefix := sscClient.getPrefix()
	nameAddress := sha512HashValue(name)[FamilyVerbAddressLen:]
	return prefix + nameAddress
}

func (sscClient *SSCClientCli) createBatchList(transactions []*transaction_pb2.Transaction) (batch_pb2.BatchList, error) {

	// Get list of TransactionHeader signatures
	transactionSignatures := []string{}
	for _, transaction := range transactions {
		transactionSignatures = append(transactionSignatures, transaction.HeaderSignature)
	}

	// Construct BatchHeader
	rawBatchHeader := batch_pb2.BatchHeader{
		SignerPublicKey: sscClient.signer.GetPublicKey().AsHex(),
		TransactionIds:  transactionSignatures,
	}
	batchHeader, err := proto.Marshal(&rawBatchHeader)
	if err != nil {
		return batch_pb2.BatchList{}, fmt.Errorf("Unable to serialize batch header: %v", err)
	}

	// Signature of BatchHeader
	batchHeaderSignature := hex.EncodeToString(sscClient.signer.Sign(batchHeader))

	// Construct Batch
	batch := batch_pb2.Batch{
		Header:          batchHeader,
		Transactions:    transactions,
		HeaderSignature: batchHeaderSignature,
	}

	// Construct BatchList
	return batch_pb2.BatchList{
		Batches: []*batch_pb2.Batch{&batch},
	}, nil
}
