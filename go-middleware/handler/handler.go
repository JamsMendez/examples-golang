package handler

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) ListUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserID)
	fmt.Println("UserID: ", userID)

	ID := r.Header.Get("X-Request-Id")
	fmt.Println("ID: ", ID)

	w.WriteHeader(http.StatusOK)
	buffer := []byte("List User [] JSON")
	_, _ = w.Write(buffer)
}
