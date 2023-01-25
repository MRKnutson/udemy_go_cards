package main

func main() {
	cards := newDeck()
	cards.printDeck()

	// hand, remainingDeck := deal(cards, 5)

	// remainingDeck.printDeck()
	// hand.printDeck()

	// cards.saveToFile("myCards")

	// cards := newDeckFromFile("myCards")
	cards.shuffleDeck()
	cards.printDeck()
}
