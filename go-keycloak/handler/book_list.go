package handler

import (
	"fmt"
	"net/http"
)

func (h *HandlerAPI) BookList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "get list book")
}
