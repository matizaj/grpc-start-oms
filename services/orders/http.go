package main

import (
	"log"
	"net/http"

	"github.com/matizaj/grpc-start-oms/common/services/orders/handler"
	"github.com/matizaj/grpc-start-oms/common/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (s *httpServer) Run() error {
	mux := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpHandler(orderService)
	orderHandler.RegisterRoutes(mux)

	log.Println("HTTP server is up and running on port ", s.addr)
	return http.ListenAndServe(s.addr, mux)
}