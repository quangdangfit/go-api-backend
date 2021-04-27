package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/go-backend/app/api"
	"go.uber.org/dig"
)

// RegisterAPI define api routes
func RegisterAPI(e *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		apiHandler *api.API,
	) error {
		_ = e.Group("/api")

		return nil
	})

	return err
}
