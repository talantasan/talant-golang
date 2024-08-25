package main

import "fmt"


func main(){
	cards := newDeck()

	hand, remainingHand := deal(cards, 7)

	hand.print()
	fmt.Println("--------------------")
	remainingHand.print()
}