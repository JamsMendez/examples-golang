package handler

import (
	"encoding/json"
	"net/http"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *HandlerAPI) SignIn(w http.ResponseWriter, r *http.Request) {
	var signIn SignInRequest

	err := json.NewDecoder(r.Body).Decode(&signIn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwt, err := h.clientKC.Login(r.Context(), signIn.Username, signIn.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	signInResp := SignInResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(signInResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
