package publisher

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/go-resty/resty/v2"
	"strings"
)

// TelegramPublisher struct
type TelegramPublisher struct {
	repository *repository.ArticleRepository
	chatID     string
	token      string
}

type requestBody struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// NewTelegramPublisher creates publisher
func NewTelegramPublisher(
	repository *repository.ArticleRepository,
	chatID string,
	token string,
) *TelegramPublisher {
	return &TelegramPublisher{
		repository: repository,
		chatID:     chatID,
		token:      token,
	}
}

// GetName returns publisher name
func (p TelegramPublisher) GetName() string {
	return "Telegram"
}

// PublishNew publishes new articles
func (p TelegramPublisher) PublishNew() {
	newArticles := p.repository.GetNewTg()
	for _, article := range newArticles {
		err := p.publishArticle(article)
		if err != nil {
			fmt.Println(fmt.Errorf("cant publish article: %s", err.Error()))
		}
	}
}

func (p TelegramPublisher) publishArticle(article *domain.Article) error {
	client := resty.New()
	body := requestBody{
		ChatID:    p.chatID,
		Text:      p.compileMessage(article),
		ParseMode: "Markdown",
	}

	res, err := client.R().
		SetBody(body).
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s/bot%s/%s", "https://api.telegram.org", p.token, "sendMessage"))
	if res != nil && res.StatusCode() > 299 {
		return fmt.Errorf("cant send message. Code: %d, response: %s. request: %s", res.StatusCode(), res.String(), res.Request.Body)
	}
	p.repository.MarkAsSentTG(article)

	return err
}

func (p TelegramPublisher) compileMessage(article *domain.Article) string {
	return fmt.Sprintf(
		"[%s](%s)\n\n%s",
		article.Title,
		article.Href,
		strings.Replace(article.Tags, "_", "\\_", -1),
	)
}
