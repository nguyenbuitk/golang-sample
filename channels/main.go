package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// creates a channel c using the make function. The channel will be used to communicate the status of the websites.
	c := make(chan string)


	// This allows multiple websites to be checked concurrently.
	for _, link := range links {
		go checkLink(link, c)
	}

	// This creates a periodic check for each website.

	// Each value received by the channel represents the link that was sent through the channel.
	for l := range c{

		// For each received value, a new goroutine is created using the go keyword. This allows the function to be executed concurrently.
		go func(link string) {
			time.Sleep(5*time.Second)
			checkLink(link,c)
		}(l)
	}
}


func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err!= nil {
        fmt.Println(link, "might  be down")
		c <- link
		return
    }

	fmt.Println(link,"is up")
	c <- link
}
