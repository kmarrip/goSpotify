//here we get the token
//with the token, we need to call the spotify for name and currently listening songs

package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type meSpotify struct {
	DisplayName string `json:"display_name"`
}
type currentSong struct {
	Item struct {
		Album struct {
			Name string `json:"name"`
		} `json:"album"`
	} `json:"item"`
}

func CallSpotifyCurrentSong(token string) (string, error) {
	getCurrentSpotify := os.Getenv("currentlyPlaying")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", getCurrentSpotify, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var responsePayload currentSong
	json.Unmarshal(body, &responsePayload)
	fmt.Println(responsePayload)
	fmt.Println(responsePayload.Item)
	fmt.Println(responsePayload.Item.Album)
	fmt.Println(responsePayload.Item.Album.Name)
	return responsePayload.Item.Album.Name, nil
}
func CallSpotifyMe(token string) (string, error) {
	//call the spotify api for user name and current song
	getUserSpotify := os.Getenv("getMeSpotify")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", getUserSpotify, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	responsePayload := &meSpotify{}
	json.Unmarshal(body, responsePayload)
	return responsePayload.DisplayName, nil
}
