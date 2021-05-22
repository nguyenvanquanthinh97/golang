package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck of 40, but received %v", len(d))
	}

	if d[0] != "Ace Spades" {
		t.Errorf("Expected first card is Ace Spaces, but received %v", d[0])
	}

	if d[len(d) - 1] != "Ten Clubs" {
		t.Errorf("Expected first card is Ten Clubs, but received %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testDeck")
	
	d := newDeck()
	d.writeFile("_testDeck")

	loadDeck := newDeckFromFile("_testDeck")

	if len(loadDeck) != 40 {
		t.Errorf("Expected 40 cards in deck, but receive %v", len(loadDeck))
	}
	os.Remove("_testDeck")
}
