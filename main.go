package main

import "github.com/hyperledger/sawtooth-sdk-go/processor"

const (
	endpoint := "tcp://127.0.0.1:4004"
)

func main() {
	processor := processor.NewTransactionProcessor("")
	processor.ShutdownOnSignal(syscall.SIGINT, syscall.SIGTERM)
}
