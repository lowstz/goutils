package files

import (
	"bufio"
	"bytes"
	"fmt"
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
