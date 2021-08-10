package repository

import "github.com/boliev/fnl-news/internal/domain"

// Article repository interface
type Article interface {
	Save(article *domain.Article)
	SaveAll(articles []*domain.Article)
	GetNewTg() []*domain.Article
	MarkAsSentTG(article *domain.Article)
}
