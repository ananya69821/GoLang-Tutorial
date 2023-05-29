package main

import (
	"io/ioutil"
	"log"
)

func main() {
	content := []byte("Hello Adina")
	err := ioutil.WriteFile("text/Adina.txt", content, 0644)

	if err != nil {
		log.Fatal(err)
	}
}