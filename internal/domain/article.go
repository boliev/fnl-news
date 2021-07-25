package domain

import "gorm.io/gorm"

// Article DTO
type Article struct {
	gorm.Model
	Title    string
	Href     string `gorm:"uniqueIndex"`
	ImageURL string
	Date     string
	Tags     string
	HTML     string `gorm:"-"`
	IsSent   bool   `gorm:"default:false"`
}
