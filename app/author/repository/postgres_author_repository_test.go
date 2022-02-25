package repository

import (
	"github.com/ansidev/gin-starter-project/constant"
	"github.com/ansidev/gin-starter-project/domain/author"
	"github.com/ansidev/gin-starter-project/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPostgresAuthorRepository(t *testing.T) {
	suite.Run(t, new(PostgresAuthorRepositoryTestSuite))
}

type PostgresAuthorRepositoryTestSuite struct {
	test.PostgresRepositoryTestSuite
	repository author.IAuthorRepository
}

func (s *PostgresAuthorRepositoryTestSuite) SetupSuite() {
	s.PostgresRepositoryTestSuite.SetupSuite()
	s.repository = NewPostgresAuthorRepository(s.Db)
}

func (s *PostgresAuthorRepositoryTestSuite) TestGetByID_ShouldReturnNoRecord() {
	_, err := s.repository.GetByID(1)

	require.Error(s.T(), err)
	require.Equal(s.T(), "record not found", err.Error())
}

func (s *PostgresAuthorRepositoryTestSuite) TestGetByID_ShouldReturnRecord() {
	_, err1 := s.SqlDb.Exec(`INSERT INTO author (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`,
		1,
		"John Doe",
		"2022-02-22 01:23:45",
		"2022-02-22 12:34:56")

	require.NoError(s.T(), err1)

	a, err2 := s.repository.GetByID(int64(1))

	require.NoError(s.T(), err2)
	require.Equal(s.T(), int64(1), a.ID)
	require.Equal(s.T(), "John Doe", a.Name)
	require.Equal(s.T(), "2022-02-22 01:23:45", a.CreatedAt.Format(constant.DateTimeFormat))
	require.Equal(s.T(), "2022-02-22 12:34:56", a.UpdatedAt.Format(constant.DateTimeFormat))
}
