package response

import (
	"encoding/json"
	"net/http"
)

type baseResponse struct {
	Result any `json:"result,omitempty"`
}

func newResponseWithData(d any) baseResponse {
	return baseResponse{
		Result: d,
	}
}

func ResponsdWithData(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(newResponseWithData(data))
}
