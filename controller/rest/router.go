package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GabrielBG99/vxchan/logger"
	"github.com/GabrielBG99/vxchan/service"
	"github.com/go-chi/chi"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	log := logger.GetLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := NewLogResponseWriter(w)

		start := time.Now()
		next.ServeHTTP(lw, r)
		elapsed := time.Since(start).Milliseconds()

		log.Infow(fmt.Sprintf("[%s %d] %s - %dms", r.Method, lw.StatusCode, r.RequestURI, elapsed),
			"elapsed", elapsed,
			"method", r.Method,
			"uri", r.RequestURI,
			"status", lw.StatusCode,
		)
	})
}

func NewRouter(svc service.Service) http.Handler {
	router := chi.NewRouter()

	router.Use(LoggerMiddleware)

	router.Route("/api", func(api chi.Router) {
		api.Get("/", v1ListBoards(svc))
		api.Get("/{boardID}", v1ListThreads(svc))
		api.Post("/{boardID}", v1CreateThread(svc))
		api.Post("/{boardID}/{threadID}", v1ReplyThread(svc))
	})

	return router
}
