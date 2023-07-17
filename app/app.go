package app

import (
	"log"

	"github.com/trongdth/token-checker/m/v2/app/dbs"
	"github.com/trongdth/token-checker/m/v2/app/repositories"
	"github.com/trongdth/token-checker/m/v2/app/services"
	"go.uber.org/dig"
)

// BuildContainer create container and inject objects to container
func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject Database
	err := dbs.Inject(container)
	if err != nil {
		log.Fatalf("Failed to inject database instance %v", err)
	}

	// Inject repositories
	err = repositories.Inject(container)
	if err != nil {
		log.Fatalf("Failed to inject repositories %v", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		log.Fatalf("Failed to inject services %v", err)
	}

	return container
}
