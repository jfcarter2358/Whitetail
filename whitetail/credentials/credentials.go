// models.credentials.go

package Credentials

import (
    "fmt"
    "encoding/json"
    "net/http"
	"io/ioutil"
	"os"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

var AccessToken string
var TokenType string

func GetAccessToken() {
	username := os.Getenv("MOC_ADMIN_OAUTH_USERNAME")
	password := os.Getenv("MOC_ADMIN_OAUTH_PASSWORD")
	oauth_url := os.Getenv("MOC_ADMIN_OAUTH_URL")

	hc := http.Client{}
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	req, _ := http.NewRequest("POST", oauth_url, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)

	if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

	defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    tokenResponse := TokenResponse{}
    fmt.Println("TOKEN RESPONSE : " + string(body))
    json.Unmarshal(body, &tokenResponse)

	AccessToken = tokenResponse.AccessToken	
	TokenType = tokenResponse.TokenType
}