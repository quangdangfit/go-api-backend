package api

import (
	"go.uber.org/dig"
)

// Inject api
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(NewAPI); err != nil {
		return err
	}

	return nil
}
