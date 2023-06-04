package main

func main() {

	cards := createDeck()

	hand, remainingCards := cards.deal(5)

	hand.print()
	remainingCards.print()

	cards.writeToFile("cards_file")

	cards := readFromFile("cards_file")

	cards.shuffle()

	cards.print()
}
