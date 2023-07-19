package services

import (
	"encoding/json"
	"os"

	"github.com/trongdth/token-checker/m/v2/app/interfaces"
	"github.com/trongdth/token-checker/m/v2/app/models"
	"github.com/trongdth/token-checker/m/v2/pkg/utils"
)

// ScanService struct
type ScanService struct {
	blockchainRepos interfaces.IBlockchainRepository
	assetRepos      interfaces.IAssetRepository
	mapBlockchain   map[string]models.TwTBlockchain
}

// NewScanService : create new instance ScanService
func NewScanService(blockchainRepos interfaces.IBlockchainRepository, assetRepos interfaces.IAssetRepository) interfaces.IScanService {
	return &ScanService{
		blockchainRepos: blockchainRepos,
		assetRepos:      assetRepos,
		mapBlockchain:   map[string]models.TwTBlockchain{},
	}
}

func (sSvc *ScanService) ScanData() error {
	var (
		arrPaths []string
		err      error
	)

	if arrPaths, err = sSvc.getArrayPath(); err != nil {
		return err
	}

	for i := 0; i < len(arrPaths); i++ {
		if err = sSvc.parseData(arrPaths[i]); err != nil {
			return err
		}
	}

	return nil
}

func (sSvc *ScanService) parseData(path string) error {
	var data map[string]interface{}

	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, &data); err != nil {
		return err
	}

	name := sSvc.getBlockchainIdentityFromPath(path)
	if data["type"] == "coin" {
		blockchain, err := sSvc.parseBlockchainData(data)
		if err != nil {
			return err
		}

		if name != "" {
			sSvc.mapBlockchain[name] = *blockchain
		}

	} else {

		if err := sSvc.parseAssetData(data, sSvc.mapBlockchain[name]); err != nil {
			return err
		}
	}

	return nil
}

func (sSvc *ScanService) parseBlockchainData(data map[string]interface{}) (*models.TwTBlockchain, error) {
	var blockchain *models.TwTBlockchain

	if err := utils.Copy(&blockchain, data); err != nil {
		return nil, err
	}

	if err := sSvc.blockchainRepos.Save(blockchain); err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (sSvc *ScanService) parseAssetData(data map[string]interface{}, blockchain models.TwTBlockchain) error {

	var asset *models.TwTAsset

	if err := utils.Copy(&asset, data); err != nil {
		return err
	}

	asset.Blockchain = &blockchain
	return sSvc.assetRepos.Save(asset)
}
