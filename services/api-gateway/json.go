package main

import (
	"encoding/json"
	"net/http"
)

// TODO: later to migrate to Shared cuz other services will use this

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
