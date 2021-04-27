package pkg

import (
	"go.uber.org/dig"

	"github.com/quangdangfit/go-backend/pkg/validation"
)

// Inject packages
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(validation.New); err != nil {
		return err
	}

	return nil
}
