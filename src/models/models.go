package models

import (
	"time"

	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	CorrelationID uuid.UUID `json:"correlation_id"`
	Amount        float64   `json:"amount"`
	RequestedAt   time.Time `json:"requested_at"`
}

type PaymentsSummaryResponse struct {
	Default  Default  `json:"default"`
	Fallback Fallback `json:"fallback"`
}

type Default struct {
	TotalRequests int `json:"total_requests"`
	TotalAmount   int `json:"total_amount"`
}

type Fallback struct {
	TotalRequests int `json:"total_requests"`
	TotalAmount   int `json:"total_amount"`
}

type ServiceHealthResponse struct {
	Failing         bool `json:"failing"`
	MinResponseTime int  `json:"min_response_time"`
}
