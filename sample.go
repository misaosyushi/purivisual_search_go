package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	postSpotify()
}

func postSpotify() {
	values := url.Values{}
	values.Add("grant_type", "client_credentials")

	req, err := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token",
		strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic NGFlYjZiOGVjYTkwNDNjNzkwNGE2ZjI3MmZjMmJhMWI6ZjcwNGRkYmQzMjlmNDAxYTk2MTI4YjU2N2Y2ZTQxNDE=")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	type Token struct {
		AccessToken string `json:"access_token"`
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(Token)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	accessToken := data.AccessToken
	fmt.Println(accessToken)

	getSpotify(accessToken)
}

func getSpotify(accessToken string)  {
	keyword := url.PathEscape("アルルカン")

	url := "https://api.spotify.com/v1/search?q=" + keyword + "&type=artist&market=JP&limit=1&offset=0"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer " + accessToken)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	type Item struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	}
	type SearchResults struct {
		Artists struct{
			Items []Item `json:"items"`
		}
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(SearchResults)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	result := data.Artists.Items[0].ExternalUrls.Spotify
	fmt.Println(result)

	sample := "SORA hogehoge"
	fmt.Println(strings.TrimLeft(sample, "SORA "))
}

