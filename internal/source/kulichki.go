package source

// Kulichki parser struct
type Kulichki struct {
	BaseSource
}

// NewKulichkiParser Kulichki source constructor
func NewKulichkiParser() *Kulichki {
	return &Kulichki{
		BaseSource: BaseSource{
			name:                        "kulichki",
			domain:                      "https://football.kulichki.net",
			path:                        "/fnl/",
			articlePattern:              "<!-- vk1 -->(.*?)</div>",
			articleListItemPattern:      "(<a href=\"/fnl/202\\d{1,2}/\\d{1,2}/\">\\d{1,2}</a>)",
			articleListItemTitlePattern: "href=\".*?\">(.*)</a>",
			articleListItemHrefPattern:  "href=\"(.*?)\"",
		},
	}
}

// ArticleFullURL returns full article url from relative link
func (s Kulichki) ArticleFullURL(relative string) string {
	return s.Domain() + relative
}

// ArticleFullTitle returns full article title for the source
func (s Kulichki) ArticleFullTitle(title string) string {
	return "ФНЛ на куличках. Обзор тура №" + title
}

// ShouldSkipArticle validates if the text is correct and we should add it to tDB
func (s Kulichki) ShouldSkipArticle(text string) bool {
	if text == "\n\n</center>\r\n</td></tr>\r\n\r\n</table>\r\n" {
		return true
	}
	return false
}
