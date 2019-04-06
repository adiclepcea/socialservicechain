package main

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func main() {

}

func sha512HashValue(value string) string {
	hashHandler := sha512.New()
	hashHandler.Write([]byte(value))
	return strings.ToLower(hex.EncodeToString(hashHandler.Sum(nil)))
}
