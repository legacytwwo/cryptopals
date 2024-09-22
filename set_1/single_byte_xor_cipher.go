package set1

func DecodeSingleByteXorCipher(buf []byte) ([]byte, error) {
	decodedBuf, err := HexDecode(buf)
	if err != nil {
		return nil, err
	}

	var result []byte
	var bestScore int

	for i := 0; i < 127; i++ {
		xoredText := make([]byte, len(decodedBuf))

		for j, v := range decodedBuf {
			xoredText[j] = v^byte(i)
		}

		score := scoreEngText(xoredText)
		if score > bestScore {
			result = xoredText
			bestScore = score
		}
	}

	return result, nil
}

func scoreEngText(text []byte) int {
	var score int

	letterFrequency := map[byte]int{
        'a': 8167, 'b': 1492, 'c': 2782, 'd': 4253, 'e': 12702,
        'f': 2228, 'g': 2015, 'h': 6094, 'i': 6966, 'j': 153, 
        'k': 772, 'l': 4025, 'm': 2406, 'n': 6749, 'o': 7507, 
        'p': 1929, 'q': 95, 'r': 5987, 's': 6327, 't': 9056, 
        'u': 2758, 'v': 978, 'w': 2360, 'x': 150, 'y': 1974, 'z': 74, ' ': 20000,
    }

	for _, v := range text {
		if val, ok := letterFrequency[v]; ok {
			score += val
		}
	}

	return score
}
