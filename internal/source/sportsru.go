package source

// Sportsru parser struct
type Sportsru struct {
	BaseSource
}

// NewSportsruParser Sportsru source constructor
func NewSportsruParser() *Sportsru {
	return &Sportsru{
		BaseSource: BaseSource{
			name:                        "sportsru",
			domain:                      "https://www.sports.ru",
			path:                        "/fnl/",
			articlePattern:              "<div class=\"material-item__content js-mediator-article\" itemprop=\"articleBody\">(.*?)<footer class=\"material-item__footer\">",
			articleListItemPattern:      "<h2 class=\"titleH2\">(.*?)</h2>",
			articleListItemTitlePattern: "<a href=\".*?\">(.*?)</a>",
			articleListItemHrefPattern:  "<a href=\"(.*?)\">.*?</a>",
		},
	}
}

// ArticleFullURL returns full article url from relative link
func (s Sportsru) ArticleFullURL(relative string) string {
	return relative
}

// ArticleFullTitle returns full article title for the source
func (s Sportsru) ArticleFullTitle(title string) string {
	return title
}

// ShouldSkipArticle validates if the text is correct and we should add it to tDB
func (s Sportsru) ShouldSkipArticle(text string) bool {
	return false
}
