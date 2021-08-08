package publisher

import (
	"github.com/boliev/fnl-news/internal/repository"
	httpclient "github.com/boliev/fnl-news/pkg/httpClient"
)

// GetPublishers returns publishers list
func GetPublishers(
	articleRepository *repository.ArticleRepository,
	telegramConfig *TelegramPublisherConfig,
	client httpclient.Client,
) []Publisher {
	var publishers []Publisher

	telegramPublisher := NewTelegramPublisher(
		articleRepository,
		telegramConfig,
		client,
	)
	publishers = append(publishers, telegramPublisher)

	return publishers
}
