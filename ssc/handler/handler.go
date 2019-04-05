package handler

import (
	"fmt"

	sscpayload "github.com/adiclepcea/socialservicechain/ssc/payload"
	sscstate "github.com/adiclepcea/socialservicechain/ssc/state"

	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2"
)

var logger = logging.Get()
var mapActions = map[string]func(sscpayload.SSCPayload, *sscstate.SSCState) error{
	"createNGO":        createNGO,
	"createSocialCase": createSocialCase,
	"createDonor":      createDonor,
	"deleteNGO":        deleteNGO,
	"deleteSocialCase": deleteSocialCase,
	"deleteDonor":      deleteDonor,
	"donation":         donation,
	"help":             help,
}

//SocialAggregator is used to process social NGOs
type SocialAggregator struct {
}

//FamilyName returns our transaction family name
func (sa *SocialAggregator) FamilyName() string {
	return "SocialServiceChain"
}

//FamilyVersions returns the versions of our transaction
func (sa *SocialAggregator) FamilyVersions() []string {
	return []string{"1.0"}
}

//Namespaces returns the namespaces for our states
func (sa *SocialAggregator) Namespaces() []string {
	return []string{sscstate.Namespace}
}

//Apply will be called when we apply the received commands
func (sa *SocialAggregator) Apply(request *processor_pb2.TpProcessRequest, context *processor.Context) error {
	header := request.GetHeader()
	player := header.GetSignerPublicKey()

	payload, err := sscpayload.FromBytes(request.GetPayload())
	if err != nil {
		return err
	}

	state := sscstate.NewSSCState(context)

	logger.Debugf("xo txn %v: player %v: payload: Name='%v', Action='%v', Space='%v'",
		request.Signature, player, payload.Name, payload.Action, payload.Space)

	fnc := mapActions[payload.Action]
	if fnc != nil {
		fnc(payload, state)
	}

	return nil
}

func createNGO(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {

	if err := validateNGO(payload, state); err != nil {
		return err
	}

	ngo := sscstate.NGO{Name: payload.Name}
	if err := state.SaveNGO(&ngo); err != nil {
		return err
	}

	displayNGO(&ngo)

	return nil
}

func validateNGO(payload sscpayload.SSCPayload, sscstate *sscstate.SSCState) error {

	ngo, err := sscstate.GetNGO(payload.Name)
	if err != nil {
		return err
	}
	if ngo != nil {
		return &processor.InvalidTransactionError{Msg: "An NGO with this name already exists"}
	}

	return nil
}

func displayNGO(ngo *sscstate.NGO) {
	fmt.Printf("NGO ID: %d, name: %s\n", ngo.ID, ngo.Name)
}

func createSocialCase(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func createDonor(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteNGO(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteSocialCase(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteDonor(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func donation(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func help(payload sscpayload.SSCPayload, state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}
