package http

import (
	"github.com/ansidev/gin-starter-project/pkg/log"
	"github.com/ansidev/gin-starter-project/post/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func registerRoutes(router *gin.Engine, postController *PostController) {
	v1 := router.Group("/post/v1")

	v1.GET("/posts/:id", postController.GetPost)
}

func NewPostController(router *gin.Engine, postService service.IPostService) {
	controller := &PostController{postService}
	registerRoutes(router, controller)
}

type PostController struct {
	postService service.IPostService
}

func (ctrl *PostController) GetPost(ctx *gin.Context) {
	postIdParam := ctx.Param("id")

	postId, err := strconv.ParseInt(postIdParam, 10, 64)

	if err != nil {
		log.Errorz("Invalid post id", zap.String("post_id", postIdParam))
	}

	post, err := ctrl.postService.GetByID(postId)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, post)
}
