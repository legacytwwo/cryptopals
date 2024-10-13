package set1

import (
	"bytes"
	"encoding/base64"
	"os"
	"testing"
)

var Aes128EcbFilePath = "./static/aes_128_ecb.txt"
var Aes128EcbOutputFilePath = "./static/aes_128_ecb_output.txt"

func TestDecodeAes128Ecb(t *testing.T) {
	tests := []struct {
		name           string
		key            []byte
		filePath       string
		outputfilePath string
	}{
		{
			name:           "valid input",
			key:            []byte("YELLOW SUBMARINE"),
			filePath:       Aes128EcbFilePath,
			outputfilePath: Aes128EcbOutputFilePath,
		},
	}
	for _, test := range tests {
		file, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		ueb := make([]byte, base64.StdEncoding.DecodedLen(len(file)))
		_, err = base64.StdEncoding.Decode(ueb, file)
		ueb = bytes.Trim(ueb, "\x00")
		if err != nil {
			t.Errorf("error while decode base64: %s", err)
		}
		ret, err := DecodeAes128Ecb(ueb, test.key)
		if err != nil {
			t.Errorf("error while decode aes 128: %s", err)
		}
		outputFile, err := os.ReadFile(test.outputfilePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		if !bytes.Equal(ret, outputFile) {
			t.Errorf("expected: %s, got: %s", outputFile, ret)
		}
	}
}
