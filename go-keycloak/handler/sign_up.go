package handler

import (
	"encoding/json"
	"net/http"

	"go-keycloak/keycloak"
)

type SignUpRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h *HandlerAPI) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUp SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := keycloak.CreateUserParams{
		Username:  signUp.Username,
		Email:     signUp.Email,
		Password:  signUp.Password,
		FirstName: signUp.FirstName,
		LastName:  signUp.LastName,
	}

	err = h.clientKC.CreateUser(r.Context(), params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
