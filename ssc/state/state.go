package sscstate

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"

	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

//Namespace will hold the sawtooth namespace
var Namespace = hexdigest("ssc")[:6]

//SSCState represents the state for this app (Social Service Chain)
//This will be used to store and read the state
type SSCState struct {
	Context *processor.Context
}

//NewSSCState creates a new state from the given context
func NewSSCState(context *processor.Context) *SSCState {
	return &SSCState{
		Context: context,
	}
}

func hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}
