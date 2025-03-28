package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TeluTrix/tarc-server/internal/controller"
	"github.com/TeluTrix/tarc-server/internal/middleware"
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

	mux.Handle("/api/version", middleware.LoggingMiddleware(http.HandlerFunc(controller.GetCurrentVersion)))

	slog.Info(fmt.Sprintf("Started tarc-server on port: %d", service.Conf.Server.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", service.Conf.Server.Port), mux)
}
