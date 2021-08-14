package main

import (
	"fmt"

	"github.com/chaithanyaMarripati/goSpotify/authorize"
	"github.com/chaithanyaMarripati/goSpotify/token"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		panic("error reading env file")
	}
	authorizeURL := authorize.ConstructAuthorizeReq()
	fmt.Println(authorizeURL)
	var code string
	fmt.Scanf("%s", &code)
	token.GetTokenFromSpotify(code)
}
