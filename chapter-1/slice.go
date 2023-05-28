package main 

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}  // declare an array
	s := arr[1:3]                  // declare a slice that views arr[1], arr[2]
	fmt.Println(s)                 // prints "[2 3]"
}