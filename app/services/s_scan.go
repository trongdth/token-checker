package services

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/trongdth/token-checker/m/v2/app/interfaces"
	"github.com/trongdth/token-checker/m/v2/app/models"
	"github.com/trongdth/token-checker/m/v2/pkg/utils"
)

// ScanService struct
type ScanService struct {
	blockchainRepos interfaces.IBlockchainRepository
	assetRepos      interfaces.IAssetRepository
}

// NewScanService : create new instance ScanService
func NewScanService(blockchainRepos interfaces.IBlockchainRepository, assetRepos interfaces.IAssetRepository) interfaces.IScanService {
	return &ScanService{
		blockchainRepos: blockchainRepos,
		assetRepos:      assetRepos,
	}
}

func (sSvc *ScanService) ScanData() error {
	var (
		wg       sync.WaitGroup
		arrPaths []string
	)

	err := filepath.Walk("./assets/blockchains",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && info.Name() == "info.json" {
				arrPaths = append(arrPaths, path)
			}

			return nil
		})
	if err != nil {
		return err
	}

	wg.Add(len(arrPaths))
	for i := 0; i < len(arrPaths); i++ {
		go func(path string) {
			if err := sSvc.parseData(path); err != nil {
				// TODO
			}
			wg.Done()
		}(arrPaths[i])

	}
	wg.Wait()

	return nil
}

func (sSvc *ScanService) parseData(path string) error {

	log.Println(path)
	var data map[string]interface{}

	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, &data); err != nil {
		return err
	}

	if data["type"] == "coin" {
		if err := sSvc.parseBlockchainData(data); err != nil {
			log.Println(err)
			return err
		}
	} else {
		if err := sSvc.parseAssetData(data); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (sSvc *ScanService) parseBlockchainData(data map[string]interface{}) error {
	var blockchain *models.TwTBlockchain

	if err := utils.Copy(&blockchain, data); err != nil {
		return err
	}

	return sSvc.blockchainRepos.Save(blockchain)
}

func (sSvc *ScanService) parseAssetData(data map[string]interface{}) error {

	var asset *models.TwTAsset

	if err := utils.Copy(&asset, data); err != nil {
		return err
	}

	return sSvc.assetRepos.Save(asset)
}
