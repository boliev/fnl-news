package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/tests/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTelegramPublisherSuccess(t *testing.T) {
	articleRepositoryMock := new(mocks.ArticleRepositoryMock)
	article1 := &domain.Article{Title: "article1", Href: "href1", Tags: "#Alania, #Torpedo", TgSent: false}
	article2 := &domain.Article{Title: "article2", Href: "href2", Tags: "#Alania", TgSent: false}
	articleRepositoryMock.On("GetNewTg").Return([]*domain.Article{article1, article2})
	articleRepositoryMock.On("MarkAsSentTG", mock.Anything).Return()
	config := &publisher.TelegramPublisherConfig{
		ChatID: "chat_id",
		Token:  "token",
	}
	clientMock := new(mocks.ClientMock)
	clientMock.On("Post", "https://api.telegram.org/bottoken/sendMessage", mock.Anything, mock.Anything).Return(nil)

	tgPublisher := publisher.NewTelegramPublisher(articleRepositoryMock, config, clientMock)
	tgPublisher.PublishNew()

	articleRepositoryMock.AssertNumberOfCalls(t, "MarkAsSentTG", 2)
}

func TestTelegramPublisherError(t *testing.T) {
	articleRepositoryMock := new(mocks.ArticleRepositoryMock)
	article1 := &domain.Article{Title: "article1", Href: "href1", Tags: "#Alania, #Torpedo", TgSent: false}
	article2 := &domain.Article{Title: "article2", Href: "href2", Tags: "#Alania", TgSent: false}
	articleRepositoryMock.On("GetNewTg").Return([]*domain.Article{article1, article2})
	articleRepositoryMock.On("MarkAsSentTG", mock.Anything).Return()
	config := &publisher.TelegramPublisherConfig{
		ChatID: "chat_id",
		Token:  "token",
	}
	clientMock := new(mocks.ClientMock)
	clientMock.On("Post", "https://api.telegram.org/bottoken/sendMessage", mock.Anything, mock.Anything).Return(fmt.Errorf("some error"))

	tgPublisher := publisher.NewTelegramPublisher(articleRepositoryMock, config, clientMock)
	tgPublisher.PublishNew()

	articleRepositoryMock.AssertNumberOfCalls(t, "MarkAsSentTG", 0)
}
