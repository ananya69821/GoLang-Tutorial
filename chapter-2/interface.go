package main

import "fmt"

// structure
type Rect struct {
	width, height float64 
}

// interface of a shape 
type Shape interface {
    Area() float64
}

// method using shape
func (r Rect) Area() float64 {
	return r.width * r.height
}

// prints the flaots
func printArea(s Shape) {
    fmt.Println(s.Area())
}

// calling the method
func main() {
	r := Rect{width: 5, height: 4}
	printArea(r)
}