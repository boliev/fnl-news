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
	"os"
)

// App struct
type App struct {
}

// Start the app
func (app App) Start() {
	cfg := app.getConfig()
	app.setupLoger(cfg.GetString("log_file"))
	var parsersConfig map[string]parser.Config
	err := cfg.UnmarshalKey("parsers", &parsersConfig)
	if err != nil {
		log.Panicf("Unable to get parsers config, %v", err)
	}

	db := app.createDB(cfg.GetString("database_dsn"))
	articleRepository := repository.CreateArticleRepository(db)
	publishers := app.getPublishers(articleRepository, cfg)
	parsers := app.getParsers(parsersConfig)
	log.Info("Starting to parse news")
	for _, prsr := range parsers {
		articles, err := prsr.Parse()
		if err != nil {
			log.Warnf("error: %s", err.Error())
			continue
		}
		articleRepository.SaveAll(articles)
	}
	for _, pblisher := range publishers {
		pblisher.PublishNew()
	}
}

func (app App) createDB(dsn string) *gorm.DB {
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

func (app App) getParsers(config map[string]parser.Config) []parser.Parser {
	var parsers []parser.Parser
	if sportboxConfig, ok := config["sportbox"]; ok {
		sportbox := parser.NewSportboxParser(sportboxConfig)
		parsers = append(parsers, sportbox)
	} else {
		log.Warnf("Unable to find config for sportbox parser")
	}

	if onefnlConfig, ok := config["onefnl"]; ok {
		onefnl := parser.NewOnefnlParser(onefnlConfig)
		parsers = append(parsers, onefnl)
	} else {
		log.Warnf("Unable to find config for onefnl parser")
	}

	return parsers
}

func (app App) getPublishers(
	articleRepository *repository.ArticleRepository, config *config.Config) []publisher.Publisher {
	var publishers []publisher.Publisher

	telegramPublisher := publisher.NewTelegramPublisher(
		articleRepository,
		config.GetString("tg_chat_id"),
		config.GetString("tg_token"),
	)
	publishers = append(publishers, telegramPublisher)

	return publishers
}

func (app App) getConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf(err.Error())
	}

	return cfg
}

func (app App) setupLoger(logFile string) {
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Warnf("cant open log file %s", logFile)
	}
	log.SetLevel(log.InfoLevel)
	log.SetOutput(f)
}
