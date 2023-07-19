package main

import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/akamensky/argparse"
	"github.com/trongdth/token-checker/m/v2/app"
	"github.com/trongdth/token-checker/m/v2/app/interfaces"
	"github.com/trongdth/token-checker/m/v2/app/models"
	"go.uber.org/zap"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// conf := config.GetConfig()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	container := app.BuildContainer()
	err = container.Invoke(func(
		scanSvc interfaces.IScanService,
		assetSvc interfaces.IAssetService,
	) error {

		// Create new parser object
		parser := argparse.NewParser("token-checker", "Program that verify token address")

		// sync command
		syncCmd := parser.NewCommand("sync", "synching asset data ...")

		// check command
		checkCmd := parser.NewCommand("check", "CHECK THE TOKEN ADDRESS IS VALID OR NOT")
		tokenAddr := checkCmd.String("a", "address", &argparse.Options{Help: "token address"})

		// Now parse the arguments
		err := parser.Parse(os.Args)
		if err != nil {
			zap.L().Error(parser.Usage(err))
			return err
		}

		if syncCmd.Happened() {
			return scanSvc.ScanData()

		} else if checkCmd.Happened() {

			if *tokenAddr != "" {
				var asset *models.TwTAsset

				if asset, err = assetSvc.VerifyToken(*tokenAddr); err != nil {
					return err
				}

				zap.L().Info("FOUND", zap.Any("TOKEN:", asset))
			}

		} else {
			return errors.New("something weird happened")
		}

		return nil
	})

	if err != nil {
		zap.L().Error("error: ", zap.Error(err))
	}
}
