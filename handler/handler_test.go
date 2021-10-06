package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test work with ../templates/unauthorized.html path. I suggest to use Gin framework for improve routing manage and testing purpose
func TestErrorAuthorization(t *testing.T) {
	http.HandleFunc("/callback", TokenHandler)
	go http.ListenAndServe(":8090", nil)

	resp, err := http.Post("http://localhost:8090/callback?error=some_error", "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Contains(t, string(body), "<title>Error page</title>")

}
