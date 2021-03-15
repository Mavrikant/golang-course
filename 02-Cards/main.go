package main

func main() {

/*	newcards := newDeck()
	hand, remaing := deal(newcards, 5) // Multiple return
	hand.print()
	remaing.print()
*/
	// greeting := "Hello GOlang"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.toString()

	// cards:=newDeck()
	// cards.saveToFile("my_cards")

	cards:=newDeckFromFile("my_cards")
	 cards.print()
	 cards.shuffle()
	cards.print()



}