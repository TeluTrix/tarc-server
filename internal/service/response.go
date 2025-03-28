package service

import (
	"net/http"
)

func ReturnReponse(resp http.ResponseWriter, status int, responseObject []byte) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write(responseObject)
}
