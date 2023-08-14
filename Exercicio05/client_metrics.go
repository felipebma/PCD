package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	var times []int64
	if err != nil {
		log.Fatal(err)
	}

	// create books variable to receive response
	var books string

	for i := 0; i < 10000; i++ {
		start := time.Now()
		// get books with keywords
		if err := client.Call("Bookstore.FindBooks", "Harry Potter", &books); err != nil {
			fmt.Println("Error:1 College.Get()", err)
		} else {
			fmt.Println(books)
		}

		end := time.Now()
		times = append(times, end.Sub(start).Nanoseconds())
	}
	fmt.Fprintf(os.Stderr, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(times)), ","), "[]"))
}
