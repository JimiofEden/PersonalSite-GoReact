package utils

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, data models.ApiResponse) {
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ReponseWithError(w http.ResponseWriter, status int, data models.ApiResponse) {
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}