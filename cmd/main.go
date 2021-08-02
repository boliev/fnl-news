package main

import (
	"github.com/boliev/fnl-news"
	"go.uber.org/dig"
)

func main() {
	container := buildContainer()
	err := container.Invoke(func(app *fnlnews.App) {
		app.Start()
	})

	if err != nil {
		panic(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	provide(container, fnlnews.DiCreateConfig)
	provide(container, fnlnews.DiCreateDB)
	provide(container, fnlnews.DiCreateArticleRepository)
	provide(container, fnlnews.DiCreatePublishers)
	provide(container, fnlnews.DiCreateParsers)
	provide(container, fnlnews.DiCreateApp)

	return container
}

func provide(container *dig.Container, constructor interface{}) {
	err := container.Provide(constructor)
	if err != nil {
		panic(err)
	}
}
