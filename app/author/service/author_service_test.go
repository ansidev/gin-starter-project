package service

import (
	"errors"
	"github.com/ansidev/gin-starter-project/author/mock"
	"github.com/ansidev/gin-starter-project/domain/author"
	"github.com/ansidev/gin-starter-project/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestAuthorService(t *testing.T) {
	suite.Run(t, new(AuthorServiceTestSuite))
}

type AuthorServiceTestSuite struct {
	test.ServiceTestSuite
	mockAuthorRepository *mock.MockIAuthorRepository
}

func (s *AuthorServiceTestSuite) SetupSuite() {
	s.ServiceTestSuite.SetupSuite()
	s.mockAuthorRepository = mock.NewMockIAuthorRepository(s.Ctrl)
}

func (s *AuthorServiceTestSuite) TestGetByID_ShouldReturnNoRecord() {
	a1 := author.Author{}

	s.mockAuthorRepository.
		EXPECT().
		GetByID(int64(1)).
		Return(a1, errors.New("record not found"))

	authorService := NewAuthorService(s.mockAuthorRepository)
	a2, err := authorService.GetByID(int64(1))

	require.Equal(s.T(), a1, a2)
	require.Error(s.T(), err)
	require.Equal(s.T(), "record not found", err.Error())
}

func (s *AuthorServiceTestSuite) TestGetByID_ShouldReturnOneRecord() {
	a1 := author.Author{
		ID:        1,
		Name:      "John Doe",
		CreatedAt: time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt: time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}

	s.mockAuthorRepository.
		EXPECT().
		GetByID(int64(1)).
		Return(a1, nil)

	authorService := NewAuthorService(s.mockAuthorRepository)
	a2, err := authorService.GetByID(int64(1))

	require.Equal(s.T(), a1, a2)
	require.NoError(s.T(), err)
}
