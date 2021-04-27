package repositories

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(NewRepository); err != nil {
		return err
	}
	return nil
}
