package fnlnews

import (
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/boliev/fnl-news/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// App struct
type App struct {
	Cfg               *config.Config
	Db                *gorm.DB
	ArticleRepository *repository.ArticleRepository
	publishers        []publisher.Publisher
	parsers           []parser.Parser
}

// Start the app
func (app App) Start() {
	log.SetLevel(log.InfoLevel)
	log.Info("Starting to parse news...")
	start := time.Now()
	// Parse
	for _, prsr := range app.parsers {
		articles, err := prsr.Parse()
		if err != nil {
			log.Warnf("error: %s", err.Error())
			continue
		}
		app.ArticleRepository.SaveAll(articles)
	}
	//Publish
	for _, pblisher := range app.publishers {
		pblisher.PublishNew()
	}
	log.Infof("Finish %s", time.Now().Sub(start).String())
}
