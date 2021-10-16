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
	"log"

	"github.com/chaithanyaMarripati/goSpotify/config"
	"github.com/chaithanyaMarripati/goSpotify/handler"
	"github.com/chaithanyaMarripati/goSpotify/spotify"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Println("can't read env files, env should have already been set")
	}
	config.SetConfigVar()

	spotify := spotify.HttpSpotify{}
	router := gin.Default()
	router.LoadHTMLFiles("../templates/*")
	router.GET("/", handler.MainApi(&spotify))
	router.GET("/callback", handler.CallbackApi())

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
