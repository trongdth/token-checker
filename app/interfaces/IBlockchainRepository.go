package interfaces

import "github.com/trongdth/token-checker/m/v2/app/models"

type IBlockchainRepository interface {
	Save(*models.TwTBlockchain) error
}
