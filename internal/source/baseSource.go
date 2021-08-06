package source

// BaseSource parser struct
type BaseSource struct {
	name                        string
	domain                      string
	path                        string
	articlePattern              string
	articleListItemPattern      string
	articleListItemTitlePattern string
	articleListItemHrefPattern  string
}

// Name returns name of the source
func (s BaseSource) Name() string {
	return s.name
}

// Domain returns domain of the source
func (s BaseSource) Domain() string {
	return s.domain
}

// Path returns path of the source
func (s BaseSource) Path() string {
	return s.path
}

// ArticlePattern returns article text pattern
func (s BaseSource) ArticlePattern() string {
	return s.articlePattern
}

// ArticleListItemPattern returns articles list pattern
func (s BaseSource) ArticleListItemPattern() string {
	return s.articleListItemPattern
}

// ArticleListItemTitlePattern returns article title pattern from the articles list
func (s BaseSource) ArticleListItemTitlePattern() string {
	return s.articleListItemTitlePattern
}

// ArticleListItemHrefPattern returns article href pattern from the articles list
func (s BaseSource) ArticleListItemHrefPattern() string {
	return s.articleListItemHrefPattern
}
