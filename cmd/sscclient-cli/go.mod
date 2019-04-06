module github.com/adiclepcea/socialservicechain/cmd/sscclient-cli

go 1.12

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/batch_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/processor_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/validator_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/validator_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/events_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/events_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/network_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/network_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/state_context_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/state_context_pb2

replace github.com/adiclepcea/socialservicechain/sscclient v0.0.0 => ../../sscclient

require (
	github.com/adiclepcea/socialservicechain/cmd/ssc v0.0.0-20190406145352-cfd7275f9005
	github.com/adiclepcea/socialservicechain/sscclient v0.0.0
	github.com/hyperledger/sawtooth-sdk-go v0.1.1
	github.com/jessevdk/go-flags v1.4.0
)
