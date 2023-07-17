package repositories

import (
	"github.com/trongdth/token-checker/m/v2/app/interfaces"
)

// BlockchainRepository : struct
type BlockchainRepository struct {
	db interfaces.IDatabase
}

// NewTVoucherTnxRepository : create new instance TVoucherTnxRepository
func NewBlockchainRepository(db interfaces.IDatabase) interfaces.IBlockchainRepository {
	return &BlockchainRepository{
		db: db,
	}
}
