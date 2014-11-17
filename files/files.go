package files

import (
	"errors"
	"io/ioutil"
	"os"
)

func ReadFromLocalFileToByte(path string) ([]byte, error) {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Check the file is exists.
func IsFileExist(path string) (bool, error) {
	absPath := AbsPath(path)
	stat, err := os.Stat(absPath)
	if err == nil {
		if stat.Mode()&os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(path + " exists but is not regular file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
