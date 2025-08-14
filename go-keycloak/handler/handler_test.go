package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-keycloak/keycloak"
)

func TestSignUpAndSignIn(t *testing.T) {
	client := keycloak.NewClientKeycloak()

	h := NewHandlerAPI(client)
	m := NewMiddleware(client)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/sign_up", h.SignUp)
	mux.HandleFunc("POST /api/v1/sign_in", h.SignIn)

	muxWithAuth := http.NewServeMux()
	muxWithAuth.HandleFunc("GET /api/v1/book/list", h.BookList)
	muxIsAuth := m.IsAuth(muxWithAuth)

	server := httptest.NewServer(mux)
	serverAuth := httptest.NewServer(muxIsAuth)

	c := NewClient(server.URL, "")

	// t.Run("user sign up", func(t *testing.T) {
	// 	var body bytes.Buffer

	// 	req := SignUpRequest{
	// 		Username:  "JamsMendez",
	// 		Email:     "jamsmendez@email.com",
	// 		Password:  "123456",
	// 		FirstName: "Jams",
	// 		LastName:  "Mendez",
	// 	}

	// 	err := json.NewEncoder(&body).Encode(req)
	// 	if err != nil {
	// 		t.Fatalf("expected json encode request sign up error nil, got %v", err)
	// 	}

	// 	resp, err := c.makeRequest(http.MethodPost, "/api/v1/sign_up", &body)
	// 	if err != nil {
	// 		t.Fatalf("expected request sign up error nil, got %v", err)
	// 	}

	// 	if resp.StatusCode != http.StatusOK {
	// 		t.Fatalf("expected status http ok, got %d", resp.StatusCode)
	// 	}
	// })

	var accessToken, refreshToken string

	t.Run("user sign in", func(t *testing.T) {
		var body bytes.Buffer

		req := SignInRequest{
			Username: "JamsMendez",
			Password: "123456",
		}

		err := json.NewEncoder(&body).Encode(req)
		if err != nil {
			t.Fatalf("expected json encode request sign in error nil, got %v", err)
		}

		resp, err := c.makeRequest(http.MethodPost, "/api/v1/sign_in", &body)
		if err != nil {
			t.Fatalf("expected request sign in error nil, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status http ok, got %d", resp.StatusCode)
		}

		var tokenResp SignInResponse

		err = json.NewDecoder(resp.Body).Decode(&tokenResp)
		if err != nil {
			t.Fatalf("expected json decode request sign in error nil, got %v", err)
		}

		accessToken = tokenResp.AccessToken
		refreshToken = tokenResp.RefreshToken

		log.Println(accessToken, "\n", refreshToken)
	})

	t.Run("is auth middleware with access-token", func(t *testing.T) {
		c.baseURL = serverAuth.URL
		c.apiKey = accessToken

		resp, err := c.makeRequest(http.MethodGet, "/api/v1/book/list", nil)
		if err != nil {
			t.Fatalf("expected request get list book error nil, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected get list book status http ok, got %d", resp.StatusCode)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("expected get list book read response, %v", err)
		}

		log.Println(resp.StatusCode, string(b))
	})

	t.Run("is auth middleware without access-token", func(t *testing.T) {
		c.baseURL = serverAuth.URL
		c.apiKey = "access token fake"

		resp, err := c.makeRequest(http.MethodGet, "/api/v1/book/list", nil)
		if err != nil {
			t.Fatalf("expected request get list book error nil, got %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Fatalf("expected get list book status http unauthorized, got %d", resp.StatusCode)
		}

		log.Println(resp.StatusCode)
	})
}
