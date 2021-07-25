package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/go-resty/resty/v2"
	"regexp"
	"strings"
)

// Sportbox parser struct
type Sportbox struct {
	Domain string
	Path   string
}

type articlesListItem struct {
	title string
	href  string
}

// NewSportboxParser creates Sportbox parser
func NewSportboxParser(domain string, path string) *Sportbox {
	return &Sportbox{
		Domain: domain,
		Path:   path,
	}
}

// GetName return name of the parser
func (s Sportbox) GetName() string {
	return "Sportbox"
}

// Parse articles from Sportbox site
func (s Sportbox) Parse() ([]*domain.Article, error) {
	var parsedArticles []*domain.Article
	articles, err := s.getArticlesList()
	if err != nil {
		return parsedArticles, err
	}

	for _, v := range articles {
		article, err := s.getArticle(v)
		if err != nil {
			// логгер?
			fmt.Println(err.Error())
			continue
		}
		parsedArticles = append(parsedArticles, article)
		fmt.Printf(
			"title: %s\nhref: %s\nimage: %s\ndate: %s\ntags: %s\n----------\n",
			article.Title, article.Href, article.ImageURL, article.Date, article.Tags,
		)
	}
	return parsedArticles, nil
}

func (s Sportbox) getArticle(item articlesListItem) (*domain.Article, error) {
	article := &domain.Article{
		Title:  item.title,
		Href:   item.href,
		IsSent: false,
	}

	articlePage, err := s.getArticlePage(item)
	if err != nil {
		return nil, err
	}

	article.ImageURL = s.getField("<img itemprop=\"image\" src=\"(.*?)\">", articlePage)
	article.Date = s.getField("<meta itemprop=\"dateCreated\" content=\"(.*?)\">", articlePage)
	article.HTML = s.getField("<div class=\"js-mediator-article\">(.*?)</div>", articlePage)

	matcher := NewTagMatcher()
	article.Tags = strings.Join(matcher.MatchTags(article.Title+" "+article.HTML), " #")
	if article.Tags != "" {
		article.Tags = "#" + article.Tags
	}

	return article, nil
}

func (s Sportbox) getField(pattern string, text string) string {
	fieldRegexp := regexp.MustCompile(fmt.Sprintf("(?msi)%s", pattern))
	field := fieldRegexp.FindStringSubmatch(text)
	if len(field) < 2 {
		return ""
	}
	return field[1]
}

func (s Sportbox) getArticlesList() ([]articlesListItem, error) {
	var articles []articlesListItem
	body, err := s.getListPage()
	if err != nil {
		return articles, fmt.Errorf("sorry cant get %s list page: %s", s.GetName(), err)
	}

	newsListRegexp := regexp.MustCompile("(?msi)<li><div class=\"_Sportbox_Spb2015_Components_TeazerBlock_TeazerBlock\">(.*?)</li>")
	newsList := newsListRegexp.FindAllStringSubmatch(body, -1)
	for _, item := range newsList {
		title, href := s.getTitleAndHref(item[1])
		articles = append(articles, articlesListItem{
			title: title,
			href:  s.Domain + href,
		})
	}

	return articles, nil
}

func (s Sportbox) getListPage() (string, error) {
	client := resty.New()

	url := s.Domain + s.Path
	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (s Sportbox) getArticlePage(item articlesListItem) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(item.href)

	if err != nil {
		return "", err
	}

	return s.cleanHTML(resp.String()), nil
}

func (s Sportbox) getTitleAndHref(text string) (string, string) {
	return s.getField("<span class=\"text\">(.*?)</span>", text), s.getField("href=\"(.*?)\"", text)
}

func (s Sportbox) cleanHTML(text string) string {
	text = regexp.MustCompile("(?msi)<a data-role-spb=\"spb_ig\"(.*?)>https://www\\.instagram\\.com/(.*?)</a>").
		ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<div class=\"spb-node-content-image-text\">(.*?)</div>").
		ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<div data-role-spb=\"spb_foto\" class=\"spb-node-content-image-wrapper\">(.*?)</div>").
		ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<a data-role-spb=\"spb_tw\"(.*?)>https://twitter\\.com/(.*?)<\a>").
		ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<a data-role-spb=\"spb_poster\"(.*?)>(.*?)</a>").
		ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<a(.*?)>").ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)</a>").ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)\n").ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<h3>Читайте также(.*)</ul>").ReplaceAllString(text, "")
	text = regexp.MustCompile("(?msi)<p><b>Следите за новостями и эфиром в (.*)").ReplaceAllString(text, "")

	return strings.TrimSpace(text)
}
