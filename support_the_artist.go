package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Query struct {
	Album Album `xml:"topalbums>album"`
}

type Album struct {
	Title     string `xml:"name"`
	Artist    string `xml:"artist>name"`
	PlayCount string `xml:"playcount"`
}

var lastFMAPIKey string = os.Getenv("LASTFM_KEY")

const APIURL string = "http://ws.audioscrobbler.com/2.0/"

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func getListeningInfo(url string) []byte {
	resp, err := http.Get(url)
	perror(err)
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	perror(err2)
	return body
}

func message(q Query) {
	fmt.Println("Your most played album in the last month is \"" + q.Album.Title +
		"\" by " + q.Album.Artist + ", which you've played " +
		q.Album.PlayCount + " times.")

}

func main() {
	url := APIURL + "?method=user.getTopAlbums&user=" + os.Getenv("LASTFM_USER") +
		"&period=1month&limit=1&api_key=" + os.Getenv("LASTFM_KEY")
	var q Query
	xml.Unmarshal(getListeningInfo(url), &q)
	message(q)
}
