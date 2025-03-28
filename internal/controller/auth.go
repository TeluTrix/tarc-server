package controller

import (
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/TeluTrix/tarc-server/internal/service"
)

func RegisterNewAPIKey(resp http.ResponseWriter, req *http.Request) schema.ApiKey {
	var apiKey schema.ApiKey

	apiKey.ApiKey = schema.GenerateAPIKey(32)

	dbResult := service.DB.Create(&apiKey)

	if dbResult.Error != nil {
		service.ReturnReponse(resp, http.StatusBadRequest)
	}
}

func ValidateAPIKey() {

}

func QueryAPIKeys() {}
