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

func TestMainApiHandler(t *testing.T) {
	spotify := mocks.SpotifyMock{}
	spotify.On("CurrentSong", "aToken").Return("MySuperSong", nil)
	spotify.On("Profile", "aToken").Return("MyName", nil)

	router := router(spotify)
	request, _ := http.NewRequest("GET", "/", nil)
	request.AddCookie(&http.Cookie{Name: "Token", Value: "aToken"})

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	spotify.AssertExpectations(t)

}

func router(spotify mocks.SpotifyMock) *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	html := template.Must(template.ParseFiles("../templates/unauthorized.html"))
	router.SetHTMLTemplate(html)
	router.GET("/", MainApi(spotify))
	return router
}
