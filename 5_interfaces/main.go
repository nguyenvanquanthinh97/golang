package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englistBot struct{}
type spanishBot struct{}

func main() {
	var eb englistBot
	var sb spanishBot

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englistBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
