package service

import (
	"context"
	orders "github.com/matizaj/grpc-start-oms/common/services/common/genproto/orders"
)

type OrderService struct {
	// db conn
}

var orderList []*orders.Order

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	orderList=append(orderList, order)
	return nil
}