package main

import (
	"github.com/worldkk1/work-ticket-based-api/cmd/server"
	"github.com/worldkk1/work-ticket-based-api/config"
	"github.com/worldkk1/work-ticket-based-api/internal/database"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewHttpServer(conf, db).Start()
}
