package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Number of Goroutines start: %d\n", runtime.NumGoroutine())

	workers := 20

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		time.Sleep(2 * time.Second)

		go func() {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"
				client := http.DefaultClient // bad no timeout by default
				resp, err := client.Get(url)
				if err != nil {
					log.Panic(err)
				}

				// bad no close body
				// cannot to reuse connection,
				// this is not ensured that the http connection could be reused
				// for another request if the keepalive http connection behavior
				//
				// memory leak go routine does not close immediately
				// defer resp.Body.Close()
				if resp != nil {
					defer resp.Body.Close() // MUST CLOSED THIS
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Panic(err)
				}

				log.Println("response", string(body))
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Number of Goroutines end: %d\n", runtime.NumGoroutine())
	fmt.Println("Duration:", time.Since(start).Seconds(), "sec")
}
