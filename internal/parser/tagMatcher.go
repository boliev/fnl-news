package parser

import (
	"sort"
	"strings"
)

var availableTags = map[string][]string{
	"Алания":        {"Алани", "Владикавказ", "Барсы"},
	"Акрон":         {"Акрон", "Тольятт"},
	"Балтика":       {"Балти", "Калининград"},
	"Велес":         {"Велес"},
	"Волгарь":       {"Волгар", "Астрахан"},
	"Динамо":        {"Динамо", "Брянск"},
	"Енисей":        {"Енисе", "Красноярск"},
	"Иртыш":         {"Иртыш", "Омск"},
	"Краснодар_2":   {"Краснодар"},
	"Нефтехимик":    {"Нефтехимик", "Нижнекамск"},
	"Оренбург":      {"Оренбург", "Газовик"},
	"СКА_Хабаровск": {"Хабаровск"},
	"Спартак_2":     {"Спартак"},
	"Текстильщик":   {"Текстильщик"},
	"Томь":          {"Томь", "Томич", "Томск"},
	"Торпедо":       {"Торпедо"},
	"Факел":         {"Факел", "Воронеж"},
	"Чайка":         {"Чайка", "Песчанокопск"},
	"Чертаново":     {"Чертанов"},
	"О_Д":           {"Долгопрудны"},
	"Камаз":         {"Камаз", "Челны"},
	"Ротор":         {"Ротор", "Волгоград"},
	"Кубань":        {"Кубан"},
	"Металлург":     {"Металлург", "Липец"},
}

// TagMatcher struct
type TagMatcher struct {
}

// NewTagMatcher constructor
func NewTagMatcher() *TagMatcher {
	return &TagMatcher{}
}

// MatchTags matches teams tags from text
func (t TagMatcher) MatchTags(text string) []string {
	var tags []string
	for tag, patterns := range availableTags {
		for _, pattern := range patterns {
			if strings.Index(text, pattern) >= 0 {
				tags = append(tags, tag)
				break
			}
		}
	}
	sort.Strings(tags)

	return tags
}

// GetAllTags returns all tags
func (t TagMatcher) GetAllTags() []string {
	var tags []string
	for k := range availableTags {
		tags = append(tags, k)
	}
	sort.Strings(tags)

	return tags
}
