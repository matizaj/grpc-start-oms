package main

import "log"

func main() {
	gRPCServer := NewGrpcServer(":9000")
	
	if err := gRPCServer.Run(); err != nil {
		log.Panic("grpc cant start")
	}
}