package main

import (
	"bytes"
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
const eventType string = "lastfm_api"

var ifTTTAPIURL string = "https://maker.ifttt.com/trigger/" + eventType + "/with/key/" + os.Getenv("IFTTT_API_KEY")

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

func sendSMS(q Query) {
	jsonReq := fmt.Sprintf(`{"value1":"%s","value2":"%s","value3":"%s"}`,
		q.Album.Title, q.Album.Artist, q.Album.PlayCount)
	jsonStr := []byte(jsonReq)
	jsonBody := bytes.NewBuffer(jsonStr)
	fmt.Println(ifTTTAPIURL)
	req, _ := http.NewRequest("POST", ifTTTAPIURL, jsonBody)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}

func main() {
	url := APIURL + "?method=user.getTopAlbums&user=" + os.Getenv("LASTFM_USER") +
		"&period=1month&limit=1&api_key=" + os.Getenv("LASTFM_KEY")
	var q Query
	xml.Unmarshal(getListeningInfo(url), &q)
	sendSMS(q)
}
