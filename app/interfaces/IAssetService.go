package interfaces

import "github.com/trongdth/token-checker/m/v2/app/models"

type IAssetService interface {
	VerifyToken(tokenAddr string) (*models.TwTAsset, error)
}
