package fnlnews

import (
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/boliev/fnl-news/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App struct
type App struct {
	cfg               *config.Config
	db                *gorm.DB
	articleRepository *repository.ArticleRepository
	parsers           []parser.Parser
	publishers        []publisher.Publisher
}

// Start the app
func (app App) Start() {
	app.bootstrap()
	log.Info("Starting to parse news...")
	// Parse
	for _, prsr := range app.parsers {
		articles, err := prsr.Parse()
		if err != nil {
			log.Warnf("error: %s", err.Error())
			continue
		}
		app.articleRepository.SaveAll(articles)
	}
	//Publish
	for _, pblisher := range app.publishers {
		pblisher.PublishNew()
	}
}

func (app *App) bootstrap() {
	log.SetLevel(log.InfoLevel)
	app.cfg = app.setupConfig()
	app.db = app.setupDB(app.cfg.GetString("database_dsn"))
	app.articleRepository = repository.CreateArticleRepository(app.db)
	app.publishers = app.getPublishers(app.articleRepository, app.cfg)
	app.parsers = app.getParsers(app.cfg)
}

func (app App) getPublishers(
	articleRepository *repository.ArticleRepository,
	config *config.Config,
) []publisher.Publisher {
	tgConfig := &publisher.TelegramPublisherConfig{
		ChatID: config.GetString("tg_chat_id"),
		Token:  config.GetString("tg_token"),
	}

	return publisher.GetPublishers(articleRepository, tgConfig)
}

func (app App) getParsers(
	config *config.Config,
) []parser.Parser {
	var parsersConfig map[string]parser.Config
	err := config.UnmarshalKey("parsers", &parsersConfig)
	if err != nil {
		log.Panicf("Unable to get parsers config, %v", err)
	}
	return parser.GetParsers(parsersConfig)
}

func (app App) setupDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	err = db.AutoMigrate(&domain.Article{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	return db
}

func (app App) setupConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf(err.Error())
	}

	return cfg
}
