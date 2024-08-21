package main

// import "fmt"

func main() {
	// var cards string = "Ace of Spades"
	// cards := "Ace of Spades" 		// Short hand declaration
	// cards = "Five of Diamonds" // Reassigning the value of the variable
	// cards := newCard()
	
	my_slices := deck{"one", "two", newCard()}

	my_slices = append(my_slices, "three") // Append to the slice


	my_slices.print()

	// fmt.Println(my_slices)

}

func newCard() string {
	return "Five of Diamonds"
}


