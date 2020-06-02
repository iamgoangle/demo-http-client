package main

import(
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"
	client := http.DefaultClient	// bad no timeout by default
	resp, err := client.Get(url)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	log.Println("response", string(body))
}
