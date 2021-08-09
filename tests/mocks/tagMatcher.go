package mocks

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/stretchr/testify/mock"
)

// TagMatcherMock mock for TagMatcher struct
type TagMatcherMock struct {
	parser.TagMatcher
	mock.Mock
}

// MatchTags mock method for TagMatcher struct
func (t *TagMatcherMock) MatchTags(text string) []string {
	args := t.Called(text)
	return args.Get(0).([]string)
}
