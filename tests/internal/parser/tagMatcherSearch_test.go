package parser

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchTagsSuccess(t *testing.T) {
	matcher := parser.NewTagMatcher()
	tags := matcher.MatchTags("Сегодня Алания прбедила, а торпедо проиграло")
	assert.Equal(t, 2, len(tags))
	assert.Equal(t, "Алания", tags[0])
	assert.Equal(t, "Торпедо", tags[1])
}

func TestSearchTagsNoTags(t *testing.T) {
	matcher := parser.NewTagMatcher()
	tags := matcher.MatchTags("Today all went great!")
	assert.Empty(t, tags)
}
