package requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func PostRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	req, _ := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return result, err
}
func PostFileRequest(url string, params map[string]string, data []byte, name string) ([]byte, error) {
	b := bytes.NewReader(data)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(name, name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, b)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", url, body)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return result, err
}
func PostHttpsRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("POST", url, b)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return result, err
}

func GetRequest(url string) ([]byte, error) {

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()
	return result, err
}
