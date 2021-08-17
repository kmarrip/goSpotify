package main

//road map
//1) "/" path will check if the user currently has access token
//2) if not 304 redirect response has to be sent, with the authorize url
//3) then the user will authorize with spotify
//4) after succesfull auth, the browser will move to token end point, where the code is given to the server
//5) exchange the code with access token and send it as a cookie -- will be using cookies for this
//6) now the user is redirected to "/" with access token, now user is given with a html page, with his name and all
//end of project
import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaithanyaMarripati/goSpotify/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		panic("error reading env file")
	}

	http.HandleFunc("/callback/", handler.TokenHandler)
	http.HandleFunc("/", handler.BaseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
