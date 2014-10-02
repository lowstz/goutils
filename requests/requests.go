package requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
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
