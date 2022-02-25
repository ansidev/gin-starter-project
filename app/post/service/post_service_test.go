package service

import (
	"errors"
	"github.com/ansidev/gin-starter-project/domain/author"
	"github.com/ansidev/gin-starter-project/domain/post"
	"github.com/ansidev/gin-starter-project/post/mock"
	"github.com/ansidev/gin-starter-project/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestPostService(t *testing.T) {
	suite.Run(t, new(PostServiceTestSuite))
}

type PostServiceTestSuite struct {
	test.ServiceTestSuite
	mockPostRepository *mock.MockIPostRepository
}

func (s *PostServiceTestSuite) SetupSuite() {
	s.ServiceTestSuite.SetupSuite()
	s.mockPostRepository = mock.NewMockIPostRepository(s.Ctrl)
}

func (s *PostServiceTestSuite) TestGetByID_ShouldReturnNoRecord() {
	p1 := post.Post{}

	s.mockPostRepository.
		EXPECT().
		GetByID(int64(1)).
		Return(p1, errors.New("record not found"))

	postService := NewPostService(s.mockPostRepository)
	p2, err := postService.GetByID(int64(1))

	require.Equal(s.T(), p1, p2)
	require.Error(s.T(), err)
	require.Equal(s.T(), "record not found", err.Error())
}

func (s *PostServiceTestSuite) TestGetByID_ShouldReturnOneRecord() {
	a1 := author.Author{
		ID:        1,
		Name:      "John Doe",
		CreatedAt: time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt: time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}

	p1 := post.Post{
		ID:        1,
		Title:     "Sample title",
		Content:   "Sample content",
		AuthorID:  a1.ID,
		Author:    a1,
		CreatedAt: time.Date(2022, 2, 22, 1, 23, 45, 0, time.UTC),
		UpdatedAt: time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}

	s.mockPostRepository.
		EXPECT().
		GetByID(int64(1)).
		Return(p1, nil)

	postService := NewPostService(s.mockPostRepository)
	p2, err := postService.GetByID(int64(1))

	require.Equal(s.T(), p1, p2)
	require.NoError(s.T(), err)
}
