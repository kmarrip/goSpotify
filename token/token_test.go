package token

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromSpotify(t *testing.T) {
	server := httptest.NewServer(setRouter())

	defer server.Close()

	os.Setenv("tokenUrl", server.URL+"/token")
	os.Setenv("clientId", "aClientId")
	os.Setenv("clientSecret", "aClientSecret")
	os.Setenv("redirectUrl", "redirectUtl")
	config.SetConfigVar()

	response, _ := GetTokenFromSpotify("aCode")

	assert.Equal(t, "aToken", response.AccessToken)

}

func setRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	router.POST("/token", func(c *gin.Context) {
		c.JSON(200, tokenResponse{AccessToken: "aToken", RefreshToken: "aRefresh", TokenType: "type", ExpiresIn: "date", Scope: "scope"})
	})
	return router
}
