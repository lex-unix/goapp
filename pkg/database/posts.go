package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lexunix/goapp/pkg/domain"
)

type PostService struct {
	DB *sqlx.DB
}

func (s *PostService) All() (*[]domain.Post, error) {
	posts := []domain.Post{}
	err := s.DB.Select(&posts, "SELECT * FROM Post")
	if err != nil {
		return nil, err
	}
	return &posts, nil

}

func (s *PostService) Get(id int64) (*domain.Post, error) {
	var p domain.Post
	if err := s.DB.Get(&p, "SELECT * FROM Post WHERE id = $1", id); err != nil {
		return nil, fmt.Errorf("PostService.Get(): %v", err)
	}
	return &p, nil
}

func (s *PostService) Create(p *domain.Post) error {
	_, err := s.DB.NamedExec(`INSERT INTO Post (title, body, userId) VALUES (:title, :body, :userid)`, p)
	return err
}

func (s *PostService) Delete(id int64) error {
	_, err := s.DB.Exec("DELETE FROM Post WHERE id = $1", id)
	return err
}

func (s *PostService) UserPosts(userID int64) (*[]domain.Post, error) {
	posts := []domain.Post{}
	err := s.DB.Select(&posts, "SELECT * FROM Post WHERE userid = $1", userID)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
