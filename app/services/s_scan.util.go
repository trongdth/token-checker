package services

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (sSvc *ScanService) getArrayPath() ([]string, error) {
	var arrPaths []string

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
		return nil, err
	}

	return arrPaths, nil
}

func (sSvc *ScanService) getBlockchainIdentityFromPath(path string) string {
	if m, _ := regexp.MatchString("^assets\\/blockchains\\/\\S*\\/info\\/info.json$", path); m {
		path = strings.ReplaceAll(path, "assets/blockchains/", "")
		path = strings.ReplaceAll(path, "/info/info.json", "")
		return strings.Trim(path, " ")

	} else if m, _ := regexp.MatchString("^assets\\/blockchains\\/\\S*\\/assets/\\S*/info.json$", path); m {
		path = strings.ReplaceAll(path, "assets/blockchains/", "")
		arr := strings.Split(path, "/asset")
		return strings.Trim(arr[0], " ")
	}
	return ""
}
