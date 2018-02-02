package badge

import (
	"crypto/rand"
	"encoding/hex"
)

// Badge is an interface for services providing identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Badge interface {
	GetID() interface{}
	Mint() error
}

// Token implements Badge
type Token struct {
	tokenLength int
	hash        string
}

func (t Token) GetID() interface{} {
	return t.hash
}

func (t Token) Mint() error {
	bitString := make([]byte, t.tokenLength)
	_, err := rand.Read(bitString)
	if err != nil {
		return err
	}
	t.hash = hex.EncodeToString(bitString)
	return nil
}
