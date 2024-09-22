package set1

import (
	"errors"
)

func CreateXorTwoBuffers(buf1, buf2 []byte) ([]byte, error) {
	if len(buf1) != len(buf2) {
		return nil, errors.New("the buffers have different lengths")
	}

	decodedBuf1, err := HexDecode(buf1)
	if err != nil {
		return nil, err
	}
	decodedBuf2, err := HexDecode(buf2)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(decodedBuf1))

	for i, v := range decodedBuf1 {
		result[i] = v^decodedBuf2[i]
	}

	return HexEncode(result), nil
}
