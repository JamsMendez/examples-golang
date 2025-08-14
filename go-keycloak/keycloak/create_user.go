package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type CreateUserParams struct {
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (c *ClientKeycloak) CreateUser(ctx context.Context, params CreateUserParams) error {
	jwt, err := c.kc.LoginAdmin(ctx, c.userAdmin, c.pwdAdmin, c.realmAdmin)
	if err != nil {
		return err
	}

	nUser := gocloak.User{
		Username:  gocloak.StringP(params.Username),
		Email:     gocloak.StringP(params.Email),
		FirstName: gocloak.StringP(params.FirstName),
		LastName:  gocloak.StringP(params.LastName),
		Enabled:   gocloak.BoolP(true),
	}

	var userID string
	userID, err = c.kc.CreateUser(ctx, jwt.AccessToken, c.realm, nUser)
	if err != nil {
		return err
	}

	err = c.kc.SetPassword(ctx, jwt.AccessToken, userID, c.realm, params.Password, false)
	if err != nil {
		return err
	}

	return nil
}
