package main

import "fmt"

func main() {
	//var card string = "Ace of spades"
	card := newCard()
	// card = "Hello World"
	// fmt.Printf("data type of card is: %T \n " + card)
	fmt.Println(card)
}

func newCard() string {
	return "Five of Dimands"
}
