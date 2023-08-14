package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	// create books variable to receive response
	var books string

	for i := 0; i < 10000; i++ {
		// get books with keywords
		if err := client.Call("Bookstore.FindBooks", "Harry Potter", &books); err != nil {
			fmt.Println("Error: Bookstore.FindBooks()", err)
		} else {
			fmt.Println(books)
		}
	}
}
