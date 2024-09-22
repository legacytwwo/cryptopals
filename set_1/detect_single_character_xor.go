package set1

func DetectSingleCharXor(buffers [][]byte) ([]byte, error) {
	var result []byte
	var bestScore int
	for _, v := range buffers {
		str, score, err := DecodeSingleByteXorCipher(v)
		if err != nil {
			return nil, err
		}
		if score > bestScore {
			bestScore = score
			result = str
		}
	}
	return result, nil
}
