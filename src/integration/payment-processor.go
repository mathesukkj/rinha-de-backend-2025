package integration

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mathesukkj/rinha-de-backend-2025/src/models"
)

func CreatePayment(request models.CreatePaymentRequest) error {
	json, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://payment-processor-default:8080/payments", "application/json", bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func HealthCheck() error {
	resp, err := http.Get("http://localhost:8001/service-health")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response models.ServiceHealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
