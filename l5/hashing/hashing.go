package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hash struct {
	Available bool
}

func (Hash) HashSha256(message string) string {
	hashed := sha256.Sum256([]byte(message))
	response := hex.EncodeToString(hashed[:])

	//fmt.Println(response)
	return response
}
