module github.com/adiclepcea/socialservicechain/cmd/ssc

go 1.12i

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/events_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/events_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/network_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/network_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/processor_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/state_context_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/state_context_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/validator_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/validator_pb2

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/hyperledger/sawtooth-sdk-go v0.1.1
	github.com/hyperledger/sawtooth-sdk-go/protobuf/events_pb2 v0.0.0 // indirect
	github.com/hyperledger/sawtooth-sdk-go/protobuf/network_pb2 v0.0.0 // indirect
	github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2 v0.0.0
	github.com/hyperledger/sawtooth-sdk-go/protobuf/state_context_pb2 v0.0.0 // indirect
	github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 // indirect
	github.com/hyperledger/sawtooth-sdk-go/protobuf/validator_pb2 v0.0.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/pebbe/zmq4 v1.0.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
)
