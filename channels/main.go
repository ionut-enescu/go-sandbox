package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://google.com", "http://facebook.com", "http://hotnews.ro", "http://cnn.com", "http://darklyrics.com"}

	c := make(chan string)

	for _, link := range urls {
		go checkLink(link, c) // runs a new go routine each iteration
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Link", url, "might be down: ", err)
		c <- url
		return
	}

	fmt.Println("Link", url, "is up")
	c <- url

}
