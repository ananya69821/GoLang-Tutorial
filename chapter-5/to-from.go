package main 

import (
	"io/ioutil"
	"log"
)

func main() {
	initial, err := ioutil.ReadFile("text/myfile.txt")
	if err != nil {
		log.Fatal(err)
	} 

	ioutil.WriteFile("text/Adina.txt", initial, 0644)

	if err != nil {
		log.Fatal(err)
	}

}