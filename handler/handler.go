package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaithanyaMarripati/goSpotify/authorize"
	"github.com/chaithanyaMarripati/goSpotify/spotify"
	"github.com/chaithanyaMarripati/goSpotify/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("../templates/*")
	router.GET("/", baseHandler)
	router.GET("/callback", tokenHandler)
	return router
}

func baseHandler(ctx *gin.Context) {
	//1) check if the user has access token in the request
	accessToken, err := ctx.Cookie("Token")
	if err != nil {
		fmt.Println("couldn't find the token cookie for this request")
		fmt.Println("so redirecting it to the authorize url")
		redirectedUrl := authorize.ConstructAuthorizeReq()
		ctx.Redirect(http.StatusTemporaryRedirect, redirectedUrl)
		return
	}
	//how we have the token cookie being sent to us for every request
	//use this token cookie, to make requests to the spotify api
	name, err := spotify.CallSpotifyMe(accessToken)
	if err != nil {
		log.Panic(err)
		return
	}

	song, err := spotify.CallSpotifyCurrentSong(accessToken)
	if err != nil {
		log.Panic(err)
		return
	}
	if song == "" {
		ctx.String(http.StatusOK, "name is %v", name)
		return
	}

	ctx.String(http.StatusOK, "name is %v\ncurrently listening to %v", name, song)
}

func tokenHandler(ctx *gin.Context) {
	errorCode := ctx.Query("error")
	if len(errorCode) > 0 {
		ctx.HTML(http.StatusInternalServerError, "unauthorized.html", nil)
	} else {
		authCode := ctx.QueryArray("code")[0]

		//now that we got the code exchange it with access token and refresh token and redirect with set cookie
		token, err := token.GetTokenFromSpotify(authCode)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "faced and issue with token generation")
			return
		}
		ctx.SetCookie("Token", token.AccessToken, 3600, "/", "", true, false)
		ctx.SetCookie("RefreshToken", token.RefreshToken, 3600, "/", "", true, false)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
