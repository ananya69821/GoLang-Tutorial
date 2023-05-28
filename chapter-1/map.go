package main 

import "fmt"

func main() {
	m := make(map[string]int)  // declare a map with string keys and int values
	m["hello"] = 1             // set "hello" = 1
	m["world"] = 2             // set "world" = 2
	fmt.Println(m)             // prints "map[hello:1 world:2]"
	
}