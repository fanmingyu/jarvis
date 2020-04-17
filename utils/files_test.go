package utils

import(
	"testing"
)

func TestIsDir(t *testing.T) {
	path := "../utils"
	result := IsDir(path)

	if !result {
		t.Errorf("test exist dir failed.")
	}

	path = "./utils"
	result = IsDir(path)

	if result {
		t.Errorf("test no exist dir failed.")
	}
}
