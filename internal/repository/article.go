package repository

import (
	"github.com/boliev/fnl-news/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ArticleRepository repository
type ArticleRepository struct {
	db *gorm.DB
}

// CreateArticleRepository creates the repository
func CreateArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

// Save saves an article
func (r ArticleRepository) Save(article *domain.Article) {
	r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&article)
}

// SaveAll saves several articles
func (r ArticleRepository) SaveAll(articles []*domain.Article) {
	for _, article := range articles {
		r.Save(article)
	}
}

// GetNewTg returns new articles
func (r ArticleRepository) GetNewTg() []*domain.Article {
	var articles []*domain.Article
	r.db.Where("tg_sent = ?", false).Find(&articles)
	return articles
}

// MarkAsSentTG mark article as sent to tg
func (r ArticleRepository) MarkAsSentTG(article *domain.Article) {
	r.db.Model(article).Update("tg_sent", true)
}
