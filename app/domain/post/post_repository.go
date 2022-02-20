package post

type IPostRepository interface {
	GetByID(id int64) (Post, error)
}
