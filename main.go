package main

import (
	"fmt"
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/controller"
	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/TeluTrix/tarc-server/internal/service"
)

func init() {
	service.ReadConfig()
	service.ConnectToDB(service.Conf)
	service.DB.AutoMigrate(schema.Folder{})
	service.DB.AutoMigrate(schema.Item{})
	service.DB.AutoMigrate(schema.Tag{})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/version", controller.GetCurrentVersion)

	http.ListenAndServe(fmt.Sprintf(":%d", service.Conf.Server.Port), mux)
}
