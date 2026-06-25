package main

import "fmt"

/*
Create a a new type of deck
which is a slice of strings
*/

// Define a new type which is extending an existing type.
// Here the deck is a slice of strings
type deck []string

// when ever someone call this function newDeck we will return a variable of type
// deck
func newDeck() deck {
	// Create a new deck empty
	cards := deck{}

	// Slice of string we are using string and deck
	//
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues:= []string{"Ace","Two", "Three", "Four"}

	// Here we have variable that we are not using so we can swap these with 
	// underscore
	for _, suit:= range cardSuits {
		for _, value:= range cardValues {
			// Here we are appending the cancatenated value back in the cards list.
			cards = append(cards, value+" of "+suit) 

			}
		}
	return cards
	}

// receiver on a function 
 
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function takes two argument.
// return two values each of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}