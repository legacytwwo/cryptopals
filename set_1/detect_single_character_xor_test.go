package set1

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var FilePath = "./static/detect_single_character_xor.txt"

func TestDetectSingleCharXor(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		expectedOutput []byte
	}{
		{
			name:           "valid input",
			filePath:       FilePath,
			expectedOutput: []byte("Now that the party is jumping\n"),
		},
	}
	for _, test := range tests {
		var input [][]byte
		file, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		data := strings.Split(string(file), "\n")
		for _, v := range data {
			encodedStr, _ := HexDecode([]byte(v))
			input = append(input, encodedStr)
		}
		ret, err := DetectSingleCharXor(input)
		if err != nil {
			t.Errorf("expected: nil error, got: %s", err)
		}
		if !bytes.Equal(ret, test.expectedOutput) {
			t.Errorf("expected: %s, got: %s", test.expectedOutput, ret)
		}
	}
}
