package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mathesukkj/rinha-de-backend-2025/src/integration"
	"github.com/mathesukkj/rinha-de-backend-2025/src/models"
)

func HandleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, ErrInvalidRequest, http.StatusBadRequest)

		return
	}

	if err := integration.CreatePayment(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HandlePaymentsSummary(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	from := query.Get("from")
	to := query.Get("to")
	if from == "" || to == "" {
		http.Error(w, ErrInvalidRequest, http.StatusBadRequest)
		return
	}

	_, err := time.Parse(time.RFC3339, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse(time.RFC3339, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	summary := models.PaymentsSummaryResponse{
		Default: models.Default{
			TotalRequests: 0,
			TotalAmount:   0,
		},
		Fallback: models.Fallback{
			TotalRequests: 0,
			TotalAmount:   0,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(summary)
}
