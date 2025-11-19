package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/worldkk1/robinhood-se-api/config"
	"github.com/worldkk1/robinhood-se-api/internal/database"
	handler "github.com/worldkk1/robinhood-se-api/internal/handlers"
	middleware "github.com/worldkk1/robinhood-se-api/internal/handlers/middleware"
	repository "github.com/worldkk1/robinhood-se-api/internal/repositories"
	usecase "github.com/worldkk1/robinhood-se-api/internal/usecases"
)

type httpServer struct {
	app  *http.ServeMux
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
	userRepo := repository.NewUserPostgresRepository(s.db)
	taskRepo := repository.NewTaskPostgresRepository(s.db)
	commentRepo := repository.NewCommentPostgresRepository(s.db)
	authUC := usecase.NewAuthUsecaseImpl(userRepo)
	taskUC := usecase.NewTaskUsecaseImpl(taskRepo)
	commentUC := usecase.NewCommentUsecaseImpl(commentRepo)
	authHTTP := handler.NewAuthHttpHandler(authUC)
	taskHTTP := handler.NewTaskHttpHandler(taskUC, commentUC)

	v1 := http.NewServeMux()
	v1.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	v1.HandleFunc("POST /auth/register", authHTTP.Register)
	v1.HandleFunc("POST /auth/login", authHTTP.Login)

	securedRouter := http.NewServeMux()
	securedRouter.HandleFunc("POST /tasks", middleware.CheckRoleAdminMiddleware(http.HandlerFunc(taskHTTP.CreateTask)))
	securedRouter.HandleFunc("GET /tasks", taskHTTP.GetTaskList)
	securedRouter.HandleFunc("GET /tasks/{id}", taskHTTP.GetTaskDetail)
	securedRouter.HandleFunc("PATCH /tasks/{id}", taskHTTP.EditTask)
	securedRouter.HandleFunc("PATCH /tasks/{id}/archive", taskHTTP.ArchiveTask)
	securedRouter.HandleFunc("POST /tasks/{id}/comments", taskHTTP.CreateTaskComment)
	securedRouter.HandleFunc("GET /tasks/{id}/comments", taskHTTP.GetTaskComments)
	securedRouter.HandleFunc("PATCH /tasks/{id}/comments/{commentId}", taskHTTP.EditTaskComment)
	securedRouter.HandleFunc("DELETE /tasks/{id}/comments/{commentId}", taskHTTP.DeleteTaskComment)
	v1.Handle("/", middleware.AuthMiddleware(securedRouter))

	s.app.Handle("/v1/", http.StripPrefix("/v1", v1))
	middlewareChain := middleware.MiddlewareChain(middleware.LoggerMiddleware)

	port := fmt.Sprintf(":%d", s.conf.Port)
	log.Println("Server running at ", port)
	err := http.ListenAndServe(port, middlewareChain(s.app))
	if err != nil {
		log.Fatalln("Server failed to start", err)
	}
}
