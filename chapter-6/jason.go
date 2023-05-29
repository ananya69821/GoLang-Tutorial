package main 

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string 
	Age int 
	Email string
}

func main() {
	p := Person{ Name: "Navid", Age: 30, Email: "Navid_Khan@icloud.com"}

	b, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

