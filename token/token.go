package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	Scope        string `json:"scope"`
}

func GetTokenFromSpotify(code string) (tokenResponse, error) {
	//now we take the clientid, client secret and auth code to exchange it for the access token and refresh token
	tokenUrl := os.Getenv("tokenUrl")
	const grantType = "authorization_code"
	const contentType = "application/x-www-form-urlencoded"
	clientId := os.Getenv("clientId")
	clientSecret := os.Getenv("clientSecret")
	redirectUri := os.Getenv("redirectUri")
	form := url.Values{}
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", grantType)
	form.Add("code", code)
	form.Add("redirect_uri", redirectUri)
	resp, err := http.Post(tokenUrl, contentType, strings.NewReader(form.Encode()))
	emptyTokenResponse := tokenResponse{}
	if err != nil {
		return emptyTokenResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return emptyTokenResponse, err
	}
	responsePayload := &tokenResponse{}
	json.Unmarshal(body, responsePayload)
	return *responsePayload, nil
}
