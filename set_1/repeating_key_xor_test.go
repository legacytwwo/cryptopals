package set1

import (
	"bytes"
	"testing"
)

func TestRepeatingKeyXor(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		key            []byte
		expectedOutput []byte
	}{
		{
			name:           "valid input",
			key:            []byte("ICE"),
			input:          []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
			expectedOutput: []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"),
		},
	}

	for _, test := range tests {
		ret := RepeatingKeyXor(test.input, test.key)
		if !bytes.Equal(HexEncode(ret), test.expectedOutput) {
			t.Errorf("expected: %s, got: %s", test.expectedOutput, ret)
		}
	}

}
