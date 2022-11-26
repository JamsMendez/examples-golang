package api

import (
	"clean-architecture-three/core/domain/entity"
	jsonS "clean-architecture-three/serializer/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userService entity.UserService
}

func NewHanlder(userService entity.UserService) UserHandler {
	return &handler{userService: userService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, status int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) entity.UserSerializer {
	return &jsonS.User{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}

  // parse param ID
	sID := r.URL.RawPath
	ID, err := strconv.Atoi(sID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	uID := uint(ID)
	user, err := h.userService.Find(uID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	body, err := h.serializer(contentType).Encode(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	setupResponse(w, contentType, body, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	user, err := h.serializer(contentType).Decode(body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.userService.Create(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	body, err = h.serializer(contentType).Encode(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	setupResponse(w, contentType, body, http.StatusCreated)
}
