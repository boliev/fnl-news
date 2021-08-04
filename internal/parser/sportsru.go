package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// Sportsru parser struct
type Sportsru struct {
	Domain string
	Path   string
}

func newSportsruParser(config Config) *Sportsru {
	return &Sportsru{
		Domain: config.Domain,
		Path:   config.Path,
	}
}

// GetName return name of the parser
func (p Sportsru) GetName() string {
	return "sportbox"
}

// Parse articles from Sportbox site
func (p Sportsru) Parse() ([]*domain.Article, error) {
	var parsedArticles []*domain.Article
	articles, err := p.getArticlesList()
	if err != nil {
		return parsedArticles, err
	}

	ch := make(chan *domain.Article)
	chError := make(chan error)
	for _, v := range articles {
		go func(v articlesListItem) {
			article, err := p.getArticle(v)
			if err != nil {
				chError <- err
			} else {
				ch <- article
			}
		}(v)
	}

	for i := 0; i < len(articles); i++ {
		select {
		case err := <-chError:
			log.Warn(err.Error())
		default:
			parsedArticles = append(parsedArticles, <-ch)
		}
	}
	return parsedArticles, nil
}

func (p Sportsru) getArticle(item articlesListItem) (*domain.Article, error) {
	article := &domain.Article{
		Title:  item.title,
		Href:   item.href,
		TgSent: false,
	}

	articlePage, err := p.getArticlePage(item)
	if err != nil {
		return nil, err
	}

	article.HTML = p.getField("<div class=\"material-item__content js-mediator-article\" itemprop=\"articleBody\">(.*?)<footer class=\"material-item__footer\">", articlePage)

	//TODO вынести в di
	matcher := NewTagMatcher()
	article.Tags = strings.Join(matcher.MatchTags(article.Title+" "+article.HTML), " #")
	if article.Tags != "" {
		article.Tags = "#" + article.Tags
	}

	return article, nil
}

func (p Sportsru) getField(pattern string, text string) string {
	fieldRegexp := regexp.MustCompile(fmt.Sprintf("(?msi)%s", pattern))
	field := fieldRegexp.FindStringSubmatch(text)
	if len(field) < 2 {
		return ""
	}
	return field[1]
}

func (p Sportsru) getArticlesList() ([]articlesListItem, error) {
	var articles []articlesListItem
	body, err := p.getListPage()
	if err != nil {
		return articles, fmt.Errorf("sorry cant get %s list page: %s", p.GetName(), err)
	}

	newsListRegexp := regexp.MustCompile("(?msi)<h2 class=\"titleH2\">(.*?)</h2>")
	newsList := newsListRegexp.FindAllStringSubmatch(body, -1)
	for _, item := range newsList {
		title, href := p.getTitleAndHref(item[1])
		articles = append(articles, articlesListItem{
			title: title,
			href:  href,
		})
	}

	return articles, nil
}

func (p Sportsru) getListPage() (string, error) {
	client := resty.New()

	url := p.Domain + p.Path
	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (p Sportsru) getArticlePage(item articlesListItem) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(item.href)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (p Sportsru) getTitleAndHref(text string) (string, string) {
	return p.getField("<a href=\".*?\">(.*?)</a>", text), p.getField("<a href=\"(.*?)\">.*?</a>", text)
}
