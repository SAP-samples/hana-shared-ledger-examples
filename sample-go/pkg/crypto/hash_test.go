package crypto_test

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"hana-shared-ledger-sample/pkg/crypto"
	"testing"
)

func TestHash(t *testing.T) {
	check := assert.New(t)

	some, data := []byte("some"), []byte("data")
	hexHash := hex.EncodeToString(crypto.Hash(some, data))

	check.Equal("87d149cb424c0387656f211d2589fb5b1e16229921309e98588419ccca8a7362", hexHash)
}

func TestHashStrings(t *testing.T) {
	check := assert.New(t)

	some, data := "some", "data"
	hexHash := hex.EncodeToString(crypto.HashStrings(some, data))

	check.Equal("87d149cb424c0387656f211d2589fb5b1e16229921309e98588419ccca8a7362", hexHash)
}
