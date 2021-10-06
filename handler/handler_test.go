package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test work with ../templates/unauthorized.html path. I suggest to use Gin framework for improve routing manage and testing purpose
func TestErrorAuthorization(t *testing.T) {
	router := SetupRouter()
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
