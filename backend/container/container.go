package container

import (
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/server"
	"backend/internal/service"
)

type Container struct {
	config *Config
	db     database.Service
	server *server.Server

	// Services
	authService *service.AuthService

	// Repositories
	userRepository *repository.UserRepository

	// Handlers
	authHandler *handler.AuthHandler
}

type Config struct {
	AccessSecret string
}

func NewContainer(config *Config) *Container {
	return &Container{
		config: config,
	}
}

func (c *Container) Build() error {

	if err := c.initDB(); err != nil {
		return err
	}

	c.initRepositories()
	c.initServices()
	c.initHandlers()
	c.setupServer(c.db)

	routerConfig := server.RouterConfig{
		AuthHandler: c.authHandler,
	}

	c.setupRoutes(routerConfig)

	return nil
}

func (c *Container) Server() *server.Server {
	return c.server
}

func (c *Container) initDB() error {
	db := database.New()
	c.db = db
	return nil
}

func (c *Container) initRepositories() {
	c.userRepository = repository.NewUserRepository(c.db)
}

func (c *Container) initServices() {
	c.authService = service.NewAuthService(
		c.userRepository,
	)
}

func (c *Container) initHandlers() {
	c.authHandler = handler.NewAuthHandler(c.authService)
}

func (c *Container) setupRoutes(routerConfig server.RouterConfig) {
	c.server.SetupRoutes(routerConfig)
}

func (c *Container) setupServer(db database.Service) {
	c.server = server.NewServer(db)
}
