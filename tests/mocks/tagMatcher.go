package mocks

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/stretchr/testify/mock"
)

type TagMatcherMock struct {
	parser.TagMatcher
	mock.Mock
}

func (t *TagMatcherMock) MatchTags(text string) []string {
	args := t.Called(text)
	return args.Get(0).([]string)
}
