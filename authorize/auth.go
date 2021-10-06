package authorize

import (
	"log"
	"net/url"

	"github.com/chaithanyaMarripati/goSpotify/config"
)

//this package handles the authorize part of the application
//here the scopes for the auth request are handled

func ConstructAuthorizeReq() string {
	log.Println("constructing the authorize request with scopes")

	baseUrl := config.EnvVariables.BaseUrl
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", config.EnvVariables.ClientId)
	params.Add("redirect_uri", config.EnvVariables.RedirectUri)
	params.Add("scope", config.EnvVariables.Scopes)
	return baseUrl + params.Encode()
}
