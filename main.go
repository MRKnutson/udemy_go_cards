package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.printDeck()

}

func newCard() string {
	return "Five of Diamonds"
}
