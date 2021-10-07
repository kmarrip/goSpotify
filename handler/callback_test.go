package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/mocks"
	"github.com/stretchr/testify/assert"
)

func TestErrorAuthorization(t *testing.T) {
	router := router(mocks.SpotifyMock{})
	res := httptest.NewRecorder()

	req := &http.Request{
		Method: http.MethodPost,
		Header: http.Header{
			"content-type": []string{"application/json"},
		},
		URL: &url.URL{
			Scheme:   "http",
			Path:     "/callback",
			RawQuery: "error=some_error",
		},
	}

	router.ServeHTTP(res, req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Contains(t, string(body), "<title>Error page</title>")

}
