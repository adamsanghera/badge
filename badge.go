package badge

import (
	"crypto/rand"
	"encoding/hex"
	"reflect"
)

// Minter is an interface for minting identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Minter interface {
	Mint() (badge interface{}, badgeType reflect.Kind, err error)
}

// RandomTokenMinter implements Minter
type RandomTokenMinter struct {
	TokenLength int
}

// NewRandomTokenMinter creates a new random token minter, which mints tokens of size n.
func NewRandomTokenMinter(tokenLength int) RandomTokenMinter {
	return RandomTokenMinter{
		TokenLength: tokenLength,
	}
}

// Mint generates a random string, and returns it.
func (t RandomTokenMinter) Mint() (badge interface{}, badgeType reflect.Kind, err error) {
	bitString := make([]byte, t.tokenLength)
	_, err = rand.Read(bitString)
	if err != nil {
		return nil, 0, err
	}
	return hex.EncodeToString(bitString), reflect.TypeOf("").Kind(), nil
}
