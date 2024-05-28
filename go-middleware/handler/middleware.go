package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type KeySession string

const UserID KeySession = "userID"

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// middleware ...
		h.ServeHTTP(w, r)
	})
}

func TrailingSlashRedirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			lastChart := r.URL.Path[len(r.URL.Path)-1]
			if lastChart == '/' {
				nPath := r.URL.Path[:len(r.URL.Path)-1]
				http.Redirect(w, r, nPath, http.StatusMovedPermanently)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

func IsAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token session ...
		ctx := context.WithValue(r.Context(), UserID, time.Now().UnixMilli())
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nReq := &http.Request{}
		*nReq = *r
		nReq.Header.Set("X-Request-Id", fmt.Sprintf("%d", time.Now().UnixMilli()))
		h.ServeHTTP(w, nReq)
	})
}

func ResponseServer(h http.Handler, servername string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", servername)
		h.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newCustomResponse(w)

		next.ServeHTTP(rw, r)

		log.Printf("method=%s path=%s status=%d duration=%s response=%s\n",
			r.Method,
			r.URL.Path,
			rw.status,
			time.Since(start),
			rw.body.String(),
		)
	})
}
