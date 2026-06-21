package main

import ("fmt"; "02_cmd/hello"; "os" )

func main() {
		fmt.Println(hello.Say(os.Args[1:]))
	}