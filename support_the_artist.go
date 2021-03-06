package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Query struct {
	Album Album `xml:"topalbums>album"`
}

type Album struct {
	Title     string `xml:"name"`
	Artist    string `xml:"artist>name"`
	PlayCount string `xml:"playcount"`
}

const APIURL string = "http://ws.audioscrobbler.com/2.0/"
const eventType string = "lastfm_api"

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

func jsonBody(q Query) *bytes.Buffer {
	jsonReq := fmt.Sprintf(`{"value1":"%s","value2":"%s","value3":"%s"}`,
		q.Album.Title, q.Album.Artist, q.Album.PlayCount)
	jsonStr := []byte(jsonReq)
	reqBody := bytes.NewBuffer(jsonStr)
	return reqBody
}

func sendSMS(q Query) {
	ifTTTAPIURL := "https://maker.ifttt.com/trigger/" + eventType +
		"/with/key/" + os.Getenv("IFTTT_API_KEY")
	req, _ := http.NewRequest("POST", ifTTTAPIURL, jsonBody(q))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	perror(err)
	fmt.Println(resp.Status)
}

func main() {
	err := godotenv.Load()
	perror(err)
	url := APIURL + "?method=user.getTopAlbums&user=" + os.Getenv("LASTFM_USER") +
		"&period=1month&limit=1&api_key=" + os.Getenv("LASTFM_KEY")
	var q Query
	xml.Unmarshal(getListeningInfo(url), &q)
	sendSMS(q)
}
