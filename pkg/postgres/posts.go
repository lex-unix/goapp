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
	_, err := s.DB.Exec(context.Background(), "insert into post (title, body, userId) values ($1, $2, 1)", p.Title, p.Body)
	return err
}

func (s *PostService) Delete(id int64) error {
	_, err := s.DB.Exec(context.Background(), "delete from post where id = $1", id)
	return err
}

func (s *PostService) UserPosts(userID int64) (*[]domain.Post, error) {
	rows, err := s.DB.Query(context.Background(), "select id, title, body, userid from post where userid = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		var p domain.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.UserID); err != nil {
			return &posts, err
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return &posts, err
	}

	return &posts, nil
}
