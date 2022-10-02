package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, response.Body)

	lw := logWriter{}
	io.Copy(lw, response.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("wrote", len(p), "bytes to Stdout")
	return len(p), nil
}
