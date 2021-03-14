package main

func main() {

	newcards := newDeck()

	hand, remaing := deal(newcards, 5)

	hand.print()
	remaing.print()

}