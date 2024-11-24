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

	// Repositories
	userRepository   *repository.UserRepository
	tenantRepository *repository.TenantRepository

	// Services
	authService *service.AuthService

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
	c.setupServer()

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
	c.tenantRepository = repository.NewTenantRepository(c.db)
}

func (c *Container) initServices() {
	c.authService = service.NewAuthService(
		c.db.DB(),
		c.tenantRepository,
		c.userRepository,
	)
}

func (c *Container) initHandlers() {
	c.authHandler = handler.NewAuthHandler(c.authService)
}

func (c *Container) setupServer() {
	c.server = server.NewServer(c.db)
}

func (c *Container) setupRoutes(routerConfig server.RouterConfig) {
	c.server.SetupRoutes(routerConfig)
}
