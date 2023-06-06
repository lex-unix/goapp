package domain

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostService interface {
	Get(id int64) (*Post, error)
	Create(p *Post) error
	Delete(id int64) error
}
