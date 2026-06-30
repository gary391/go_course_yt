package main

import "fmt"

func main() {

	number_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for int_val := range number_slice {
		if int_val%2 == 0 {
			fmt.Printf("even:  %v\n", int_val)
		} else {
			fmt.Printf("odd :  %v\n", int_val)
		}
		//

	}
}
