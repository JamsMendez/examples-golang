package keycloak

import (
	"context"
	"fmt"
)

type UserInfo struct {
	Username string
	Email    string
}

func (c *ClientKeycloak) UserInfo(ctx context.Context, accessToken string) (*UserInfo, error) {
	user, err := c.kc.GetUserInfo(ctx, accessToken, c.realm)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user info not found")
	}

	userInfo := &UserInfo{}

	if user.Nickname != nil {
		userInfo.Username = *user.Nickname
	}

	if user.Email != nil {
		userInfo.Email = *user.Email
	}

	return userInfo, nil
}
