package sscstate

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"

	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

//NGO represents an Non Governmental Organization
type NGO struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

//A Need is a need a SocialCase has or one that a Donation Fulfills
type Need struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	UM     string  `json:"um"`
	ID     int64   `json:"id"`
}

//A SocialCase represents one social case that has needs
type SocialCase struct {
	Name  string `json:"name"`
	Needs []Need `json:"needs"`
	ID    int64  `json:"id"`
}

//A Donor is the source of a donation
type Donor struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

//A Donation will satisfy a need of a SocialCase when it will be assigned
type Donation struct {
	Name         string `json:"name"`
	ID           int64  `json:"id"`
	Needs        []Need `json:"needs"`
	DonorID      int64  `json:"donor_id"`
	SocialCaseID int64  `json:"social_case_id"`
}

//A DonationAssignment will transfer fulfill the SocialCase need
type DonationAssignment struct {
	DonorID      int64  `json:"donor_id"`
	Needs        []Need `json:"needs"`
	SocialCaseID int64  `json:"social_case_id"`
}

//Namespace will hold the sawtooth namespace
var Namespace = hexdigest("ssc")[:6]

//SSCState represents the state for this app (Social Service Chain)
//This will be used to store and read the state
type SSCState struct {
	Context *processor.Context
	storage StateStorage
}

//NewSSCState creates a new state from the given context
func NewSSCState(context *processor.Context) *SSCState {
	return &SSCState{
		Context: context,
		storage: &MemoryStateStorage{},
	}
}

//GetNGO will return the NGO as it is stored into the storage
func (sscstate *SSCState) GetNGO(name string) (*NGO, error) {
	return sscstate.storage.GetNGO(name)
}

//SaveNGO will store the ngo into the storage
func (sscstate *SSCState) SaveNGO(ngo *NGO) error {
	return sscstate.storage.SaveNGO(ngo)
}

func hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}
