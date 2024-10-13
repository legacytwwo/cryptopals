package set1

import (
	"crypto/aes"
	"errors"
)

func DecodeAes128Ecb(data, key []byte) ([]byte, error) {
	if len(data)%16 != 0 {
		return nil, errors.New("input encrypted data is invalid")
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decodedCipher := make([]byte, len(data))
	for i := 0; i < len(data); i += aes.BlockSize {
		cipher.Decrypt(decodedCipher[i:i+aes.BlockSize], data[i:i+aes.BlockSize])
	}
	return decodedCipher, nil
}
