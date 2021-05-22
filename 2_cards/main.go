package main

func main() {
	// Declare Variable
	// Method 1: var card string = "Ace of Spades"
	// Method 2: use ":=" to tell card is a new variable of type base of the right-handed value
	// When re-assign value to card, we don't need use ":" anymore (":" mean telling new variable)
	// card := "Ace of Spades"

	// Declare slice type
	// cards := [] string {
	// 	"Ace of Spades",
	// 	newCard(),
	// }

	cards := newDeck()

	// cards = append(cards, "Six of Spades")

	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// fmt.Println("----------------------------")
	
	// remainingCards.print()

	cards.writeFile("my_cards")

	cards = newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
