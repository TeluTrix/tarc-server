package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/model"
	"github.com/TeluTrix/tarc-server/internal/service"
)

const v = "0.1.0"

func GetCurrentVersion(resp http.ResponseWriter, req *http.Request) {
	var version model.Version
	version.Version = v

	json, err := json.Marshal(version)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		slog.Error(err.Error())
	}

	service.ReturnReponse(resp, http.StatusOK, json)
}
