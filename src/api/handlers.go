package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mathesukkj/rinha-de-backend-2025/src/integration"
	"github.com/mathesukkj/rinha-de-backend-2025/src/models"
)

func (api *API) HandleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if request.Amount <= 0 {
		http.Error(w, ErrInvalidRequest, http.StatusBadRequest)

		return
	}

	request.RequestedAt = time.Now()

	if err := integration.CreatePayment(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := api.db.CreatePayment(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func HandlePaymentsSummary(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	from := query.Get("from")
// 	to := query.Get("to")

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(summary)
// }
