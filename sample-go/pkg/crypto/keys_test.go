package crypto_test

import (
	"crypto/ecdsa"
	"github.com/stretchr/testify/assert"
	"hana-shared-ledger-sample/pkg/crypto"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	check := assert.New(t)

	t.Run("P256", func(t *testing.T) {
		privateKey, err := crypto.GeneratePrivateKey(crypto.P256)
		check.NotNil(privateKey)
		check.NoError(err)
		_, ok := privateKey.(*ecdsa.PrivateKey)
		check.True(ok)
	})
}

func TestPublicKey_Base64(t *testing.T) {
	check := assert.New(t)

	t.Run("P256", func(t *testing.T) {
		privateKey, err := crypto.GeneratePrivateKey(crypto.P256)
		publicKey, err := crypto.GetPublicKey(privateKey)
		check.NoError(err)
		base64PublicKey, err := crypto.Base64FromPublicKey(publicKey)
		check.NoError(err)
		check.NotEmpty(base64PublicKey)

		p, err := crypto.LoadPublicKeyFromBase64(base64PublicKey)
		check.NoError(err)
		newPublicKey := p.(*ecdsa.PublicKey)

		check.True(publicKey.(*ecdsa.PublicKey).Equal(newPublicKey))
	})
}

func TestPrivateKey_Base64(t *testing.T) {
	check := assert.New(t)

	t.Run("P256", func(t *testing.T) {
		privateKey, err := crypto.GeneratePrivateKey(crypto.P256)
		base64PrivateKey, err := crypto.Base64FromPrivateKey(privateKey)
		check.NoError(err)
		check.NotEmpty(base64PrivateKey)

		p, err := crypto.LoadPrivateKeyFromBase64(base64PrivateKey)
		check.NoError(err)
		newPrivateKey := p.(*ecdsa.PrivateKey)

		check.True(privateKey.(*ecdsa.PrivateKey).Equal(newPrivateKey))
	})
}

func TestSignAndVerify(t *testing.T) {
	check := assert.New(t)

	t.Run("P256", func(t *testing.T) {
		testData := []byte("testdata")

		privateKey, _ := crypto.GeneratePrivateKey(crypto.P256)

		sig, err := crypto.Sign(testData, privateKey)
		check.NoError(err)

		err = crypto.Verify(&privateKey.(*ecdsa.PrivateKey).PublicKey, testData, sig)
		check.NoError(err)

		wrongPrivateKey, _ := crypto.GeneratePrivateKey(crypto.P256)

		err = crypto.Verify(&wrongPrivateKey.(*ecdsa.PrivateKey).PublicKey, testData, sig)
		check.EqualError(err, "could not verify signature")
	})
}
