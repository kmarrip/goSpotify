package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/mocks"
	"github.com/stretchr/testify/assert"
)

func TestErrorAuthorization(t *testing.T) {
	router := setRouter(mocks.SpotifyMock{})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/callback?error=some_error", nil)

	router.ServeHTTP(res, req)

	assert.Contains(t, res.Body.String(), "<title>Error page</title>")
}
