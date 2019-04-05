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
var mapActions = map[string]func(*sscstate.SSCState) error{
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

func (sa *SocialAggregator) FamilyName() string {
	return "SocialAggregator"
}

func (sa *SocialAggregator) FamilyVersions() []string {
	return []string{"1.0"}
}

func (sa *SocialAggregator) Namespaces() []string {
	return []string{sscstate.Namespace}
}

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
		fnc(state)
	}

	return nil
}

func createNGO(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func createSocialCase(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func createDonor(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteNGO(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteSocialCase(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func deleteDonor(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func donation(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}

func help(state *sscstate.SSCState) error {
	return fmt.Errorf("Not yet implemented")
}
