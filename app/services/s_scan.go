package services

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/trongdth/token-checker/m/v2/app/interfaces"
)

// ScanService struct
type ScanService struct {
	blockchainRepos interfaces.IBlockchainRepository
}

// NewScanService : create new instance ScanService
func NewScanService(blockchainRepos interfaces.IBlockchainRepository) interfaces.IScanService {
	return &ScanService{
		blockchainRepos: blockchainRepos,
	}
}

func (bcSvc *ScanService) ScanData() error {
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
			log.Println(path)
			wg.Done()
		}(arrPaths[i])

	}
	wg.Wait()
	return nil
}
