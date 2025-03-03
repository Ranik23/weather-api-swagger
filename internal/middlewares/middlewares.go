package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)


func PanicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Recovered from panic", slog.Any("err", err))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}


func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info(fmt.Sprintf("[%s] %s %s %s", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start)))
	})
}