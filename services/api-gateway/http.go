package main

import (
	"encoding/json"
	"microservices_with_go/shared/contracts"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse json data", http.StatusBadRequest)
		return
	}

	if reqBody.UserID == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: "ok"}

	// TODO: Call trip service
	writeJson(w, http.StatusCreated, response)
}
