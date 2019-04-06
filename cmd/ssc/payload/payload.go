package sscpayload

import (
	"fmt"

	cbor "github.com/brianolson/cbor_go"
)

// SSCPayload is the payload we can receive in our app
type SSCPayload struct {
	Name   string
	Action string
	Space  int
}

//FromBytes creates a SSCPayload struct from a bytes payload
func FromBytes(payload []byte) (*SSCPayload, error) {
	sscPayload := SSCPayload{}
	fmt.Println(string(payload))
	err := cbor.Loads(payload, &sscPayload)
	if err != nil {
		return nil, err
	}
	fmt.Println(sscPayload)
	return &sscPayload, nil
}
