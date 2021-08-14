package token

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetTokenFromSpotify(code string) {
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
	if err != nil {
		panic("couldn't make post request to token endpoint")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
