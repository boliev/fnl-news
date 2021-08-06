package source

// Onefnl parser struct
type Onefnl struct {
	BaseSource
}

// NewOnefnlParser Onefnl source constructor
func NewOnefnlParser() *Onefnl {
	return &Onefnl{
		BaseSource: BaseSource{
			name:                        "onefnl",
			domain:                      "https://1fnl.ru",
			path:                        "/news/",
			articlePattern:              "<div class=\"news-item-page\">(.*?)</div>",
			articleListItemPattern:      "<div class=\"news-info\">(.*?)</a>",
			articleListItemTitlePattern: "class=\"news-title\">(.*)",
			articleListItemHrefPattern:  "href=\"(.*?)\"",
		},
	}
}

// ArticleFullURL returns full article url from relative link
func (s Onefnl) ArticleFullURL(relative string) string {
	return s.Domain() + s.Path() + relative
}

// ArticleFullTitle returns full article title for the source
func (s Onefnl) ArticleFullTitle(title string) string {
	return title
}

// ShouldSkipArticle validates if the text is correct and we should add it to tDB
func (s Onefnl) ShouldSkipArticle(text string) bool {
	return false
}
