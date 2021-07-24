package fnlnews

import (
	"fmt"
	"github.com/boliev/fnl-news/internal/parser"
)

// App struct
type App struct {
}

// Start the app
func (app App) Start() {
	fmt.Println("Starting to parse news")
	fmt.Println(" - Sportbox")
	sportbox := &parser.Sportbox{
		Domain: "https://news.sportbox.ru",
		Path:   "/Vidy_sporta/Futbol/Russia/1st_division",
	}

	err := sportbox.Parse()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
