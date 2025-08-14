package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
)

type ClientKeycloak struct {
	kc *gocloak.GoCloak

	clientID     string
	clientSecret string
	realm        string

	userAdmin  string
	pwdAdmin   string
	realmAdmin string
}

func NewClientKeycloak() *ClientKeycloak {
	return &ClientKeycloak{
		kc: gocloak.NewClient("http://localhost:8080"),

		clientID:     "library-openid-auth",
		clientSecret: "ptTJzX9FPuCsMvBWMpZoLLoD6m02Vw8O",
		realm:        "go-services-realm",

		userAdmin:  "admin",
		pwdAdmin:   "admin",
		realmAdmin: "master",
	}
}
