package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/worldkk1/robinhood-se-api/config"
	"github.com/worldkk1/robinhood-se-api/internal/database"
)

type httpServer struct {
	app *http.ServeMux
	db   database.Database
	conf *config.Config
}

func NewHttpServer(conf *config.Config, db database.Database) Server {
	app := http.NewServeMux()
	return &httpServer{
		app:  app,
		db:   db,
		conf: conf,
	}
}

func (s *httpServer) Start() {

	s.app.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	port := fmt.Sprintf(":%d", s.conf.Port)
	log.Println("Server running at ", port)
	err := http.ListenAndServe(port, s.app)
	if err != nil {
		log.Fatalln("Server failed to start", err)
	}
}
