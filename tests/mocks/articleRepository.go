package mocks

import (
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/stretchr/testify/mock"
)

// ArticleRepositoryMock mock for source struct
type ArticleRepositoryMock struct {
	repository.Article
	mock.Mock
}

// GetNewTg returns new articles for telegram publisher
func (r *ArticleRepositoryMock) GetNewTg() []*domain.Article {
	args := r.Called()
	return args.Get(0).([]*domain.Article)
}

// MarkAsSentTG returns new articles for telegram publisher
func (r *ArticleRepositoryMock) MarkAsSentTG(article *domain.Article) {
	r.Called(article)
}
