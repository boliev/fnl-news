package fnlnews

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/boliev/fnl-news/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App struct
type App struct {
}

// Start the app
func (app App) Start() {
	cfg := app.getConfig()
	var parsersConfig map[string]parser.Config
	err := cfg.UnmarshalKey("parsers", &parsersConfig)
	if err != nil {
		fmt.Printf("Unable to get parsers config, %v", err)
	}

	db := app.createDB(cfg.GetString("database_dsn"))
	articleRepository := repository.CreateArticleRepository(db)
	publishers := app.getPublishers(articleRepository, cfg)
	parsers := app.getParsers(parsersConfig)
	fmt.Println("Starting to parse news")
	for _, prsr := range parsers {
		fmt.Printf(" - %s\n", prsr.GetName())
		articles, err := prsr.Parse()
		if err != nil {
			fmt.Printf("error: %s", err.Error())
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
		panic(fmt.Sprintf("error: %s", err.Error()))
	}
	err = db.AutoMigrate(&domain.Article{})
	if err != nil {
		panic(fmt.Sprintf("error: %s", err.Error()))
	}

	return db
}

func (app App) getParsers(config map[string]parser.Config) []parser.Parser {
	var parsers []parser.Parser
	if sportboxConfig, ok := config["sportbox"]; ok {
		sportbox := parser.NewSportboxParser(sportboxConfig)
		parsers = append(parsers, sportbox)
	} else {
		fmt.Printf("Unable to find config for sportbox parser")
	}

	if onefnlConfig, ok := config["onefnl"]; ok {
		onefnl := parser.NewOnefnlParser(onefnlConfig)
		parsers = append(parsers, onefnl)
	} else {
		fmt.Printf("Unable to find config for onefnl parser")
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
		panic(err.Error())
	}

	return cfg
}
