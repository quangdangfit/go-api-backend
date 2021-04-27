package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-backend/app/api"
	"github.com/quangdangfit/go-backend/app/dbs"
	"github.com/quangdangfit/go-backend/app/repositories"
	"github.com/quangdangfit/go-backend/app/router"
	"github.com/quangdangfit/go-backend/app/services"
	"github.com/quangdangfit/go-backend/pkg"
	"github.com/quangdangfit/go-backend/pkg/logger"
)

// BuildContainer create container and inject objects to container
func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject packages
	err := pkg.Inject(container)
	if err != nil {
		logger.Fatal("Failed to inject packages ", err)
	}

	// Inject Database
	err = dbs.Inject(container)
	if err != nil {
		logger.Fatal("Failed to inject database instance ", err)
	}

	// Inject repositories
	err = repositories.Inject(container)
	if err != nil {
		logger.Fatal("Failed to inject repositories ", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		logger.Fatal("Failed to inject services ", err)
	}

	// Inject APIs
	err = api.Inject(container)
	if err != nil {
		logger.Fatal("Failed to inject APIs ", err)
	}

	return container
}

// InitGinEngine init gin engine
func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "api-key"},
		MaxAge:           12 * time.Hour,
	}))

	err := router.RegisterAPI(app, container)
	if err != nil {
		logger.Fatal("Cannot register API: ", err)
	}

	return app
}
