package mocks

import (
	"github.com/boliev/fnl-news/internal/source"
	"github.com/stretchr/testify/mock"
)

// SourceMock mock for source struct
type SourceMock struct {
	source.Source
	mock.Mock
}

// Domain mock function for source struct
func (s *SourceMock) Domain() string {
	args := s.Called()
	return args.String(0)
}

// Path mock function for source struct
func (s *SourceMock) Path() string {
	args := s.Called()
	return args.String(0)
}

// Name mock function for source struct
func (s *SourceMock) Name() string {
	args := s.Called()
	return args.String(0)
}

// ArticleListItemPattern mock function for source struct
func (s *SourceMock) ArticleListItemPattern() string {
	args := s.Called()
	return args.String(0)
}

// ArticlePattern mock function for source struct
func (s *SourceMock) ArticlePattern() string {
	args := s.Called()
	return args.String(0)
}

// ArticleListItemTitlePattern mock function for source struct
func (s *SourceMock) ArticleListItemTitlePattern() string {
	args := s.Called()
	return args.String(0)
}

// ArticleListItemHrefPattern mock function for source struct
func (s *SourceMock) ArticleListItemHrefPattern() string {
	args := s.Called()
	return args.String(0)
}

// ArticleFullTitle mock function for source struct
func (s *SourceMock) ArticleFullTitle(title string) string {
	args := s.Called(title)
	return args.String(0)
}

// ArticleFullURL mock function for source struct
func (s *SourceMock) ArticleFullURL(url string) string {
	args := s.Called(url)
	return args.String(0)
}

// ShouldSkipArticle mock function for source struct
func (s *SourceMock) ShouldSkipArticle(url string) bool {
	args := s.Called(url)
	return args.Bool(0)
}
