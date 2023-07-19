package repositories

import (
	"context"

	"github.com/trongdth/token-checker/m/v2/app/dbs/indexes"
	"github.com/trongdth/token-checker/m/v2/app/interfaces"
	"github.com/trongdth/token-checker/m/v2/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AssetRepository : struct
type AssetRepository struct {
	db              interfaces.IDatabase
	assetCollection mongo.Collection
}

// NewAssetRepository : create new instance AssetRepository
func NewAssetRepository(db interfaces.IDatabase) interfaces.IAssetRepository {
	assetCollectName := indexes.GetAssetCollection().Name
	return &AssetRepository{
		db:              db,
		assetCollection: *db.GetInstance().Collection(assetCollectName),
	}
}

func (aRepos *AssetRepository) Save(asset *models.TwTAsset) error {
	asset.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeOut)
	defer cancel()

	filter := bson.D{bson.E{Key: "id", Value: asset.ID}}
	opts := options.Update().SetUpsert(true)
	_, err := aRepos.assetCollection.UpdateOne(ctx, filter, bson.M{"$set": asset}, opts)
	return err
}

func (aRepos *AssetRepository) FindAsset(tokenAddr string) (*models.TwTAsset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeOut)
	defer cancel()

	var asset *models.TwTAsset

	if err := aRepos.assetCollection.FindOne(ctx, bson.M{"id": tokenAddr}).Decode(&asset); err != nil {
		return nil, err
	}

	return asset, nil
}
