package handler

import (
	"fmt"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	//1) check if the user has access token in the request
	tokenCookie, err := r.Cookie("token")
	fmt.Println(tokenCookie)
	if err != nil {
		fmt.Println("couldn't find the token cookie for this request")
		fmt.Println("so redirecting it to the authorize url")
	}
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {

}
