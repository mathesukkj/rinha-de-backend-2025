package api

import "net/http"

const (
	CreatePaymentsPath  = "POST /payments"
	PaymentsSummaryPath = "GET /payments-summary"
)

type Router struct {
	router *http.ServeMux
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	RegisterRoutes(router)

	return router
}

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc(CreatePaymentsPath, HandleCreatePayment)
	// router.HandleFunc(PaymentsSummaryPath, HandlePaymentsSummary)
}
