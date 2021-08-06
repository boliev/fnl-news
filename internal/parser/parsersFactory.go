package parser

import (
	"github.com/boliev/fnl-news/internal/source"
	"github.com/boliev/fnl-news/pkg/httpClient"
)

// GetParsers returns parsers list
func GetParsers() []*Parser {
	var parsers []*Parser
	var tagMatcher = NewTagMatcher()
	var client = httpclient.NewResty()
	var client1251 = httpclient.NewResty1251()

	sportBox := NewParser(source.NewSportboxSource(), tagMatcher, client)
	parsers = append(parsers, sportBox)
	kulichki := NewParser(source.NewKulichkiParser(), tagMatcher, client1251)
	parsers = append(parsers, kulichki)
	oneFnl := NewParser(source.NewOnefnlParser(), tagMatcher, client)
	parsers = append(parsers, oneFnl)
	sportsru := NewParser(source.NewSportsruParser(), tagMatcher, client)
	parsers = append(parsers, sportsru)

	return parsers
}
