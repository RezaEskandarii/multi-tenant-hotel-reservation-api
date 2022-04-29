package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateSHA256 generates sha256 hash
func GenerateSHA256(str string) string {
	h := sha256.New()
	salt := "12345/*$%^&&*(())__+==STR"
	h.Write([]byte(fmt.Sprintf("%s%s", str, salt)))
	return hex.EncodeToString(h.Sum(nil))
}
