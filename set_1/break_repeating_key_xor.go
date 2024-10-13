package set1

import (
	"errors"
	"sort"
)

type KeysizeWeight struct {
	keysize int
	weight  float64
}

func HammingDistance(buf1 []byte, buf2 []byte) (int, error) {
	if len(buf1) != len(buf2) {
		return 0, errors.New("the buffers have different lengths")
	}
	var result int
	for i := 0; i < len(buf1); i++ {
		for j := 0; j < 8; j++ {
			mask := byte(0x01 << j)
			if (buf1[i] & mask) != (buf2[i] & mask) {
				result++
			}
		}
	}
	return result, nil
}

func FindKeysizeWeights(input []byte) []KeysizeWeight {
	keysize, maxKeysize := 2, 40
	result := make([]KeysizeWeight, 0)
	for keysize <= maxKeysize {
		var hd float64
		for i := 0; i < 3; i++ {
			border := i * keysize
			dist, err := HammingDistance(input[border:border+keysize], input[border+keysize:border+(keysize*2)])
			if err != nil {
				break
			}
			hd += float64(dist)
		}
		hd = hd / float64(keysize)
		result = append(result, KeysizeWeight{keysize: keysize, weight: hd})
		keysize++
	}
	sort.Slice(result, func(i, j int) bool { return result[i].weight < result[j].weight })
	return result
}

func BreakCiphertextToBlocks(input []byte, keysize int) [][]byte {
	result := make([][]byte, keysize)
	for i := range keysize {
		result[i] = append(result[i], input[i])
		for j := i + keysize; j <= len(input)-1; j += keysize {
			result[i] = append(result[i], input[j])
		}
	}
	return result
}

func BreakRepeatingKeyXor(input []byte) []byte {
	var result []byte
	var bestScore int
	keysizeWeights := FindKeysizeWeights(input)
	for _, v := range keysizeWeights[:3] {
		blocks := BreakCiphertextToBlocks(input, v.keysize)
		repeatingKey := make([]byte, v.keysize)
		for i := range len(blocks) {
			_, _, key, err := DecodeSingleByteXorCipher(blocks[i])
			if err != nil {
				break
			}
			repeatingKey[i] = byte(key)
		}
		r := RepeatingKeyXor(input, repeatingKey)
		score := scoreEngText(r)
		if score > bestScore {
			bestScore = score
			result = r
		}
	}
	return result
}
