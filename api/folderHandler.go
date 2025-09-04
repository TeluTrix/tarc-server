package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/TeluTrix/tarc-server/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFolder(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var folder schema.Folder
	err := decoder.Decode(&folder)
	if err != nil {
		slog.Error(err.Error())
	}

	returnFolder, err := service.CreateFolder(folder)
	if err != nil {
		slog.Error(err.Error())
		service.WriteResponse(resp, nil, http.StatusInternalServerError)
	}

	service.WriteResponse(resp, returnFolder, http.StatusCreated)

}
func GetFolder(resp http.ResponseWriter, req *http.Request) {
	id, _ := uuid.Parse(req.FormValue("id"))

	folder, err := service.GetFolder(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error(err.Error())
		service.WriteResponse(resp, nil, http.StatusInternalServerError)
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error(err.Error())
		service.WriteResponse(resp, struct{}{}, http.StatusNotFound)
	}

	service.WriteResponse(resp, folder, http.StatusOK)
}
func GetFolders(resp http.ResponseWriter, req *http.Request) {
	folders, err := service.GetFolders()
	if err != nil {
		service.WriteResponse(resp, nil, http.StatusInternalServerError)
	}

	service.WriteResponse(resp, folders, http.StatusOK)
}
func UpdateFolder(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var folder schema.Folder
	err := decoder.Decode(&folder)
	if err != nil {
		slog.Error(err.Error())
	}
}
func DeleteFolder(resp http.ResponseWriter, req *http.Request) {}
