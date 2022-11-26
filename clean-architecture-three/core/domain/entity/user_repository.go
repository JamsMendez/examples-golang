package entity

type UserRepository interface {
  Find(ID uint) (*User, error) 
  Create(user *User) error
}
