package api

import "net/http"

const (
	CreatePaymentsPath  = "POST /payments"
	PaymentsSummaryPath = "GET /payments-summary"
)

type Router struct {
	router *http.ServeMux
}

func NewRouter(api *API) *http.ServeMux {
	router := http.NewServeMux()
	RegisterRoutes(router, api)

	return router
}

func RegisterRoutes(router *http.ServeMux, api *API) {
	router.HandleFunc(CreatePaymentsPath, api.HandleCreatePayment)
	// router.HandleFunc(PaymentsSummaryPath, HandlePaymentsSummary)
}
