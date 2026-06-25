package main

func main() {

	// assign a variable and
	// declares and assigns a value
	// var card string = "Ace of Spades"
	// Here the go complier will figure out the card variable
	// := Lets you initialize and assign to the variable.
	// var card string
	// card = "Five of Diamands"
	// card := newCard()
	// will take string.
	// Here are are creaing and assigning the variable in a single operation
	/*
		var creates a variable
		card - The nsma of the variable
		string = string type
		Ace of Spades is the value of type string.

		- bool - true false
		- string - "Hi"
		- int - 0, -1000
		- float64 - 0.1
	*/
	// newCard can be defined in any part of the program
	// and can be called before defining it.
	// printState()
	// fmt.Println(newCard())
	// Slice of type string and not Slice of trings!!!
	// cards := deck{
	// 	"Ace of Hearts", newCard()}
	// Append function creates a new slice instead of modifying the existing one.
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	// Iterate over the slice
	// The reason we are using the := because for each iteration we are
	// initializing/assigning and throwing away the previous values.
	// every variable we declare in the code must be used in our code.
	// for i, card:= range cards {
	// 	fmt.Println(i,card)
	// }

	cards := newDeck()
	// cards.print()
	hand, remaingDecks := deal(cards, 5)
	hand.print()
	remaingDecks.print()

}

// Function should have a return type.
// What data type we are going to return from the function

// func newCard() string {
// 	return "Five of Diamonds"
// }
