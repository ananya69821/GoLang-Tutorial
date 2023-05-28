package main

import "fmt"

// type of structure
type Rectangle struct {
	width float64
	height float64
}

// method dealing with the structure
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// feeding its parameters
func main() {
	r := Rectangle{width: 5, height: 4}
	fmt.Println(r.area())
}