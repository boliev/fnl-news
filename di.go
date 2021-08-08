package fnlnews

import (
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/internal/publisher"
	"github.com/boliev/fnl-news/internal/repository"
	"github.com/boliev/fnl-news/pkg/config"
	"github.com/boliev/fnl-news/pkg/httpClient"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DiCreateDB di function for database
func DiCreateDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.GetString("database_dsn")), &gorm.Config{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	err = db.AutoMigrate(&domain.Article{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	return db
}

// DiCreateConfig di function for config
func DiCreateConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf(err.Error())
	}

	return cfg
}

// DiCreateArticleRepository di function for article repository
func DiCreateArticleRepository(db *gorm.DB) *repository.ArticleRepository {
	return repository.CreateArticleRepository(db)
}

// DiCreateHTTPClient di function http client
func DiCreateHTTPClient() httpclient.Client {
	return httpclient.NewResty()
}

// DiCreateApp di function for app
func DiCreateApp(
	cfg *config.Config,
	db *gorm.DB,
	articleRepository *repository.ArticleRepository,
	publishers []publisher.Publisher,
	parsers []*parser.Parser,
) *App {
	return &App{
		Cfg:               cfg,
		Db:                db,
		ArticleRepository: articleRepository,
		publishers:        publishers,
		parsers:           parsers,
	}
}

// DiCreatePublishers di function for publishers lis
func DiCreatePublishers(
	articleRepository *repository.ArticleRepository,
	config *config.Config,
	client httpclient.Client,
) []publisher.Publisher {
	tgConfig := &publisher.TelegramPublisherConfig{
		ChatID: config.GetString("tg_chat_id"),
		Token:  config.GetString("tg_token"),
	}

	return publisher.GetPublishers(articleRepository, tgConfig, client)
}

// DiCreateParsers di function for parsers lis
func DiCreateParsers() []*parser.Parser {
	return parser.GetParsers()
}
