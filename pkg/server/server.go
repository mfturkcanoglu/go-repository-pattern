package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/controller"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/database"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/dto"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/repository"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/service"
)

var (
	logger *log.Logger     = log.Default()
	client database.Client = database.NewSqliteClient(logger)

	todoRepository repository.Repository[entity.Todo]        = repository.NewTodoRepository(logger, client)
	todoService    service.Service[entity.Todo, dto.TodoDto] = service.NewTodoService(logger, todoRepository)
	todoController controller.Controller                     = controller.NewTodoController(logger, todoService)
)

type Server struct {
	router *chi.Mux
	logger *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: logger,
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(jsonMiddleware)
	s.router.Use(middleware.StripSlashes)

	s.logger.Println("Register endpoints")
	// Register todo endpoints
	s.router.Get("/api/todo/{id}", todoController.GetByID)
	s.router.Get("/api/todo", todoController.ListAll)
	s.router.Post("/api/todo", todoController.Post)
	s.router.Put("/api/todo", todoController.Put)
	s.router.Delete("/api/todo", todoController.Delete)
}

func (s *Server) ListenAndServe() {

	httpServer := http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      s.router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	// Initialize database connection and run auto migrations
	client.Connect()
	client.Migrate()

	s.logger.Fatalln(httpServer.ListenAndServe())
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
