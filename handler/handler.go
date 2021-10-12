package handler

import (
	"log"
	"net/http"

	"github.com/chaithanyaMarripati/goSpotify/authorize"
	"github.com/chaithanyaMarripati/goSpotify/spotify"
	"github.com/gin-gonic/gin"
)

func MainApi(spotify spotify.Spotify) gin.HandlerFunc {
	return func(context *gin.Context) {
		accessToken, err := context.Cookie("Token")
		if err != nil {
			log.Println("Couldn't find the token cookie for this request, so redirecting it to the authorize url")
			redirectedUrl := authorize.ConstructAuthorizeReq()
			context.Redirect(http.StatusTemporaryRedirect, redirectedUrl)
			return
		}

		name, err := spotify.Profile(accessToken)
		if err != nil {
			log.Println(err)
			return
		}

		song, err := spotify.CurrentSong(accessToken)
		if err != nil {
			log.Println(err)
			return
		}
		if song == "" {
			context.String(http.StatusOK, "name is %v", name)
			return
		}

		context.String(http.StatusOK, "name is %v\ncurrently listening to %v", name, song)
	}
}
