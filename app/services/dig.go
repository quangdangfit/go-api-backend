package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(NewService); err != nil {
		return err
	}

	return nil
}
