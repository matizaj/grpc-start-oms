package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr=":7000"

func NewGRPCClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc failed to connect", err)
		return nil, err
	}

	defer conn.Close()
	return conn, nil
}

func main() {
	httpServer := NewHttpServer(addr)

	log.Println("kitchen server is up and running on port ", addr)

	if err := httpServer.Run(); err != nil {
		log.Fatal("kitche http sevrver failed")
	}
}