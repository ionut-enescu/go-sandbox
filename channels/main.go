package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"http://google.com", "http://facebook.com", "http://hotnews.ro", "http://cnn.com", "http://darklyrics.com"}

	c := make(chan string)

	for _, link := range urls {
		go checkLink(link, c) // runs a new go routine each iteration
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Link", url, "might be down: ", err)
		c <- "Might be down: " + url
		return
	}

	fmt.Println("Link", url, "is up")
	c <- "It's up: " + url

}
