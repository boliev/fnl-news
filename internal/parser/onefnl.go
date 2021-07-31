package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// Onefnl parser struct
type Onefnl struct {
	Domain string
	Path   string
}

// NewOnefnlParser creates Onefnl parser
func NewOnefnlParser(config Config) *Onefnl {
	return &Onefnl{
		Domain: config.Domain,
		Path:   config.Path,
	}
}

// GetName return name of the parser
func (p Onefnl) GetName() string {
	return "onefnl"
}

// Parse parses articles from onefnl
func (p Onefnl) Parse() ([]*domain.Article, error) {
	var parsedArticles []*domain.Article
	articles, err := p.getArticlesList()
	if err != nil {
		return parsedArticles, err
	}

	for _, v := range articles {
		article, err := p.getArticle(v)
		if err != nil {
			//логгер?
			log.Warn(err.Error())
			continue
		}
		parsedArticles = append(parsedArticles, article)
	}
	return parsedArticles, nil
}

func (p Onefnl) getArticle(item articlesListItem) (*domain.Article, error) {
	article := &domain.Article{
		Title:  item.title,
		Href:   item.href,
		TgSent: false,
	}

	articlePage, err := p.getArticlePage(item)
	if err != nil {
		return nil, err
	}

	article.HTML = p.getField("<div class=\"news-item-page\">(.*?)</div>", articlePage)
	article.ImageURL = p.Domain + p.getField("src=\"(.*?)\"", article.HTML)
	article.Date = p.getField("<span class=\"date\">(.*?)</span>", articlePage)

	matcher := NewTagMatcher()
	article.Tags = strings.Join(matcher.MatchTags(article.Title+" "+article.HTML), " #")
	if article.Tags != "" {
		article.Tags = "#" + article.Tags
	}

	return article, nil
}

func (p Onefnl) getArticlesList() ([]articlesListItem, error) {
	var articles []articlesListItem
	body, err := p.getListPage()
	if err != nil {
		return articles, fmt.Errorf("sorry cant get %s list page: %s", p.GetName(), err)
	}

	newsListRegexp := regexp.MustCompile("(?msi)<div class=\"news-info\">(.*?)</a>")
	newsList := newsListRegexp.FindAllStringSubmatch(body, -1)

	for _, item := range newsList {
		title, href := p.getTitleAndHref(item[1])
		articles = append(articles, articlesListItem{
			title: title,
			href:  p.Domain + p.Path + href,
		})
	}

	return articles, nil
}

func (p Onefnl) getTitleAndHref(text string) (string, string) {
	return p.getField("class=\"news-title\">(.*)", text), p.getField("href=\"(.*?)\"", text)
}

func (p Onefnl) getField(pattern string, text string) string {
	fieldRegexp := regexp.MustCompile(fmt.Sprintf("(?msi)%s", pattern))
	field := fieldRegexp.FindStringSubmatch(text)
	if len(field) < 2 {
		return ""
	}
	return field[1]
}

func (p Onefnl) getListPage() (string, error) {
	client := resty.New()

	url := p.Domain + p.Path
	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (p Onefnl) getArticlePage(item articlesListItem) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(item.href)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
