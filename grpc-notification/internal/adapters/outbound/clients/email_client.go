package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/KKogaa/grpc-notification/internal/core/entities"
)

type EmailClient struct {
	endpoint string
	token    string
}

func NewEmailClient(endpoint string, token string) *EmailClient {
	return &EmailClient{
		endpoint: endpoint,
		token:    token,
	}
}

func (e *EmailClient) SendEmail(email entities.Email) (entities.Email, error) {

	payload := map[string]interface{}{
		"to":             email.To,
		"subject":        email.Subject,
		"transaction_id": email.TransactionId,
		"amount":         email.Amount,
		"description":    email.Description,
		"status":         email.Status,
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", e.endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Workflow-Api-Key", e.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return email, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Failed to send email: %s", resp.Status)
		return email, errors.New("send failed")
	}

	return email, nil
}
