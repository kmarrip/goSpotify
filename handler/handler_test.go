package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setRouter(spotify mocks.SpotifyMock) *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	unauthorized := template.Must(template.ParseFiles("../templates/unauthorized.html"))
	main := template.Must(template.ParseFiles("../templates/main.html"))
	router.SetHTMLTemplate(unauthorized)
	router.SetHTMLTemplate(main)
	router.GET("/", MainApi(spotify))
	router.GET("/callback", CallbackApi())
	return router
}

func TestMainApiHandler(t *testing.T) {
	spotify := mocks.SpotifyMock{}
	spotify.On("CurrentSong", "aToken").Return("MySuperSong", nil)
	spotify.On("Profile", "aToken").Return("MyName", nil)

	router := setRouter(spotify)
	request, _ := http.NewRequest("GET", "/", nil)
	request.AddCookie(&http.Cookie{Name: "Token", Value: "aToken"})

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "<p>MyName</p>")
	assert.Contains(t, response.Body.String(), "<p>MySuperSong</p>")
	spotify.AssertExpectations(t)
}
