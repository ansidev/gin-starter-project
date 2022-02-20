package db

import (
	"database/sql"
	"fmt"
	"github.com/ansidev/gin-starter-project/pkg/log"
	_ "github.com/lib/pq"
)

const (
	PostgresDsnFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func NewPostgresClient(config SqlDbConfig) *sql.DB {
	address := fmt.Sprintf(
		PostgresDsnFormat,
		config.DbHost,
		config.DbPort,
		config.DbUsername,
		config.DbPassword,
		config.DbName,
	)
	sqlDb, err := sql.Open(config.DbDriver, address)

	log.FatalIf(err, "Failed to connect to database")

	return sqlDb
}
