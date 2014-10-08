package pgp

import (
	"bufio"
	"bytes"
	"code.google.com/p/go.crypto/openpgp"
	//"code.google.com/p/go.crypto/openpgp/armor"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFromLocalFileToByte(path string) ([]byte, error) {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SaveToLocalFileWithByte(path string, data []byte) error {
	error := ioutil.WriteFile(path, data, 0777)
	return error
}

func ReadFromLocalFileToLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func SaveToLocalFileWithLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func PgpEncode(flatFilePath string, pgpFilePath string, publicKeyPath string) error {
	publicKey, err := os.Open(publicKeyPath)
	if err != nil {
		return err
	}
	defer publicKey.Close()
	entitylist, err := openpgp.ReadArmoredKeyRing(publicKey)
	//entitylist, err := openpgp.ReadKeyRing(publicKey)

	if err != nil {
		fmt.Println(1, err)
		return err
	}
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entitylist, nil, nil, nil)
	if err != nil {
		return err
	}
	flatData, err := ReadFromLocalFileToByte(flatFilePath)
	if err != nil {
		fmt.Println(2, err)
		return err
	}
	_, err = w.Write(flatData)
	if err != nil {
		fmt.Println(3, err)
		return err
	}
	err = w.Close()
	if err != nil {
		fmt.Println(4, err)
		return err
	}
	bytesp, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	err = SaveToLocalFileWithByte(pgpFilePath, bytesp)
	if err != nil {
		return err
	}
	return nil

}
func PgpEncodeWithStringKey(flatFilePath string, pgpFilePath string, publicKey string) error {

	entitylist, err := openpgp.ReadArmoredKeyRing(strings.NewReader(publicKey))
	//entitylist, err := openpgp.ReadKeyRing(publicKey)

	if err != nil {
		fmt.Println(1, err)
		return err
	}
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entitylist, nil, nil, nil)
	if err != nil {
		return err
	}
	flatData, err := ReadFromLocalFileToByte(flatFilePath)
	if err != nil {
		fmt.Println(2, err)
		return err
	}
	_, err = w.Write(flatData)
	if err != nil {
		fmt.Println(3, err)
		return err
	}
	err = w.Close()
	if err != nil {
		fmt.Println(4, err)
		return err
	}
	bytesp, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	err = SaveToLocalFileWithByte(pgpFilePath, bytesp)
	if err != nil {
		return err
	}
	return nil

}
func PgpDecodeWithStringKey(pgpFilePath string, flatFilePath string, privateKey string) error {
	entitylist, err := openpgp.ReadArmoredKeyRing(strings.NewReader(privateKey))

	if err != nil {
		fmt.Println(2, err)
		return err
	}

	encryptedMessage, err := ReadFromLocalFileToByte(pgpFilePath)
	if err != nil {
		fmt.Println(3, err)
		return err
	}
	decbuf := bytes.NewBuffer(encryptedMessage)

	md, err := openpgp.ReadMessage(decbuf, entitylist, nil, nil)
	/*

	   result, err := armor.Decode(decbuf)
	   if err != nil {
	       fmt.Println(4, err)
	       return err
	   }
	   md, err := openpgp.ReadMessage(result.Body, entitylist, nil, nil)
	*/
	if err != nil {
		fmt.Println(err)
		return err
	}

	results, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		fmt.Println(6, err)
		return err
	}
	err = SaveToLocalFileWithByte(flatFilePath, results)
	return err
}
