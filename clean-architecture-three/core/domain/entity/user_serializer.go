package entity

type UserSerializer interface {
  Decode(input []byte) (*User, error) 
  Encode(user *User) ([]byte, error)
}
