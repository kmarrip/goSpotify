package handler

import (
	"net/http"

	"github.com/chaithanyaMarripati/goSpotify/token"
	"github.com/gin-gonic/gin"
)

func CallbackApi() gin.HandlerFunc {
	return func(context *gin.Context) {
		errorCode := context.Query("error")
		if len(errorCode) > 0 {
			context.HTML(http.StatusOK, "unauthorized.html", nil)
		} else {
			authCode := context.QueryArray("code")[0]

			//now that we got the code exchange it with access token and refresh token and redirect with set cookie
			token, err := token.GetTokenFromSpotify(authCode)
			if err != nil {
				context.String(http.StatusInternalServerError, "faced and issue with token generation")
				return
			}
			context.SetCookie("Token", token.AccessToken, 3600, "/", "", true, false)
			context.SetCookie("RefreshToken", token.RefreshToken, 3600, "/", "", true, false)
			context.Redirect(http.StatusTemporaryRedirect, "/")
		}

	}
}
