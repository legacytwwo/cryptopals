package set1

import (
	"bytes"
	"testing"
)

func TestSingleByteXorCipher(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldError    bool
	}{
		{
			name: "valid input",
			input: []byte(
				"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			),
			expectedOutput: []byte("Cooking MC's like a pound of bacon"),
			shouldError:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, _, err := DecodeSingleByteXorCipher(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(ret, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, ret)
			}
		})
	}
}
