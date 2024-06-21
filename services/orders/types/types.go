package types

import (
	"context"
	orders "github.com/matizaj/grpc-start-oms/common/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) ([]*orders.Order, error)
}