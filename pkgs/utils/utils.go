package utils

import (
	"encoding/base64"
	"log"
)

func Encrypt(data string) string {
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	return encodedString
}

func Decrypt(encoded string) string {

	var bytes, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Printf("Err := %v", err.Error())
	}
	var decodedString = string(bytes)
	return decodedString
}
