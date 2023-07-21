package services

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (sSvc *ScanService) getArrayPath() ([]string, []string, error) {
	var (
		arrAssetPaths      []string
		arrBlockchainPaths []string
	)

	err := filepath.Walk("./assets/blockchains",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && info.Name() == "info.json" {
				if m, _ := sSvc.isBlockchainPath(path); m {
					arrBlockchainPaths = append(arrBlockchainPaths, path)
				} else if m, _ := sSvc.isAssetPath(path); m {
					arrAssetPaths = append(arrAssetPaths, path)
				}

			}

			return nil
		})
	if err != nil {
		return nil, nil, err
	}

	return arrBlockchainPaths, arrAssetPaths, nil
}

func (sSvc *ScanService) getBlockchainIdentityFromPath(path string) string {
	if m, _ := sSvc.isBlockchainPath(path); m {
		path = strings.ReplaceAll(path, "assets/blockchains/", "")
		path = strings.ReplaceAll(path, "/info/info.json", "")
		return strings.Trim(path, " ")

	} else if m, _ := sSvc.isAssetPath(path); m {
		path = strings.ReplaceAll(path, "assets/blockchains/", "")
		arr := strings.Split(path, "/asset")
		return strings.Trim(arr[0], " ")
	}
	return ""
}

func (sSvc *ScanService) isBlockchainPath(path string) (bool, error) {
	return regexp.MatchString("^assets\\/blockchains\\/\\S*\\/info\\/info.json$", path)
}

func (sSvc *ScanService) isAssetPath(path string) (bool, error) {
	return regexp.MatchString("^assets\\/blockchains\\/\\S*\\/assets/\\S*/info.json$", path)
}
