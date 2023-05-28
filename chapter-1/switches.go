package main 

import "fmt"

func main() {
	var x int = 5
	switch x {
	case 0: 
		fmt.Println("X is zero")
	case 1: 
		fmt.Println("X is one")
	default: 
		fmt.Println("X is more than that")
	}
	}