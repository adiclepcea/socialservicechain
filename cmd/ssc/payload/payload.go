package sscpayload

// SSCPayload is the payload we can receive in our app
type SSCPayload struct {
	Name   string
	Action string
	Space  int
}

//FromBytes creates a SSCPayload struct from a bytes payload
func FromBytes(payload []byte) (SSCPayload, error) {
	return SSCPayload{}, nil
}
