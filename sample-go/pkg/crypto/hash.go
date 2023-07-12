package crypto

import "crypto/sha256"

// Hash returns the sha256 hash of the concatenated given byte slices
func Hash(data ...[]byte) []byte {
	sha := sha256.New()
	for _, d := range data {
		sha.Write(d)
	}

	return sha.Sum(nil)
}

// HashStrings returns the sha256 hash of the concatenated given strings
func HashStrings(strings ...string) []byte {
	sha := sha256.New()
	for _, str := range strings {
		sha.Write([]byte(str))
	}

	return sha.Sum(nil)
}
