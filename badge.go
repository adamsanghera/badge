package badge

import (
	"crypto/rand"
	"encoding/hex"
)

// Minter is an interface for minting identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Minter interface {
	Mint() (interface{}, error)
}

// RandomTokenMinter implements Minter
type RandomTokenMinter struct {
	tokenLength int
}

// Mint generates a random string, and returns it.
func (t RandomTokenMinter) Mint() (interface{}, error) {
	bitString := make([]byte, t.tokenLength)
	_, err := rand.Read(bitString)
	if err != nil {
		return nil, err
	}
	return hex.EncodeToString(bitString), nil
}
