package main 

import ( 
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "Anonymous", "The username to be greeted.")
	email := flag.String("email", "no-email-provided", "The email of the user.")

	flag.Parse()

	fmt.Printf("Hello, %s! We've sent a greeting to your email: %s\n", *username, *email)
}
// Run this: go build greet.go
// Run this: ./greet -username=Navid -email=navid_khan@icloud.com