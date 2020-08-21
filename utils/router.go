package utils

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Recoverer ...
func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("X-Sessionid")
		trackingID := r.Header.Get("X-Trackingid")
		if r.URL != nil {
			sessionParam := r.URL.Query()["session"]
			if len(sessionParam) > 0 {
				sessionID = sessionParam[0]
				r.Header.Set("X-Sessionid", sessionID)
			}
			trackingParam := r.URL.Query()["tracking"]
			if len(trackingParam) > 0 {
				trackingID = trackingParam[0]
				r.Header.Set("X-Trackingid", trackingID)
			}
		}
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// newRouter returns new router with preinstalled middleware
func newRouter() *chi.Mux {
	chi.RegisterMethod("COPY")
	chi.RegisterMethod("LOCK")
	chi.RegisterMethod("MKCOL")
	chi.RegisterMethod("MOVE")
	chi.RegisterMethod("PROPFIND")
	chi.RegisterMethod("PROPPATCH")
	chi.RegisterMethod("UNLOCK")
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		recoverer,
		middleware.NoCache,
	)
	return router
}

var routerInitOnce sync.Once
var router *chi.Mux

// Router returns map feature services route interface
func Router() *chi.Mux {
	routerInitOnce.Do(func() {
		router = newRouter()
	})
	return router
}
