package services

import (
	"github.com/trongdth/token-checker/m/v2/app/interfaces"
)

// AssetService struct
type AssetService struct {
	assetRepos interfaces.IAssetRepository
}

// NewAssetService : create new instance AssetService
func NewAssetService(assetRepos interfaces.IAssetRepository) interfaces.IAssetService {
	return &AssetService{
		assetRepos: assetRepos,
	}
}

func (aSvc *AssetService) VerifyToken(tokenAddr string) {

}
