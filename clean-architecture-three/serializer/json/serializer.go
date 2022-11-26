package json

import (
	"clean-architecture-three/core/domain/entity"
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {}

func (u *User) Decode(input []byte) (*entity.User, error) {
  user := &entity.User{}

  if err := json.Unmarshal(input, user); err != nil {
    msg := fmt.Sprintf("%s %s", err.Error(), "serializer.User.Decode")
    return nil, errors.New(msg)
  }

  return user, nil
}

func (u *User) Encode(input *entity.User) ([]byte, error) {
	buffer, err := json.Marshal(input)
	if err != nil {
		msg := fmt.Sprintf("%s %s", err.Error(), "serializer.User.Encode")
		return nil, errors.New(msg)
	}

	return buffer, nil
}
