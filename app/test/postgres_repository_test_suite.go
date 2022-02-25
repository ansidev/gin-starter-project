package test

import (
	"database/sql"
	"fmt"
	"github.com/ansidev/gin-starter-project/pkg/db"
	gormPkg "github.com/ansidev/gin-starter-project/pkg/gorm"
	"github.com/ansidev/gin-starter-project/pkg/log"
	ep "github.com/fergusstrange/embedded-postgres"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	waitTime = 10
	maxTry   = 5
)

func init() {
	log.InitLogger("console")
}

type PostgresRepositoryTestSuite struct {
	suite.Suite
	TestDb *ep.EmbeddedPostgres
	SqlDb  *sql.DB
	Db     *gorm.DB
}

func (s *PostgresRepositoryTestSuite) initDbConnection(dbConfig db.SqlDbConfig) {
	s.SqlDb = db.NewPostgresClient(dbConfig)
	dialector := postgres.New(postgres.Config{
		Conn:                 s.SqlDb,
		PreferSimpleProtocol: true,
	})
	s.Db = gormPkg.InitGormDb(dialector)
}

func (s *PostgresRepositoryTestSuite) SetupSuite() {
	dbConfig, testDb := GetTestDbConfig()

	s.TestDb = testDb

	err1 := s.TestDb.Start()

	// Workaround solution to auto start embed Postgres server on failed
	count := 0
	for err1 != nil && count <= maxTry {
		log.Error(err1)
		time.Sleep(time.Duration(waitTime*(count+1)) * time.Second)
		err1 = s.TestDb.Start()
		count++
	}

	require.NoError(s.T(), err1)

	s.initDbConnection(dbConfig)
}

func (s *PostgresRepositoryTestSuite) BeforeTest(suite string, method string) {
	log.Info(fmt.Sprintf("Suite: %s, Before running %s", suite, method))
	err := goose.Up(s.SqlDb, "../../migration")
	require.NoError(s.T(), err)
}

func (s *PostgresRepositoryTestSuite) AfterTest(suite string, method string) {
	log.Info(fmt.Sprintf("Suite: %s, After running %s", suite, method))
	err := goose.Reset(s.SqlDb, "../../migration")
	require.NoError(s.T(), err)
}

func (s *PostgresRepositoryTestSuite) TearDownSuite() {
	err := s.TestDb.Stop()
	require.NoError(s.T(), err)
}
