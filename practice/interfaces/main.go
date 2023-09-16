package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}

func main() {

	eb := englishBot{}
	sn := spanishBot{}

	printGreeting(eb)
	printGreeting(sn)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// its not necesary declarate de name of the receiver if you are not going to use it
func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "¡Hola!"
}
