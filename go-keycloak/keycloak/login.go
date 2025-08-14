package keycloak

import "context"

type JWT struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

func (c *ClientKeycloak) Login(ctx context.Context, username, password string) (*JWT, error) {
	jwt, err := c.kc.Login(ctx, c.clientID, c.clientSecret, c.realm, username, password)
	if err != nil {
		return nil, err
	}

	cJWT := JWT {
		AccessToken: jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn: jwt.ExpiresIn,
	}

	return &cJWT, nil
}
