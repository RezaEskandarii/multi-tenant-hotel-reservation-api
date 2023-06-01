package hash_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	salt = "12345/*$%^&&*(())__+==STR"
)

// GenerateSHA256 generates sha256 hash
func GenerateSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", str, salt)))
	return hex.EncodeToString(h.Sum(nil))
}

// CompareSHA256 compare raw string with given sha256 hash string.
func CompareSHA256(rawStr, hash string) bool {
	var hashedStr = GenerateSHA256(rawStr)
	return hash == hashedStr
}
