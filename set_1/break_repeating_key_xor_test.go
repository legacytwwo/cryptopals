package set1

import (
	"bytes"
	"encoding/base64"
	"os"
	"testing"
)

var RepeatingKeyFilePath = "./static/break_repeating_key_xor.txt"
var RepeatingKeyOutputFilePath = "./static/break_repeating_key_xor_output.txt"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		name   string
		buf1   []byte
		buf2   []byte
		output int
	}{
		{
			name:   "valid input",
			buf1:   []byte("this is a test"),
			buf2:   []byte("wokka wokka!!!"),
			output: 37,
		},
	}
	for _, test := range tests {
		ret, _ := HammingDistance(test.buf1, test.buf2)
		if ret != test.output {
			t.Errorf("expected: %v, got: %v", test.output, ret)
		}
	}
}

func TestBreakRepeatingKeyXor(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		outputfilePath string
	}{
		{
			name:           "valid input",
			filePath:       RepeatingKeyFilePath,
			outputfilePath: RepeatingKeyOutputFilePath,
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
		ret := BreakRepeatingKeyXor(ueb)
		outputFile, err := os.ReadFile(test.outputfilePath)
		if err != nil {
			t.Errorf("error while open file: %s", err)
		}
		if !bytes.Equal(ret, outputFile) {
			t.Errorf("expected: %s, got: %s", outputFile, ret)
		}
	}
}
