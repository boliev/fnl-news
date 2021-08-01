package domain

import "gorm.io/gorm"

// Article DTO
type Article struct {
	gorm.Model
	Title  string
	Href   string `gorm:"uniqueIndex"`
	Tags   string
	HTML   string `gorm:"-"`
	TgSent bool   `gorm:"default:false"`
}
