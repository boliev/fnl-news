package fnlnews

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/domain"
	"github.com/boliev/fnl-news/internal/parser"
	"github.com/boliev/fnl-news/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App struct
type App struct {
}

// Start the app
func (app App) Start() {
	db := app.createDB()
	articleRepository := repository.CreateArticleRepository(db)
	parsers := app.getParsers()
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
}

func (app App) createDB() *gorm.DB {
	dsn := "host=localhost user=fnluser password=123456 dbname=fnl port=5432 sslmode=disable"
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

func (app App) getParsers() []parser.Parser {
	var parsers []parser.Parser

	sportbox := parser.NewSportboxParser("https://news.sportbox.ru", "/Vidy_sporta/Futbol/Russia/1st_division")
	parsers = append(parsers, sportbox)

	onefnl := parser.NewOnefnlParser("https://1fnl.ru", "/news/")
	parsers = append(parsers, onefnl)

	return parsers
}
