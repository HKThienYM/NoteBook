package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	services "thien.com/server/Services"
)

const (
	port = ":50051"
)

func main() {
	fmt.Println("Hello world!")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	services.RegisterService(server)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
