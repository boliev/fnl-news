package mocks

import (
	"github.com/boliev/fnl-news/internal/source"
	"github.com/stretchr/testify/mock"
)

type SourceMock struct {
	source.Source
	mock.Mock
}

func (s *SourceMock) Domain() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) Path() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) Name() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) ArticleListItemPattern() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) ArticlePattern() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) ArticleListItemTitlePattern() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) ArticleListItemHrefPattern() string {
	args := s.Called()
	return args.String(0)
}

func (s *SourceMock) ArticleFullTitle(title string) string {
	args := s.Called(title)
	return args.String(0)
}

func (s *SourceMock) ArticleFullURL(url string) string {
	args := s.Called(url)
	return args.String(0)
}

func (s *SourceMock) ShouldSkipArticle(url string) bool {
	args := s.Called(url)
	return args.Bool(0)
}
