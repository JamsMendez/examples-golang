package mysql

import (
	"database/sql"
	"time"

	"clean-architecture-three/core/domain/entity"
)

type mysqlRepository struct {
	conn     *sql.DB
	database string
	timeout  time.Duration
}

func NewMySQL(url string) (*mysqlRepository, error) {
	conn, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return &mysqlRepository{conn: conn}, nil
}

func (m *mysqlRepository) Find(ID uint) (*entity.User, error) {
	// sql query ...

	user := &entity.User{
		ID:       1,
		Username: "JamsMendez",
	}

	return user, nil
}

func (m *mysqlRepository) Create(user *entity.User) error {
	// sql create ...
	user = &entity.User{
		ID:       2,
		Username: "JamsMendez",
		Password: "**********",
		Name:     "Jams",
		LastName: "Mendez",
	}

	return nil
}
