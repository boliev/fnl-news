package parser

// TagMatcher interface
type TagMatcher interface {
	MatchTags(text string) []string
}
