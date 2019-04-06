module github.com/adiclepcea/socialservicechain/cmd/sscclient-cli

go 1.12

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/batch_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 => ../../../../hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2

replace github.com/adiclepcea/socialservicechain/sscclient v0.0.0 => ../../sscclient

require (
	github.com/adiclepcea/socialservicechain/sscclient v0.0.0
	github.com/hyperledger/sawtooth-sdk-go v0.1.1
	github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2 v0.0.0 // indirect
	github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 // indirect
	github.com/jessevdk/go-flags v0.0.0-20141203071132-1679536dcc89
)
