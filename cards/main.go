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

	cardsSlice := []string{"8 of Spades", "Queen of Hearts", newCard()}
	fmt.Println(cardsSlice)

	cardsSlice = append(cardsSlice, "Six of Diamonds")
	fmt.Println(cardsSlice)

	for i, card := range cardsSlice {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
