package parser

import "github.com/boliev/fnl-news/internal/domain"

// Parser interface
type Parser interface {
	GetName() string
	Parse() ([]*domain.Article, error)
}
