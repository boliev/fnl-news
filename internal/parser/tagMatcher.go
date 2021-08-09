package parser

type TagMatcher interface {
	MatchTags(text string) []string
}
