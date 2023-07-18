package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/lexunix/goapp/pkg/domain"
)

type UserService struct {
	DB *sqlx.DB
}

func (s *UserService) User(username string) (*domain.User, error) {
	var user domain.User
	err := s.DB.Get(&user, "SELECT * FROM account WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Get(id int64) (*domain.User, error) {
	var user domain.User
	err := s.DB.Get(&user, "SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Create(u *domain.User) error {
	_, err := s.DB.NamedExec(`INSERT INTO account (username, email, password) VALUES (:username, :email, :password)`, u)
	return err
}

func (s *UserService) Delete(id int64) error {
	_, err := s.DB.Exec("DELETE FROM account WHERE id = $1", id)
	return err
}
