package main

import (
	"fmt"
	"impl"
	"log"
	"net"
	"net/rpc"
)

func server() {

	// cria instância da bookstore
	bookstore := impl.NewBookstore()

	// cria um novo consumer RPC e registra a livraria
	server := rpc.NewServer()
	err := server.RegisterName("Bookstore", bookstore)

	// cria um listener TCP
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer func(ln net.Listener) {
		var err = ln.Close()
		if err != nil {
			log.Fatal("listen error:", err)
		}
	}(ln)

	// aguarda por invocações
	fmt.Println("Servidor está pronto (RPC-TCP)...")
	server.Accept(ln)

}

func main() {

	go server()

	_, _ = fmt.Scanln()
}
