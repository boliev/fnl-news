package repository

import (
	"github.com/boliev/fnl-news/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ArticlePostgreRepository repository
type ArticlePostgreRepository struct {
	db *gorm.DB
}

// CreateArticlePostgreRepository creates the repository
func CreateArticlePostgreRepository(db *gorm.DB) *ArticlePostgreRepository {
	return &ArticlePostgreRepository{
		db: db,
	}
}

// Save saves an article
func (r ArticlePostgreRepository) Save(article *domain.Article) {
	r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&article)
}

// SaveAll saves several articles
func (r ArticlePostgreRepository) SaveAll(articles []*domain.Article) {
	for _, article := range articles {
		r.Save(article)
	}
}

// GetNewTg returns new articles
func (r ArticlePostgreRepository) GetNewTg() []*domain.Article {
	var articles []*domain.Article
	r.db.Where("tg_sent = ?", false).Find(&articles)
	return articles
}

// MarkAsSentTG mark article as sent to tg
func (r ArticlePostgreRepository) MarkAsSentTG(article *domain.Article) {
	r.db.Model(article).Update("tg_sent", true)
}
