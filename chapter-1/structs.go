package main 

import "fmt"

// setting up the collection
type Point struct {
	X int 
	Y int
} 


// how to call the struct 
func main() {
	p := Point{4, 7}
	fmt.Println(p.X)
}
