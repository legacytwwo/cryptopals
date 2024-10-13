package set1

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var DetectAesEcbFilePath = "./static/detect_aes_ecb_mode.txt"
var DetectAesEcbOutputFilePath = "./static/detect_aes_ecb_mode_output.txt"

func TestDetect(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		outputfilePath string
	}{
		{
			name:           "valid input",
			filePath:       DetectAesEcbFilePath,
			outputfilePath: DetectAesEcbOutputFilePath,
		},
	}
	for _, test := range tests {
		var ret []byte
		file, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		data := strings.Split(string(file), "\n")
		for _, v := range data {
			encodedStr, _ := HexDecode([]byte(v))
			if status := DetectAesEcbMode(encodedStr); status {
				ret = encodedStr
				break
			}
		}
		outputFile, err := os.ReadFile(test.outputfilePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		if !bytes.Equal(HexEncode(ret), outputFile) {
			t.Errorf("expected: %s, got: %s", outputFile, ret)
		}
	}
}
