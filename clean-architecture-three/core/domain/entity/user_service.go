package entity

type UserService interface {
  Find(ID uint) (*User, error)
  Create(user *User) error
}
