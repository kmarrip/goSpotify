package config

import "os"

type EnvVar struct {
	BaseUrl          string
	ClientId         string
	RedirectUri      string
	Scopes           string
	TokenUrl         string
	ClientSecret     string
	GetMeSpotify     string
	CurrentlyPlaying string
}

var EnvVariables EnvVar

func SetConfigVar() {
	EnvVariables = EnvVar{
		BaseUrl:          os.Getenv("baseUrl"),
		ClientId:         os.Getenv("clientId"),
		RedirectUri:      os.Getenv("redirectUrl"),
		Scopes:           os.Getenv("scopes"),
		TokenUrl:         os.Getenv("tokenUrl"),
		ClientSecret:     os.Getenv("clientSecret"),
		GetMeSpotify:     os.Getenv("getMeSpotify"),
		CurrentlyPlaying: os.Getenv("currentlyPlaying"),
	}
}
