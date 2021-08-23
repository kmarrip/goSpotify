package authorize

import (
	"fmt"
	"net/url"

	"github.com/chaithanyaMarripati/goSpotify/config"
)

//this package handles the authorize part of the application
//here the scopes for the auth request are handled

func ConstructAuthorizeReq() string {
	fmt.Println("constructing the authorize request with scopes")
	scopes := config.EnvVariables.Scopes
	baseUrl := config.EnvVariables.BaseUrl
	clientId := config.EnvVariables.ClientId
	const responseType = "code"
	redirectUri := config.EnvVariables.RedirectUri
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("response_type", responseType)
	params.Add("redirect_uri", redirectUri)
	params.Add("scope", scopes)
	return baseUrl + params.Encode()
}
