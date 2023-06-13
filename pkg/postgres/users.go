package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lexunix/goapp/pkg/domain"
)

type UserService struct {
	DB *pgxpool.Pool
}

func (s *UserService) User(username string) (*domain.User, error) {
	var user domain.User

	row := s.DB.QueryRow(context.Background(), "select id, username, email, password from account where username = $1", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("UserService.User(): %v", err)
	}

	return &user, nil
}

func (s *UserService) Get(id int64) (*domain.User, error) {
	var user domain.User
	row := s.DB.QueryRow(context.Background(), "select id, username, email, password from account where id = $1", id)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("PostService.Get(): %v", err)
	}
	return &user, nil
}

func (s *UserService) Create(u *domain.User) error {
	_, err := s.DB.Exec(context.Background(), "insert into account (username, email, password) values ($1, $2, $3)", u.Username, u.Email, u.Password)
	return err
}

func (s *UserService) Delete(id int64) error {
	_, err := s.DB.Exec(context.Background(), "delete from account where id = $1", id)
	return err
}
