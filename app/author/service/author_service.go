package service

import "github.com/ansidev/gin-starter-project/domain/author"

type IAuthorService interface {
	GetByID(id int64) (author.Author, error)
}

func NewAuthorService(authorRepository author.IAuthorRepository) IAuthorService {
	return &AuthorService{authorRepository}
}

type AuthorService struct {
	authorRepository author.IAuthorRepository
}

func (s *AuthorService) GetByID(id int64) (author.Author, error) {
	return s.authorRepository.GetByID(id)
}
