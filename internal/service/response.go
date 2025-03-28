package service

import (
	"fmt"
	"log/slog"
	"net/http"
)

func ReturnReponse(resp http.ResponseWriter, status int, responseObject []byte) {
	slog.Info(fmt.Sprintf("Response-Code:%d, Reponse-Body:%s", status, string(responseObject)))

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write(responseObject)
}
