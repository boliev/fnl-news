package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/charmap"
	"regexp"
	"strings"
)

// Kulichki parser struct
type Kulichki struct {
	Domain string
	Path   string
}

// newKulichkiParser creates Onefnl parser
func newKulichkiParser(config Config) *Kulichki {
	return &Kulichki{
		Domain: config.Domain,
		Path:   config.Path,
	}
}

// GetName return name of the parser
func (p Kulichki) GetName() string {
	return "kulichki"
}

// Parse parses articles from onefnl
func (p Kulichki) Parse() ([]*domain.Article, error) {
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

func (p Kulichki) getArticle(item articlesListItem) (*domain.Article, error) {
	article := &domain.Article{
		Title:  item.title,
		Href:   item.href,
		TgSent: false,
	}

	articlePage, err := p.getArticlePage(item)
	if err != nil {
		return nil, err
	}
	decoder := charmap.Windows1251.NewDecoder()
	articlePage, err = decoder.String(articlePage)

	articlePage = p.getField("<!-- vk1 -->(.*?)</div>", articlePage)
	if articlePage == "\n\n</center>\r\n</td></tr>\r\n\r\n</table>\r\n" {
		return nil, nil
	}
	matcher := NewTagMatcher()
	article.Tags = strings.Join(matcher.GetAllTags(), " #")
	if article.Tags != "" {
		article.Tags = "#" + article.Tags
	}

	return article, nil
}

func (p Kulichki) getArticlesList() ([]articlesListItem, error) {
	var articles []articlesListItem
	body, err := p.getListPage()
	if err != nil {
		return articles, fmt.Errorf("sorry cant get %s list page: %s", p.GetName(), err)
	}
	decoder := charmap.Windows1251.NewDecoder()
	body, err = decoder.String(body)
	if err != nil {
		return articles, fmt.Errorf("sorry cant decode from window-1251")
	}
	body = p.getField("<font size=\"2\">Туры: </font>(.*?)</td>", body)

	newsListRegexp := regexp.MustCompile("(?msi)<a(.*?)</a>")
	newsList := newsListRegexp.FindAllStringSubmatch(body, -1)

	for _, item := range newsList {
		title, href := p.getTitleAndHref(item[1])
		articles = append(articles, articlesListItem{
			title: title,
			href:  p.Domain + href,
		})
	}

	return articles, nil
}

func (p Kulichki) getTitleAndHref(text string) (string, string) {
	return "ФНЛ на куличках. Обзор тура " + p.getField("href=\".*?\">(.*)", text), p.getField("href=\"(.*?)\"", text)
}

func (p Kulichki) getField(pattern string, text string) string {
	fieldRegexp := regexp.MustCompile(fmt.Sprintf("(?msi)%s", pattern))
	field := fieldRegexp.FindStringSubmatch(text)
	if len(field) < 2 {
		return ""
	}
	return field[1]
}

func (p Kulichki) getListPage() (string, error) {
	client := resty.New()

	url := p.Domain + p.Path
	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (p Kulichki) getArticlePage(item articlesListItem) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(item.href)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
