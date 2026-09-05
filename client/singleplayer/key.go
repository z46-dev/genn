package singleplayer

import (
	"crypto/rand"
	"encoding/hex"
)

var MetaAllowedKey string

func init() {
	var key [32]byte
	var err error
	if _, err = rand.Read(key[:]); err != nil {
		panic("generate random key: " + err.Error())
	}

	MetaAllowedKey = hex.EncodeToString(key[:])
}
