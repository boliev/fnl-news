package publisher

import (
	"github.com/boliev/fnl-news/internal/repository"
)

// GetPublishers returns publishers list
func GetPublishers(
	articleRepository *repository.ArticleRepository,
	telegramConfig *TelegramPublisherConfig,
) []Publisher {
	var publishers []Publisher

	telegramPublisher := NewTelegramPublisher(
		articleRepository,
		telegramConfig,
	)
	publishers = append(publishers, telegramPublisher)

	return publishers
}
