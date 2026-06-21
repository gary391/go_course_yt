package hello

import "strings"

// Variable and type declarations in Go are done using the var keyword, 
// followed by the variable name and type.
func Say(names [] string) string {
	if len(names) == 0 {
		names = [] string{"world"}
	}
	//func Join(elems []string, sep string) string
	return "Hello, " + strings.Join(names, ", ") + "!"
}