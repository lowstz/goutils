package files

import (
	"testing"
)

func TestIsFileExists(t *testing.T) {
	type Test struct {
		path   string
		ispath bool
	}

	var tests = []Test{
		{".", false},
		{"..", false},
		{"./files.go", true},
		{"./files_test.go", true},
		{"./file_test.go", false},
	}

	for i, test := range tests {
		exists, _ := IsFileExist(test.path)
		if exists != test.ispath {
			t.Errorf("test %d: %s %#v != %#v",
				i, test.path, exists, test.ispath)
		}
	}
}
