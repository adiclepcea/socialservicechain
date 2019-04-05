package handler

import (
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2"
)

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
	return []string{"???????"}
}

func (sa *SocialAggregator) Apply(request *processor_pb2.TpProcessRequest, context *processor.Context) error {
	return nil
}
