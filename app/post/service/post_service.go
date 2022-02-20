package service

import "github.com/ansidev/gin-starter-project/domain/post"

type IPostService interface {
	GetByID(id int64) (post.Post, error)
}

func NewPostService(postRepository post.IPostRepository) IPostService {
	return &PostService{postRepository}
}

type PostService struct {
	postRepository post.IPostRepository
}

func (s *PostService) GetByID(id int64) (post.Post, error) {
	return s.postRepository.GetByID(id)
}
