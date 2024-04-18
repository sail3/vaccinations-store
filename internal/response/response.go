package response

import (
	"encoding/json"
	"net/http"
)

type baseResponse struct {
	Result any `json:"result,omitempty"`
	Error  any `json:"error,omitempty"`
}

func newResponseWithData(d any) baseResponse {
	return baseResponse{
		Result: d,
	}
}

func newResponseWithError(e any) baseResponse {
	return baseResponse{
		Error: e,
	}
}

func ResponsdWithData(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(newResponseWithData(data))
}

func ResponseWithError(w http.ResponseWriter, errCode int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	return json.NewEncoder(w).Encode(newResponseWithError(err))
}
