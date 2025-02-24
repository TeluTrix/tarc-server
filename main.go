package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/service"
)

type Version struct {
	Version string
}

const v = "0.1.0"

func init() {
	service.ReadConfig()
	service.ConnectToDB(service.Conf)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/version", returnVersion)

	http.ListenAndServe(fmt.Sprintf(":%d", service.Conf.Server.Port), mux)
}

func returnVersion(resp http.ResponseWriter, req *http.Request) {
	var version Version
	version.Version = v

	json, err := json.Marshal(version)
	if err != nil {
		slog.Error(err.Error())
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(json)
}
