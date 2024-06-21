package main

import (
	"log"
	"net"

	"github.com/matizaj/grpc-start-oms/common/services/orders/handler"
	"github.com/matizaj/grpc-start-oms/common/services/orders/service"
	"google.golang.org/grpc"
)
type gRPCServer struct {
	addr string
}

func NewGrpcServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Panic("tcp listener failed", err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	// register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersHandler(grpcServer, orderService)

	log.Printf("Grpc Server is up and running on  port %s \n", s.addr)
	return grpcServer.Serve(lis)
}