package app

import (
	"encoding/json"
	"net/http"
)

// Response : response message for API
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

// NewResponse creates a new response for returning data
func newResponse(object interface{}, message string, status int) *Response {
	return &Response{
		Data:    object,
		Message: message,
		Status:  status,
	}
}

// HandleJSONResponse : to output json response
func (res *Response) HandleJSONResponse() http.HandlerFunc {
	data, err := json.Marshal(res)
	if err != nil {
		return ThrowError(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}
}

// ThrowError :
func ThrowError(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
