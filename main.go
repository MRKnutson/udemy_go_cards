package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// remainingDeck.printDeck()
	// hand.printDeck()

	// cards.saveToFile("myCards")

	cards := newDeckFromFile("myCards")
	cards.printDeck()
}
