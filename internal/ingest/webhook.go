// Package ingest receives webhooks from mermaid-demo-payments.
package ingest

import (
	"encoding/json"
	"net/http"
)

type Envelope struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

// Handle is the single entry point for every upstream event.
func Handle(w http.ResponseWriter, r *http.Request) {
	var env Envelope
	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		http.Error(w, "bad envelope", http.StatusBadRequest)
		return
	}
	switch env.Type {
	case "charge.succeeded":
		enqueueReceipt(env)
	case "charge.failed":
		enqueueFailureNotice(env)
	}
	w.WriteHeader(http.StatusAccepted)
}

func enqueueReceipt(Envelope)       {}
func enqueueFailureNotice(Envelope) {}
