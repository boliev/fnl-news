package parser

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsersFactorySuccess(t *testing.T) {
	parsers := parser.GetParsers()
	assert.NotEmpty(t, parsers)
}
