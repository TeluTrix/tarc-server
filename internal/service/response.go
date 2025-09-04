package service

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, v any, status int) {
	data, err := json.Marshal(v)
	if err != nil && v != nil {
		slog.Error("Failed to marshal JSON response", "error", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
