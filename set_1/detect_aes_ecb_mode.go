package set1

import (
	"crypto/aes"
)

func DetectAesEcbMode(buffer []byte) bool {
	if len(buffer)%16 != 0 {
		return false
	}
	blocks := make(map[string]bool)
	for i := 0; i < len(buffer); i += aes.BlockSize {
		key := string(buffer[i : i+aes.BlockSize])
		if _, ok := blocks[key]; ok {
			return true
		}
		blocks[key] = true
	}
	return false
}
