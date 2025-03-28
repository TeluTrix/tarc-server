package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		slog.Info(fmt.Sprintf("Remote-Host: '%s', Endpoint: '%s'", req.RemoteAddr, req.RequestURI))
		next.ServeHTTP(w, req)
	})
}
