package handler

import (
	"encoding/json"
	"net/http"
)

// Response structure for sending messages and data in JSON format
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JsonResponse sends a JSON response to the client
func JsonResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
