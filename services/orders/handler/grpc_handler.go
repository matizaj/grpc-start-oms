package handler

import (
	"context"
	"log"

	orders "github.com/matizaj/grpc-start-oms/common/services/common/genproto/orders"
	"github.com/matizaj/grpc-start-oms/common/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// svc injection
	orderService types.OrderService
	// unimplemented
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersHandler(grpc *grpc.Server,  orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{orderService: orderService, }
	log.Println(gRPCHandler)
	// register OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, r *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId: 42,
		CustomerId: 42,
		ProductId: 42,
		Quantity: 42,
	}

	if err := h.orderService.CreateOrder(ctx, order);err != nil {
		log.Println("failed to create order ", err)
		return nil, err
	}

	res := &orders.CreateOrderResponse{Status: "order created!"}
	return res, nil
}