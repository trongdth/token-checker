package repositories

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) (err error) {
	if err = container.Provide(NewAssetRepository); err != nil {
		return err
	}

	if err = container.Provide(NewBlockchainRepository); err != nil {
		return err
	}
	return nil
}
