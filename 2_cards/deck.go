package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck {}
	cardSuits := deck {"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := deck {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, item := range d {
		fmt.Println(item)
	}
}

func deal(d deck, pos int) (deck, deck) {
	// slice function 
	return d[:pos], d[pos:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeFile(filename string) error {
	// 0666 is a linux permission unix code
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)

	// option 1
	// var result deck
	// if error == nil {
	// 	result = deck(strings.Split(string(bs), ","))
	// }

	// return result

	// option 2
	if error != nil {
		// option #1 - Log the error and return a call to newDeck()
		// option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle () {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
