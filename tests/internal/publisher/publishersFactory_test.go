package parser

import (
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishersFactorySuccess(t *testing.T) {
	articleRepositoryMock := new(mocks.ArticleRepositoryMock)
	config := &publisher.TelegramPublisherConfig{
		ChatID: "chat_id",
		Token:  "token",
	}
	client := new(mocks.ClientMock)
	publishers := publisher.GetPublishers(articleRepositoryMock, config, client)
	assert.NotEmpty(t, publishers)
}
