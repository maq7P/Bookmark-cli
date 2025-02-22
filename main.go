package main

import (
	"bookmarkapp/cli"
	"bookmarkapp/services"
)

func main() {
	// Инициализация сервиса закладок
	bookmarkService := services.NewBookmarksService(
		&services.FileService{},
		&services.ConfigService{},
	)

	// Инициализация cli закладок
	bookmarkCli := cli.NewBookmarkCli(&bookmarkService)
	bookmarkCli.Run()
}