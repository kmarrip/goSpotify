package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

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
	accessToken, err := ctx.Request.Cookie("Token")
	if err != nil || accessToken.Expires.Unix() < time.Now().Unix() {
		fmt.Println("Token is expired or not found")
		authState := uuid.New().String()
		redirectedUrl := authorize.ConstructAuthorizeReq(authState)

		tok, err := ctx.Cookie("RefreshToken")
		if err != nil {
			ctx.SetCookie("State", authState, 120, "/", "", true, true)
			ctx.Redirect(http.StatusTemporaryRedirect, redirectedUrl)
			return
		}

		data, err := token.RefreshSpotifyToken(tok)
		if err != nil {
			ctx.SetCookie("State", authState, 120, "/", "", true, true)
			ctx.Redirect(http.StatusTemporaryRedirect, redirectedUrl)
			return
		}
		ctx.SetCookie("Token", data.AccessToken, 3600, "/", "", true, false)
		if data.RefreshToken == "" {
			data.RefreshToken = tok
		}
		
		ctx.SetCookie("RefreshToken", data.RefreshToken, 3600 * 24 * 7, "/", "", true, false)
		accessToken = &http.Cookie{Value: data.AccessToken}
	}
	//how we have the token cookie being sent to us for every request
	//use this token cookie, to make requests to the spotify api
	name, err := spotify.CallSpotifyMe(accessToken.Value)
	if err != nil {
		log.Panic(err)
		return
	}

	song, err := spotify.CallSpotifyCurrentSong(accessToken.Value)
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
		authState := ctx.QueryArray("state")[0]

		stateCookieVal, err := ctx.Cookie("State")
		if err != nil || authState != stateCookieVal {
			ctx.String(http.StatusInternalServerError, "faced and issue with state verification")
			return
		}

		//now that we got the code exchange it with access token and refresh token and redirect with set cookie
		authToken, err := token.GetTokenFromSpotify(authCode)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "faced and issue with token generation")
			return
		}
		ctx.SetCookie("Token", authToken.AccessToken, 3600, "/", "", true, false)
		ctx.SetCookie("RefreshToken", authToken.RefreshToken, 3600 * 24 * 7, "/", "", true, false)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
