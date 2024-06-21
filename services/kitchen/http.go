package main

import (
	"html/template"
	"log"
	"net/http"

	orders "github.com/matizaj/grpc-start-oms/common/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (s *httpServer) Run() error {
	mux := http.NewServeMux()

	conn, err := NewGRPCClient(":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		client := orders.NewOrderServiceClient(conn)
		_, err := client.CreateOrder(r.Context(), &orders.CreateOrderRequest{
			CustomerId: 66,
			ProductId: 66,
			Quantity: 66,
		})

		if err != nil {
			log.Fatal("client error", err)
			return
		}
		o, err  := client.GetOrders(r.Context(), &orders.GetOrderRequest{CustomerID: 22})
		if err != nil {
			log.Fatal("client error", err)
			return
		}

		t:= template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, o); err != nil {
			log.Fatal("template error", err)
		}
	})

	log.Println("HTTP server is up and running on port ", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`