package crypto

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
)

// NewSigFromBase64 just decodes a base64 string.
// The result is most likely a asn.1 representation of a specific signature.
func NewSigFromBase64(sigBase64 string) ([]byte, error) {
	decodedSig, err := base64.StdEncoding.DecodeString(sigBase64)
	if err != nil {
		return nil, err
	}
	return decodedSig, nil
}

// Verify checks a given signature against the original hash and corresponding public key
func Verify(publicKey interface{}, hash []byte, signature []byte) error {
	switch pub := publicKey.(type) {
	case *ecdsa.PublicKey:
		ok := ecdsa.VerifyASN1(pub, hash, signature)
		if !ok {
			return errors.New("could not verify signature")
		}
		return nil
	}

	return errors.New("unsupported public key type, maybe private key was provided")

}
