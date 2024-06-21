package main

import "log"

func main() {
	gRPCServer := NewGrpcServer(":9000")
	httpServer := NewHttpServer(":8000")


	if err := httpServer.Run(); err != nil {
		log.Panic("http cant start", err)
	}

	if err := gRPCServer.Run(); err != nil {
		log.Panic("grpc cant start", err)
	}
}