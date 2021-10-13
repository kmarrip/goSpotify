package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaithanyaMarripati/goSpotify/authorize"
	"github.com/chaithanyaMarripati/goSpotify/spotify"
	"github.com/chaithanyaMarripati/goSpotify/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("../templates/*")
	router.GET("/", baseHandler)
	router.GET("/callback", tokenHandler)
	return router
}

func checkExpiration(ctx *gin.Context) bool {

	expirationDate, exp_err := ctx.Cookie("ExpirationDate")

	if exp_err == nil {

		exp_time, parse_error := strconv.ParseInt(expirationDate, 10, 64)

		if parse_error == nil {

			delta := exp_time - time.Now().Unix()

			// 300 seconds are around 5 minutes
			return delta < 300
		}
	}

	return true
}

func setCookies(refreshToken string, accessToken string, expiresData string, ctx *gin.Context) {
	ctx.SetCookie("Token", accessToken, 3600, "/", "", true, false)
	ctx.SetCookie("RefreshToken", refreshToken, 3600, "/", "", true, false)
	ctx.SetCookie("ExpirationDate", expiresData, 3600, "/", "", true, false)
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

func baseHandler(ctx *gin.Context) {
	//1) check if the user has access token in the request
	accessToken, token_err := ctx.Cookie("Token")

	if token_err != nil {
		fmt.Println("couldn't find the token cookie for this request")
		fmt.Println("so redirecting it to the authorize url")
		authState := uuid.New().String()
		redirectedUrl := authorize.ConstructAuthorizeReq(authState)
		ctx.SetCookie("State", authState, 120, "/", "", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, redirectedUrl)
		return
	}

	refreshToken, token_err := ctx.Cookie("RefreshToken")

	if checkExpiration(ctx) == true && token_err == nil {
		authToken, err := token.GetRefreshedToken(refreshToken)

		if err == nil {
			setCookies(authToken.RefreshToken, authToken.AccessToken, authToken.ExpiresIn, ctx)
		}

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

		setCookies(authToken.RefreshToken, authToken.AccessToken, authToken.ExpiresIn, ctx)
	}
}
