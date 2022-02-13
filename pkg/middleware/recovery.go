package middleware

import (
	"net/http"

	"github.com/00security/profefe/pkg/log"
)

func RecoveryHandler(logger *log.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorw("panic serving request", "rid", RequestIDFromContext(req.Context()), "uri", req.RequestURI, "panic", err)
			}
		}()
		handler.ServeHTTP(w, req)
	})
}
