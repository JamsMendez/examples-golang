package handler

import "go-keycloak/keycloak"

type HandlerAPI struct {
	clientKC *keycloak.ClientKeycloak
}

func NewHandlerAPI(client *keycloak.ClientKeycloak) *HandlerAPI {
	return &HandlerAPI{
		clientKC: client,
	}
}
