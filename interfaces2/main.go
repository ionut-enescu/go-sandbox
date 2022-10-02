package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error: you must specify exactly 1 arg! You passed", len(os.Args)-1, "args")
		os.Exit(1)
	}

	// ------------- SOLUTION 1 ------------------

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error reading file", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("------------- SOLUTION 1 ------------------")

	bytes, _ := io.Copy(os.Stdout, file)
	fmt.Println("wrote", bytes, "bytes to Stdout using solution 1")

	// -------------  SOLUTION 2 -----------------

	fw := fileWriter{}

	bs, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error reading file", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("------------- SOLUTION 2 ------------------")
	fw.Write(bs)
}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("wrote", len(bs), "bytes to Stdout using solution 2")
	return len(bs), nil
}
