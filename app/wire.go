//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"database/sql"
	authorRepository "github.com/ansidev/gin-starter-project/author/repository"
	authorService "github.com/ansidev/gin-starter-project/author/service"
	"github.com/ansidev/gin-starter-project/pkg/db"
	postRepository "github.com/ansidev/gin-starter-project/post/repository"
	postService "github.com/ansidev/gin-starter-project/post/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitSqlClient(config db.SqlDbConfig) *sql.DB {
	wire.Build(db.NewPostgresClient)

	return &sql.DB{}
}

func InitAuthorService(db *gorm.DB) authorService.IAuthorService {
	wire.Build(authorService.NewAuthorService, authorRepository.NewPostgresAuthorRepository)
	return &authorService.AuthorService{}
}

func InitPostService(db *gorm.DB) postService.IPostService {
	wire.Build(postService.NewPostService, postRepository.NewPostgresPostRepository)
	return &postService.PostService{}
}
