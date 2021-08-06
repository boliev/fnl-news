package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/source"
	httpclient "github.com/boliev/fnl-news/pkg/httpClient"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// Parser struct
type Parser struct {
	source     source.Source
	tagMatcher *TagMatcher
	client     httpclient.Client
}

type articlesListItem struct {
	title string
	href  string
}

// NewParser Parser constructor
func NewParser(source source.Source, tagMatcher *TagMatcher, client httpclient.Client) *Parser {
	return &Parser{
		source:     source,
		tagMatcher: tagMatcher,
		client:     client,
	}
}

// Parse start parsing source
func (p Parser) Parse() ([]*domain.Article, error) {
	articles, err := p.getArticlesList()
	if err != nil {
		return nil, err
	}

	var parsedArticles []*domain.Article
	ch := make(chan *domain.Article)
	chError := make(chan error)
	chSkip := make(chan bool)
	for _, v := range articles {
		go func(v articlesListItem) {
			article, err := p.getArticle(v)
			if err != nil {
				chError <- err
			} else if article != nil {
				ch <- article
			} else {
				chSkip <- true
			}
		}(v)
	}

	for i := 0; i < len(articles); i++ {
		select {
		case err := <-chError:
			log.Warn(err.Error())
		case article := <-ch:
			parsedArticles = append(parsedArticles, article)
		case _ = <-chSkip:
			continue
		}
	}

	return parsedArticles, nil
}

func (p Parser) getArticlesList() ([]articlesListItem, error) {
	var articles []articlesListItem
	body, err := p.client.Get(p.source.Domain() + p.source.Path())
	if err != nil {
		return articles, fmt.Errorf("sorry cant get list page for %s: %s", p.source.Name(), err)
	}

	newsListRegexp := regexp.MustCompile("(?msi)" + p.source.ArticleListItemPattern())
	newsList := newsListRegexp.FindAllStringSubmatch(body, -1)

	for _, item := range newsList {
		title, href := p.getTitleAndHref(item[1])
		articles = append(articles, articlesListItem{
			title: p.source.ArticleFullTitle(title),
			href:  p.source.ArticleFullURL(href),
		})
	}

	return articles, nil
}

func (p Parser) getArticle(item articlesListItem) (*domain.Article, error) {
	article := &domain.Article{
		Title:  item.title,
		Href:   item.href,
		TgSent: false,
	}

	articlePage, err := p.client.Get(item.href)
	if err != nil {
		return nil, err
	}

	article.HTML = p.getField(p.source.ArticlePattern(), articlePage)

	if p.source.ShouldSkipArticle(article.HTML) {
		return nil, nil
	}

	article.Tags = strings.Join(p.tagMatcher.MatchTags(article.Title+" "+article.HTML), " #")
	if article.Tags != "" {
		article.Tags = "#" + article.Tags
	}

	return article, nil
}

func (p Parser) getField(pattern string, text string) string {
	fieldRegexp := regexp.MustCompile(fmt.Sprintf("(?msi)%s", pattern))
	field := fieldRegexp.FindStringSubmatch(text)
	if len(field) < 2 {
		return ""
	}
	return field[1]
}

func (p Parser) getTitleAndHref(text string) (string, string) {
	return p.getField(p.source.ArticleListItemTitlePattern(), text),
		p.getField(p.source.ArticleListItemHrefPattern(), text)
}
