package main

import (
	"context"
	"database/sql"
	"fmt"
	authorController "github.com/ansidev/gin-starter-project/author/controller/http"
	"github.com/ansidev/gin-starter-project/config"
	"github.com/ansidev/gin-starter-project/constant"
	gormPkg "github.com/ansidev/gin-starter-project/pkg/gorm"
	"github.com/ansidev/gin-starter-project/pkg/log"
	postController "github.com/ansidev/gin-starter-project/post/controller/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	appEnv string
	sqlDb  *sql.DB
	gormDb *gorm.DB
)

func init() {
	log.InitLogger("console")

	appEnv = os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = constant.DefaultProdEnv
	}

	if appEnv == constant.DefaultProdEnv {
		config.LoadConfig("/app", constant.DefaultProdConfig, &config.AppConfig)
	} else {
		config.LoadConfig(".", constant.DefaultDevConfig, &config.AppConfig)
	}
}

func main() {
	// Flush log buffer if necessary
	defer log.Sync()

	router := gin.Default()
	if appEnv == constant.DefaultProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	// Default route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name":        constant.AppName,
			"version":     constant.AppVersion,
			"releaseDate": constant.AppReleaseDate,
		})
	})

	initInfrastructureServices()
	initControllers(router)

	server := &http.Server{
		Addr:    initAddress(),
		Handler: router,
	}

	// Listen from a different goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create channel to signify a signal being sent
	exit := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	<-exit
	log.Info("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

func initInfrastructureServices() {
	sqlDb = InitSqlClient(config.AppConfig.SqlDbConfig)
	dialector := postgres.New(postgres.Config{
		Conn:                 sqlDb,
		PreferSimpleProtocol: true,
	})
	gormDb = gormPkg.InitGormDb(dialector)
}

func initControllers(router *gin.Engine) {
	authorService := InitAuthorService(gormDb)
	authorController.NewAuthorController(router, authorService)

	postService := InitPostService(gormDb)
	postController.NewPostController(router, postService)
}

func initAddress() string {
	return fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port)
}
