package set1

import (
	"testing"
	"bytes"
)

func TestFixedXor(t *testing.T) {
	tests := []struct {
		name           string
		input          [2][]byte
		expectedOutput []byte
		shouldError    bool
	}{
		{
			name: "valid input",
			input: [2][]byte{
				[]byte("1c0111001f010100061a024b53535009181c"),
				[]byte("686974207468652062756c6c277320657965"),
			},
			expectedOutput: []byte("746865206b696420646f6e277420706c6179"),
			shouldError:    false,
		},
		{
			name: "different lenght",
			input: [2][]byte{
				[]byte("1c0111001f010100061a024b53535009181c1cad2"),
				[]byte("686974207468652062756c6c277320657965"),
			},
			expectedOutput: nil,
			shouldError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := CreateXorTwoBuffers(test.input[0], test.input[1])
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
