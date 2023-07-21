package domain

type Post struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}

type PostService interface {
	All() (*[]Post, error)
	Get(id int64) (*Post, error)
	Create(p *Post) error
	Delete(id int64) error
	UserPosts(userID int64) (*[]Post, error)
}
