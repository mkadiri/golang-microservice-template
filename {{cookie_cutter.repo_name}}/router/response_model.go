package router

import (
	"encoding/json"
	"net/http"
)

type RestResponse struct {
	Message interface{}
	StatusCode int
}

func (resp *RestResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp.Message)
}