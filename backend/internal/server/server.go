package server

import (
	"backend/internal/database"
	"backend/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	auth fiber.Router
}

type RouterConfig struct {
	AuthHandler *handler.AuthHandler
}

type Server struct {
	App *fiber.App
	db  database.Service
}

func NewServer(db database.Service) *Server {
	app := fiber.New()
	return &Server{
		App: app,
		db:  db,
	}
}

func (s *Server) SetupRoutes(config RouterConfig) {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiRouter := s.App.Group("/api")

	router := Router{
		auth: apiRouter.Group("/auth"),
	}

	setupPublicRoutes(router.auth, config)

	s.App.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(s.db.Health())
	})
}

func setupPublicRoutes(router fiber.Router, config RouterConfig) {
	router.Post("/register", config.AuthHandler.Register)
}

func (s *Server) Listen() error {
	if err := s.App.Listen(":8080"); err != nil {
		return nil
	}

	return nil
}
