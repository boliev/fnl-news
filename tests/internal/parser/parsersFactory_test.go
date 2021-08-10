package parser

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsersFactorySuccess(t *testing.T) {
	parsers := parser.GetParsers(new(mocks.TagMatcherMock), new(mocks.ClientMock), new(mocks.ClientMock))
	assert.NotEmpty(t, parsers)
}
