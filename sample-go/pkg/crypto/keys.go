package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// PrivateKeyType holds the valid types to generate private keys see iota
type PrivateKeyType int

const (
	// P256 is an ECDSA NIST P256 private key
	P256 PrivateKeyType = iota
)

// GeneratePrivateKey returns a generic private key of the requested type
func GeneratePrivateKey(keyType PrivateKeyType) (privateKey interface{}, err error) {
	switch keyType {
	case P256:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	}

	return nil, errors.New("unsupported private key type")
}

// GetPublicKey returns a generic public key retrieved from the given private key
func GetPublicKey(privateKey interface{}) (publicKey interface{}, err error) {
	switch priv := privateKey.(type) {
	case *ecdsa.PrivateKey:
		return &priv.PublicKey, nil
	}

	return nil, errors.New("no valid private key was provided")
}

// Base64FromPrivateKey returns a base64 encoded PEM(ASN.1) representation of the provided private key
func Base64FromPrivateKey(privateKey interface{}) (string, error) {
	switch priv := privateKey.(type) {
	case *ecdsa.PrivateKey:
		byte, err := x509.MarshalECPrivateKey(priv)
		if err != nil {
			return "", err
		}

		block := &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: byte,
		}

		buf := new(bytes.Buffer)
		if err := pem.Encode(buf, block); err != nil {
			return "", err
		}

		return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
	}

	return "", errors.New("unsupported private key")
}

// LoadPrivateKeyFromBase64 returns the private key extracted from the given base64 encoded PEM format
func LoadPrivateKeyFromBase64(privateKeyBase64 string) (privateKey interface{}, err error) {
	privateKeyPEM, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, errors.New("not able to find private key in pem block")
	}
	if len(rest) != 0 {
		return nil, errors.New("invalid private key format")
	}

	switch block.Type {
	case "EC PRIVATE KEY":
		return x509.ParseECPrivateKey(block.Bytes)
	}

	return nil, fmt.Errorf("unknown private key pem format: %v", block.Type)
}

// LoadPublicKeyFromBase64 returns the public key extracted from the given base64 encoded PEM format
func LoadPublicKeyFromBase64(publicKeyBase64 string) (publicKey interface{}, err error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("not able to find public key in pem block")
	}
	if len(rest) != 0 {
		return nil, errors.New("invalid public key format")
	}

	return ParsePKIXPublicKey(block.Bytes)
}

// Base64FromPublicKey returns a base64 encoded PEM(ASN.1) representation of the provided public key
func Base64FromPublicKey(publicKey interface{}) (string, error) {
	byte, err := MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: byte,
	}

	buf := new(bytes.Buffer)
	if err := pem.Encode(buf, block); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// Sign signs data, that should normally be provided as hash and returns an asn.1 encoded signature
func Sign(hash []byte, privateKey interface{}) ([]byte, error) {
	switch priv := privateKey.(type) {
	case *ecdsa.PrivateKey:
		return ecdsa.SignASN1(rand.Reader, priv, hash)
	}

	return nil, errors.New("unsupported private key for signing")
}
