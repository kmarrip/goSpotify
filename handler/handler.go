package handler

import (
	"fmt"
	"net/http"

	"github.com/chaithanyaMarripati/goSpotify/authorize"
	"github.com/chaithanyaMarripati/goSpotify/spotify"
	"github.com/chaithanyaMarripati/goSpotify/token"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	//1) check if the user has access token in the request
	tokenCookie, err := r.Cookie("Token")
	if err != nil {
		fmt.Println("couldn't find the token cookie for this request")
		fmt.Println("so redirecting it to the authorize url")
		redirectedUrl := authorize.ConstructAuthorizeReq()
		http.Redirect(w, r, redirectedUrl, http.StatusTemporaryRedirect)
		return
	}
	//how we have the token cookie being sent to us for every request
	//use this token cookie, to make requests to the spotify api
	accessToken := tokenCookie.Value
	name, err := spotify.CallSpotifyMe(accessToken)
	if err != nil {
		return
	}
	song, errr := spotify.CallSpotifyCurrentSong(accessToken)
	fmt.Println(song)
	if errr != nil {
		return
	}
	if song == "" {
		fmt.Fprintf(w, "name is %v", name)
		return
	}
	fmt.Fprintf(w, "name is %v\ncurrently listening to %v", name, song)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query()["code"][0]
	//now that we got the code exchange it with access token and refresh token and redirect with set cookie
	token, err := token.GetTokenFromSpotify(authCode)
	if err != nil {
		fmt.Fprintf(w, "faced and issue with token generation")
		return
	}
	accessTokenCookie := &http.Cookie{
		Name:     "Token",
		Value:    token.AccessToken,
		Path:     "/",
		SameSite: http.SameSiteDefaultMode,
	}
	refreshTokenCookie := &http.Cookie{
		Name:     "RefreshToken",
		Value:    token.RefreshToken,
		Path:     "/",
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(w, accessTokenCookie)
	http.SetCookie(w, refreshTokenCookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
