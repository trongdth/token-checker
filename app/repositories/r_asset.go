package repositories

import "github.com/trongdth/token-checker/m/v2/app/interfaces"

// AssetRepository : struct
type AssetRepository struct {
	db interfaces.IDatabase
}

// NewAssetRepository : create new instance AssetRepository
func NewAssetRepository(db interfaces.IDatabase) interfaces.IAssetRepository {
	return &AssetRepository{
		db: db,
	}
}
