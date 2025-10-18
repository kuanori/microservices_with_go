package main

import (
	"encoding/json"
	"log"
	"microservices_with_go/services/api-gateway/grpc_clients"
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

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	// Don't forget to close the client to avoid resource leaks
	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("failed to send preview trip: %v", err)
		http.Error(w, "Failed to preview trip", http.StatusInternalServerError)
		return
	}

	// ======== HTTP
	// jsonBody, _ := json.Marshal(reqBody)
	// reader := bytes.NewReader(jsonBody)

	// resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	// if err != nil {
	// 	http.Error(w, "failed to call trip service: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// defer resp.Body.Close()

	// var respBody any
	// if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
	// 	http.Error(w, "failed to parse json data from trip service", http.StatusBadRequest)
	// 	return
	// }

	// response := contracts.APIResponse{Data: respBody}

	response := contracts.APIResponse{Data: tripPreview}

	writeJson(w, http.StatusOK, response)
}
