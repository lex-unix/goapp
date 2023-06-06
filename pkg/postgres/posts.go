package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lexunix/goapp/pkg/domain"
)

type PostService struct {
	DB *pgxpool.Pool
}

func (s *PostService) Get(id int64) (*domain.Post, error) {
	var p domain.Post
	row := s.DB.QueryRow(context.Background(), "select id, title, body from post where id = $1", id)
	if err := row.Scan(&p.ID, &p.Title, &p.Body); err != nil {
		return nil, fmt.Errorf("PostService.Get(): %v", err)
	}

	return &p, nil
}

func (s *PostService) Create(p *domain.Post) error {
	_, err := s.DB.Exec(context.Background(), "insert into post (title, body) values ($1, $2)", p.Title, p.Body)
	return err
}

func (s *PostService) Delete(id int64) error {
	_, err := s.DB.Exec(context.Background(), "delete from post where id = $1", id)
	return err
}
