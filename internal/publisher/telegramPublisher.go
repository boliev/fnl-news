package publisher

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/repository"
	httpclient "github.com/boliev/fnl-news/pkg/httpClient"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// TelegramPublisher struct
type TelegramPublisher struct {
	repository repository.Article
	client     httpclient.Client
	chatID     string
	token      string
}

type requestBody struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// TelegramPublisherConfig config for telegram publisher
type TelegramPublisherConfig struct {
	ChatID string
	Token  string
}

// NewTelegramPublisher creates publisher
func NewTelegramPublisher(
	repository repository.Article,
	config *TelegramPublisherConfig,
	client httpclient.Client,
) *TelegramPublisher {
	return &TelegramPublisher{
		repository: repository,
		client:     client,
		chatID:     config.ChatID,
		token:      config.Token,
	}
}

// PublishNew publishes new articles
func (p TelegramPublisher) PublishNew() {
	newArticles := p.repository.GetNewTg()
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(newArticles))
	for _, article := range newArticles {
		go func(a *domain.Article, wg *sync.WaitGroup) {
			err := p.publishArticle(a)
			if err != nil {
				log.Warnf("cant publish article: %s", err.Error())
			}
			wg.Done()
		}(article, &waitGroup)
	}

	waitGroup.Wait()
}

func (p TelegramPublisher) publishArticle(article *domain.Article) error {
	body := requestBody{
		ChatID:    p.chatID,
		Text:      p.compileMessage(article),
		ParseMode: "Markdown",
	}

	url := fmt.Sprintf("%s/bot%s/%s", "https://api.telegram.org", p.token, "sendMessage")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	err := p.client.Post(url, body, headers)

	if err != nil {
		return fmt.Errorf("cant send message. %s", err.Error())
	}
	p.repository.MarkAsSentTG(article)

	return nil
}

func (p TelegramPublisher) compileMessage(article *domain.Article) string {
	return fmt.Sprintf(
		"[%s](%s)\n\n%s",
		article.Title,
		article.Href,
		strings.Replace(article.Tags, "_", "\\_", -1),
	)
}
