package set1

func RepeatingKeyXor(input, key []byte) []byte {
	var result []byte
	p := 0
	for _, v := range input {
		if p >= len(key) {
			p = 0
		}
		result = append(result, v^key[p])
		p += 1
	}
	return result
}
