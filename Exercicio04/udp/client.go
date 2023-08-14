package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	udpServer, err := net.ResolveUDPAddr("udp", ":8082")

	if err != nil {
		fmt.Println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		fmt.Println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	for i := 0; i < 10000; i++ {

		_, err = conn.Write([]byte("Harry Potter"))
		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			fmt.Println("Read data failed:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(received))
	}
}
