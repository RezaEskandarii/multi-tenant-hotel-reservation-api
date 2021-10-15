package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateSHA256 generates sha256 hash
func GenerateSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
