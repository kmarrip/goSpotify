package spotify

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	t.Run("should call profile", func(t *testing.T) {
		httpSpotify := HttpSpotify{}
		router := setRouter()
		server := httptest.NewServer(router)

		os.Setenv("getMeSpotify", server.URL+"/profile")
		config.SetConfigVar()

		name, _ := httpSpotify.Profile("aToken")

		assert.Equal(t, "pong", name)

		defer server.Close()
	})
}

func TestCurrentSong(t *testing.T) {
	t.Run("should call currentSong", func(t *testing.T) {
		httpSpotify := HttpSpotify{}
		router := setRouter()
		server := httptest.NewServer(router)

		os.Setenv("currentlyPlaying", server.URL+"/current")
		config.SetConfigVar()

		song, _ := httpSpotify.CurrentSong("aToken")

		assert.Equal(t, "ping", song)

		defer server.Close()
	})
}

func setRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	router.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"display_name": "pong",
		})
	})
	router.GET("/current", func(c *gin.Context) {
		song := currentSong{Item: Item{Album: Album{Name: "ping"}}}
		c.JSON(200, song)
	})
	return router
}
