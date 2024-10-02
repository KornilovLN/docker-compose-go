package main

import (
	"firstapp/chapter_os/about"
	"firstapp/chapter_os/demo"
	logging "firstapp/chapter_os/log"
	"firstapp/chapter_os/webserver"
)

func init() {
	logging.Init()
}

func main() {
	about.AboutInfo()

	// Демо работы с файлами, (пакет fileops)
	demo.DemonstrateFileOperations()

	// Запуск веб-сервера:
	// Там выводятся логи Gin:
	// localhost:<PORT>/logs
	// (port см. в webserver.go)
	webserver.StartWebServer()
}
