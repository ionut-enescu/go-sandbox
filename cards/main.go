package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println(card)
	card = "King of Kings"
	fmt.Println(card)
	card = newCard()
	fmt.Println(card)

	cardsSlice := deck{"8 of Spades", "Queen of Hearts", newCard()}
	fmt.Println(cardsSlice)

	cardsSlice = append(cardsSlice, "Six of Diamonds")
	fmt.Println(cardsSlice)

	cardsSlice.printDeck()

	newDeck().printDeck()

	newDealedDeck, remainingCardsDeck := deal(newDeck(), 3)
	fmt.Println("-----newDealedDeck---------")
	newDealedDeck.printDeck()
	fmt.Println("-----remainingCardsDeck---------")
	remainingCardsDeck.printDeck()

	fmt.Println("-----toString()---------")
	fmt.Println(newDeck().toString())

	fmt.Println("-----saveToFile()---------")
	newDeck().saveToFile("MyCards")

	fmt.Println("-----newDeckFromFile()---------")
	newDeckFromFile("MyCards").printDeck()
	newDeckFromFile("MyCard").printDeck()

}

func newCard() string {
	return "Five of Diamonds"
}
