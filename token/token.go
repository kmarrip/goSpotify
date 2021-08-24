package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chaithanyaMarripati/goSpotify/config"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	ExpiresAt    string `json:"expiresAt"`
}

func GetTokenFromSpotify(code string) (tokenResponse, error) {
	//now we take the clientid, client secret and auth code to exchange it for the access token and refresh token
	tokenUrl := config.EnvVariables.TokenUrl
	const grantType = "authorization_code"
	const contentType = "application/x-www-form-urlencoded"
	clientId := config.EnvVariables.ClientId
	clientSecret := config.EnvVariables.ClientSecret
	redirectUri := config.EnvVariables.RedirectUri
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
	currentTime := time.Now().Add(time.Duration(responsePayload.ExpiresIn)).Format(time.RFC3339)
	responsePayload.ExpiresAt = currentTime
	return *responsePayload, nil
}
