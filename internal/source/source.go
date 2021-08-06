package source

// Source interface
type Source interface {
	Name() string
	Domain() string
	Path() string
	ArticlePattern() string
	ArticleListItemPattern() string
	ArticleListItemTitlePattern() string
	ArticleListItemHrefPattern() string
	ArticleFullURL(relative string) string
	ArticleFullTitle(title string) string
	ShouldSkipArticle(text string) bool
}
