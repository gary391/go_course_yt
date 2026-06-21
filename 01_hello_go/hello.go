// The main function has to be in the package main.
package main 

// import is used to include other packages in our program. 
// Here we are importing the fmt package,
//  which provides functions for formatted I/O.
import (
	"fmt"
)
// The main package is the starting point of a Go program.
func main() {
	const name, age = "Matt", 666
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.

}
