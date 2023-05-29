package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	b := []byte(`{"Name":"Navid","Age":30,"Email":"navid_khan@icloud.com"}`)
	var p Person 

	err := json.Unmarshal(b, &p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}