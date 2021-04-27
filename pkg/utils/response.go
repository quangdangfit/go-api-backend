package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/quangdangfit/go-backend/app/schema"
	"github.com/quangdangfit/go-backend/pkg/errors"
)

// ErrorResponse return response with error
func ErrorResponse(c *gin.Context, status int, err error, wrapMsg ...string) {
	e := errors.WithMessage(err, wrapMsg...)
	c.JSON(status, schema.Response{Error: errors.FromError(e)})
}
