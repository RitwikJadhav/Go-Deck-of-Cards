package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("--Cards saved to the local disk--")
	cards.saveToFile("my-cards")
	newCardsDeck := newDeckFromFile("my-cards")
	fmt.Println("--Cards retrieved from the local disk--")
	fmt.Println(newCardsDeck)
	newCardsDeck.shuffle()
	fmt.Println("--Cards deck shuffled in a random order--")
	newCardsDeck.print()
}
