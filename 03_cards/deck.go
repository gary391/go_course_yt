package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Here we have variable that we are not using so we can swap these with
	// underscore
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// Here we are appending the cancatenated value back in the cards list.
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

// receiver on a function
// This will allow us to call print method on deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function takes two argument.
// deck is an argument
// return two values each of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

// deck to string
// receiver of a deck
// can call -> cards.toString()
// here d is the receiver i.e. the deck we are opearting on.
func (d deck) toString() string {
	// slice of strings
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Here we don't have a deck yet so we don't need a receiver.
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	// check the value of err
	if err != nil {
		// Option  #1 - log the error and return a call to newDeck()
		// Option  #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1) // exit if something went wrong code value is non-zero
	}
	
	fmt.Println(string(bs))
	s := strings.Split(string(bs), ",") // slice of strings
	return deck(s) // here deck is an extension of slice of string to support type conversatio

}
// Supports inplace shuffling

func (d deck) shuffle() {
	// seed value
	// using time -->unixnano
	source := rand.NewSource(time.Now().UnixNano())// int64 return value
	// New rand object
	r :=rand.New(source) // source should be int64


	for i := range d {
		// The pseudo random generator depends a seed value
		// by default we are using the same seed value 
		newPosition := r.Intn(len(d)-1) // this will return an int value
		// swap element in the 
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
