package http

import (
	"github.com/ansidev/gin-starter-project/author/service"
	"github.com/ansidev/gin-starter-project/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func registerRoutes(router *gin.Engine, authorController *AuthorController) {
	v1 := router.Group("/author/v1")

	v1.GET("/authors/:id", authorController.GetAuthor)
}

func NewAuthorController(router *gin.Engine, authorService service.IAuthorService) {
	controller := &AuthorController{authorService}
	registerRoutes(router, controller)
}

type AuthorController struct {
	authorService service.IAuthorService
}

func (ctrl *AuthorController) GetAuthor(ctx *gin.Context) {
	authorIdParam := ctx.Param("id")

	authorId, err := strconv.ParseInt(authorIdParam, 10, 64)

	if err != nil {
		log.Errorz("Invalid author id", zap.String("author_id", authorIdParam))
	}

	author, err := ctrl.authorService.GetByID(authorId)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, author)
}
