package interfaces

import "github.com/trongdth/token-checker/m/v2/app/models"

type IAssetRepository interface {
	Save(*models.TwTAsset) error
}
