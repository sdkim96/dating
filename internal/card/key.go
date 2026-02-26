package card

import (
	"crypto/sha256"
	"encoding/hex"
)

func createKey(tenant string, name string) string {
	data := tenant + ":" + name
	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}
