package parser

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSuccess(t *testing.T) {
	sourceMock := new(mocks.SourceMock)
	sourceMock.On("Domain").Return("https://test.com")
	sourceMock.On("Path").Return("/fnl/")
	sourceMock.On("Name").Return("Name")
	sourceMock.On("ArticleListItemPattern").Return("(.*?,.*?\n)")
	sourceMock.On("ArticleListItemTitlePattern").Return("(.*?),.*?\n")
	sourceMock.On("ArticleListItemHrefPattern").Return(".*?,(.*?)\n")
	sourceMock.On("ArticlePattern").Return("(.*)")
	sourceMock.On("ArticleFullTitle", "title1").Return("title1")
	sourceMock.On("ArticleFullTitle", "title2").Return("title2")
	sourceMock.On("ArticleFullTitle", "title3").Return("title3")
	sourceMock.On("ArticleFullURL", "link1").Return("link1")
	sourceMock.On("ArticleFullURL", "link2").Return("link2")
	sourceMock.On("ArticleFullURL", "link3").Return("link3")
	sourceMock.On("ShouldSkipArticle", "text1").Return(true)
	sourceMock.On("ShouldSkipArticle", "text2").Return(false)
	sourceMock.On("ShouldSkipArticle", "text3").Return(false)
	matcher := new(mocks.TagMatcherMock)
	clientMock := new(mocks.ClientMock)
	body := "title1,link1\ntitle2,link2\ntitle3,link3\n"
	clientMock.On("Get", "https://test.com/fnl/").Return(body, nil)
	clientMock.On("Get", "link1").Return("text1", nil)
	clientMock.On("Get", "link2").Return("", fmt.Errorf("some error"))
	clientMock.On("Get", "link3").Return("text3", nil)
	matcher.On("MatchTags", "title2 text2").Return([]string{})
	matcher.On("MatchTags", "title3 text3").Return([]string{"Alania", "Torpedo"})

	p := parser.NewParser(sourceMock, matcher, clientMock)

	articles, err := p.Parse()
	assert.Equal(t, 1, len(articles))
	assert.Nil(t, err)
	assert.Equal(t, "title3", articles[0].Title)
}

func TestGetListError(t *testing.T) {
	sourceMock := new(mocks.SourceMock)
	sourceMock.On("Domain").Return("https://test.com")
	sourceMock.On("Path").Return("/fnl/")
	sourceMock.On("Name").Return("Name")
	matcher := new(mocks.TagMatcherMock)
	clientMock := new(mocks.ClientMock)
	clientMock.On("Get", "https://test.com/fnl/").Return("", fmt.Errorf("some error"))
	p := parser.NewParser(sourceMock, matcher, clientMock)

	articles, err := p.Parse()
	assert.Empty(t, articles)
	assert.NotNil(t, err)
}

func TestNewsTextIsEmptyNoTags(t *testing.T) {
	sourceMock := new(mocks.SourceMock)
	sourceMock.On("Domain").Return("https://test.com")
	sourceMock.On("Path").Return("/fnl/")
	sourceMock.On("Name").Return("Name")
	sourceMock.On("ArticleListItemPattern").Return("(.*?,.*?\n)")
	sourceMock.On("ArticleListItemTitlePattern").Return("(.*?),.*?\n")
	sourceMock.On("ArticleListItemHrefPattern").Return(".*?,(.*?)\n")
	sourceMock.On("ArticlePattern").Return("<p>(.*)</p>")
	sourceMock.On("ArticleFullTitle", "title1").Return("title1")
	sourceMock.On("ArticleFullURL", "link1").Return("link1")
	sourceMock.On("ShouldSkipArticle", "").Return(false)
	matcher := new(mocks.TagMatcherMock)
	clientMock := new(mocks.ClientMock)
	clientMock.On("Get", "https://test.com/fnl/").Return("title1,link1\n", nil)
	clientMock.On("Get", "link1").Return("text1", nil)
	p := parser.NewParser(sourceMock, matcher, clientMock)
	matcher.On("MatchTags", "title1 ").Return([]string{})

	articles, err := p.Parse()
	assert.Equal(t, 1, len(articles))
	assert.Empty(t, articles[0].Tags)
	assert.Nil(t, err)
}
