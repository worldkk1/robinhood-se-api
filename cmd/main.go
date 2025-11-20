package main

import (
	"github.com/worldkk1/robinhood-se-api/cmd/server"
	"github.com/worldkk1/robinhood-se-api/config"
	"github.com/worldkk1/robinhood-se-api/internal/database"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewHttpServer(conf, db).Start()
}
