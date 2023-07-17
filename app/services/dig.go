package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(NewAssetService); err != nil {
		return err
	}

	if err = container.Provide(NewScanService); err != nil {
		return err
	}

	return nil
}
