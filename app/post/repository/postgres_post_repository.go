package repository

import (
	"github.com/ansidev/gin-starter-project/domain/post"
	"github.com/ansidev/gin-starter-project/pkg/log"
	"gorm.io/gorm"
)

func NewPostgresPostRepository(db *gorm.DB) post.IPostRepository {
	return &postgresPostRepository{db}
}

type postgresPostRepository struct {
	db *gorm.DB
}

func (r *postgresPostRepository) GetByID(id int64) (post.Post, error) {
	var p post.Post
	result := r.db.Preload("Author").First(&p, id)

	if result.Error != nil {
		log.Error("Error while querying post by id", result.Error)
		return post.Post{}, result.Error
	}

	return p, nil
}
