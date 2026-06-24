package main

import "fmt"

/*
Create a a new type of deck
which is a slice of strings
*/

// Define a new type which is extending an existing type.
type deck []string
// receiver on a function 
 
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
