package repository

import (
	"github.com/ansidev/gin-starter-project/domain/author"
	"github.com/ansidev/gin-starter-project/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewPostgresAuthorRepository(db *gorm.DB) author.IAuthorRepository {
	return &postgresAuthorRepository{db}
}

type postgresAuthorRepository struct {
	db *gorm.DB
}

func (r *postgresAuthorRepository) GetByID(id int64) (author.Author, error) {
	var a author.Author
	result := r.db.First(&a, id)

	if result.Error != nil {
		log.Errorz("Error while querying author by id ", zap.Error(result.Error))
		return author.Author{}, result.Error
	}

	return a, nil
}
