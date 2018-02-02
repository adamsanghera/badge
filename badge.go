package badge

import (
	"crypto/rand"
	"encoding/hex"
)

// Badge is an interface for services providing identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Badge interface {
	Mint() (interface{}, error)
}

// RandomToken implements Badge
type RandomToken struct {
	tokenLength int
}

// Mint generates a random string, and returns it.
func (t RandomToken) Mint() (interface{}, error) {
	bitString := make([]byte, t.tokenLength)
	_, err := rand.Read(bitString)
	if err != nil {
		return nil, err
	}
	return hex.EncodeToString(bitString), nil
}
