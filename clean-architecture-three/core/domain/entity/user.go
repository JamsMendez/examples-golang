package entity

import "errors"

type User struct {
	ID       uint
	Username string
	Password string
	Name     string
	LastName string
}

var (
	ErrUserNotFound = errors.New("User not found")
	ErrUserInvalid = errors.New("User invalid")
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService {
		userRepository,
	}
}

func (u *userService) Find(ID uint) (*User, error) {
	return u.userRepository.Find(ID)
}

func (u *userService) Create(user *User) error {
	// validate user ...
	// update user input or create new *User

	return u.userRepository.Create(user)
}
