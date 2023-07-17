package indexes

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAssetCollection() *Collection {
	collectionName := "twt_assets"
	indexes := []mongo.IndexModel{}
	return &Collection{
		Name:   collectionName,
		Indexs: indexes,
	}
}
