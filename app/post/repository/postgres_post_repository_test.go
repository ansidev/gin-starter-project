package repository

import (
	"github.com/ansidev/gin-starter-project/constant"
	"github.com/ansidev/gin-starter-project/domain/post"
	"github.com/ansidev/gin-starter-project/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPostgresPostRepository(t *testing.T) {
	suite.Run(t, new(PostgresPostRepositoryTestSuite))
}

type PostgresPostRepositoryTestSuite struct {
	test.PostgresRepositoryTestSuite
	repository post.IPostRepository
}

func (s *PostgresPostRepositoryTestSuite) SetupSuite() {
	s.PostgresRepositoryTestSuite.SetupSuite()
	s.repository = NewPostgresPostRepository(s.Db)
}

func (s *PostgresPostRepositoryTestSuite) TestGetByID_ShouldReturnNoRecord() {
	_, err := s.repository.GetByID(int64(1))

	require.Error(s.T(), err)
	require.Equal(s.T(), "record not found", err.Error())
}

func (s *PostgresPostRepositoryTestSuite) TestGetByID_ShouldReturnRecord() {
	_, err1 := s.SqlDb.Exec(`INSERT INTO author (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`,
		1,
		"John Doe",
		"2022-02-22 01:23:45",
		"2022-02-22 12:34:56")

	require.NoError(s.T(), err1)

	_, err2 := s.SqlDb.Exec(`INSERT INTO post (id, title, content, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		1,
		"Sample title",
		"Sample content",
		1,
		"2022-02-22 01:23:45",
		"2022-02-22 12:34:56")

	require.NoError(s.T(), err2)

	p, err3 := s.repository.GetByID(int64(1))

	require.NoError(s.T(), err3)
	require.Equal(s.T(), int64(1), p.ID)
	require.Equal(s.T(), "Sample title", p.Title)
	require.Equal(s.T(), "Sample content", p.Content)
	require.Equal(s.T(), int64(1), p.AuthorID)
	require.Equal(s.T(), "2022-02-22 01:23:45", p.CreatedAt.Format(constant.DateTimeFormat))
	require.Equal(s.T(), "2022-02-22 12:34:56", p.UpdatedAt.Format(constant.DateTimeFormat))
}
