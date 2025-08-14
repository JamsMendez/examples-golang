package handler

import (
	"log"
	"net/http"

	"go-keycloak/keycloak"
)

type Middleware struct {
	client *keycloak.ClientKeycloak
}

func NewMiddleware(client *keycloak.ClientKeycloak) *Middleware {
	return &Middleware{
		client: client,
	}
}

func (m Middleware) IsAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Access-Token")
		if accessToken == "" {
			http.Error(w, "access token not found", http.StatusUnauthorized)
			return
		}

		userInfo, err := m.client.UserInfo(r.Context(), accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Printf("%+v\n", userInfo)

		h.ServeHTTP(w, r)
	})
}

// func middleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		h.ServeHTTP(w, r)
// 	})
// }
