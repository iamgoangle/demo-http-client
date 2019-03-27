package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewGetRequest() {
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func MyRequest() {
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cache-Control", "no-cache")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))
}

func MakeRequest() {
	client := http.Client{}
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}

func main() {
	NewGetRequest()
	NewGetRequestWithHeader()
	MakeRequest()
}
