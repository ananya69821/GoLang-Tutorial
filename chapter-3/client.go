package main 

import (
	"fmt"
	"io/ioutil"
    "net/http"
)

func main () {
	resp, err := http.Get("https://swapi.dev/api/people/1")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body)) // turns the website into html page response
}
