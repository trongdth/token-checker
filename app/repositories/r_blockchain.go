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

// BlockchainRepository : struct
type BlockchainRepository struct {
	db                   interfaces.IDatabase
	blockchainCollection mongo.Collection
}

// NewBlockchainRepository : create new instance BlockchainRepository
func NewBlockchainRepository(db interfaces.IDatabase) interfaces.IBlockchainRepository {
	blockchainCollectName := indexes.GetBlockchainCollection().Name
	return &BlockchainRepository{
		db:                   db,
		blockchainCollection: *db.GetInstance().Collection(blockchainCollectName),
	}
}

func (bcRepos *BlockchainRepository) Save(blockchain *models.TwTBlockchain) error {
	blockchain.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeOut)
	defer cancel()

	filter := bson.D{bson.E{Key: "id", Value: blockchain.ID}}
	opts := options.Update().SetUpsert(true)
	_, err := bcRepos.blockchainCollection.UpdateOne(ctx, filter, bson.M{"$set": blockchain}, opts)
	return err
}
