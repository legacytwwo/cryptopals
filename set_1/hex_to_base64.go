package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func ConverHexToBase64(hexBytes []byte) ([]byte, error) {
	decodedBytes, err := HexDecode(hexBytes)
	if err != nil {
		return nil, err
	}
	base64Bytes := make([]byte, base64.StdEncoding.EncodedLen(len(decodedBytes)))
	base64.StdEncoding.Encode(base64Bytes, decodedBytes)
	return base64Bytes, nil
}

func HexDecode(buf []byte) ([]byte, error) {
	result := make([]byte, hex.DecodedLen(len(buf)))
	_, err := hex.Decode(result, buf)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func HexEncode(buf []byte) []byte {
	result := make([]byte, hex.EncodedLen(len(buf)))
	hex.Encode(result, buf)
	return result
}
