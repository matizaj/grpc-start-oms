package handler

import (
	"net/http"

	"github.com/matizaj/grpc-start-oms/common/services/common"
	orders "github.com/matizaj/grpc-start-oms/common/services/common/genproto/orders"
	"github.com/matizaj/grpc-start-oms/common/services/orders/types"
)

type httpHandler struct {
	orderService types.OrderService
}

func NewHttpHandler(orderService types.OrderService) *httpHandler {
	return &httpHandler{orderService}
}

func (h *httpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *httpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var p orders.CreateOrderRequest
	if err := common.ReadJSON(r, &p); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	order := &orders.Order{
		OrderId: 22,
		CustomerId: p.GetCustomerId(),
		ProductId: p.GetProductId(),
		Quantity: p.GetQuantity(),
	}

	if err := h.orderService.CreateOrder(r.Context(), order);err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := &orders.CreateOrderResponse{Status: "success!!"}

	common.WriteJSON(w, http.StatusOK, res)

}