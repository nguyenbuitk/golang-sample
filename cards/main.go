package main

// import "fmt"

func main() {

	card := newDeckFromFile("deck.txt")
	card.shuffle()
	card.print()
}
