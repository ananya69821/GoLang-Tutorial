package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println((i))
	}
}

func printLetters() {
    for i := 'a'; i < 'a'+10; i++ {
        fmt.Println(string(i))
    }
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(time.Second)  // sleep for 1 second
}