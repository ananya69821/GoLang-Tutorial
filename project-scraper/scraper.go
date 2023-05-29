package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	response, err := http.Get("http://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print each book title
	document.Find(".product_pod h3 a").Each(func(index int, element *goquery.Selection) {
		title, _ := element.Attr("title")
		fmt.Println(title)
	})
}
