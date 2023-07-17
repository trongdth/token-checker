package indexes

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlockchainCollection() *Collection {
	collectionName := "twt_blockchains"
	indexes := []mongo.IndexModel{}
	return &Collection{
		Name:   collectionName,
		Indexs: indexes,
	}
}
