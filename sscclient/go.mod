module github.com/adiclepcea/socialservicechain/sscclient

go 1.12

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2 v0.0.0 => ../../../hyperledger/sawtooth-sdk-go/protobuf/batch_pb2

replace github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0 => ../../../hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2

require (
	github.com/brianolson/cbor_go v1.0.0
	github.com/btcsuite/btcd v0.0.0-20190315201642-aa6e0f35703c // indirect
	github.com/golang/protobuf v1.3.1
	github.com/hyperledger/sawtooth-sdk-go v0.1.1
	github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2 v0.0.0
	github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2 v0.0.0
	gopkg.in/yaml.v2 v2.2.2
)
