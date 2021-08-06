package source

// Sportbox parser struct
type Sportbox struct {
	BaseSource
}

// NewSportboxSource Sportbox source constructor
func NewSportboxSource() *Sportbox {
	return &Sportbox{
		BaseSource: BaseSource{
			name:                        "sportbox",
			domain:                      "https://news.sportbox.ru",
			path:                        "/Vidy_sporta/Futbol/Russia/1st_division",
			articlePattern:              "<div class=\"js-mediator-article\">(.*?)</div>",
			articleListItemPattern:      "<li><div class=\"_Sportbox_Spb2015_Components_TeazerBlock_TeazerBlock\">(.*?)</li>",
			articleListItemTitlePattern: "<span class=\"text\">(.*?)</span>",
			articleListItemHrefPattern:  "href=\"(.*?)\"",
		},
	}
}

// ArticleFullURL returns full article url from relative link
func (s Sportbox) ArticleFullURL(relative string) string {
	return s.Domain() + relative
}

// ArticleFullTitle returns full article title for the source
func (s Sportbox) ArticleFullTitle(title string) string {
	return title
}

// ShouldSkipArticle validates if the text is correct and we should add it to tDB
func (s Sportbox) ShouldSkipArticle(text string) bool {
	return false
}
