package yo

import (
	"encoding/json"
	"fmt"
	"github.com/magicshui/goutils/requests"
)

type YoPostData struct {
	ApiToken string `json:""api_token`
	UserName string `json:"username"`
	Link     string `json:"link"`
}

func Yo(token string, username string) {
	url := "https://api.justyo.co/yo/"
	y := new(YoPostData)
	y.ApiToken = token
	y.UserName = username
	y.Link = "error"
	fmt.Println("yo")
	data, _ := json.Marshal(&y)
	r, err := requests.PostHttpsRequest(url, data)
	fmt.Println(string(r), err)

}
