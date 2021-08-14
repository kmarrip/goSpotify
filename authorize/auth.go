package authorize

import (
	"fmt"
	"net/url"
	"os"
)

//this package handles the authorize part of the application
//here the scopes for the auth request are handled

func ConstructAuthorizeReq() string {
	fmt.Println("constructing the authorize request with scopes")
	scopes := os.Getenv("scopes")
	baseUrl := os.Getenv("baseUrl")
	clientId := os.Getenv("clientId")
	const responseType = "code"
	redirectUri := os.Getenv("redirectUri")
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("response_type", responseType)
	params.Add("redirect_uri", redirectUri)
	params.Add("scope", scopes)
	return baseUrl + params.Encode()
}
