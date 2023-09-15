package main

import "fmt"

// var card string

func main() {
	cards := newDeck()
	//// cards = append(cards, "Six of Spades")
	//
	//hand, deck := deal(cards, 5)
	//
	//hand.print()
	//deck.print()

	//greeting := "Hi there!"

	//fmt.Println([]byte(greeting))

	cards.saveToFile("box")

	other := loadFromFile("box")

	other.randomize()
	fmt.Println(other)

	//fmt.Println(cards.toString())

}
