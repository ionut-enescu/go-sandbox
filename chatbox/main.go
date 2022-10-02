package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	e := englishBot{}
	s := spanishBot{}

	print(e)
	print(s)

}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func print(b bot) {
	fmt.Println(b.getGreeting())
}
