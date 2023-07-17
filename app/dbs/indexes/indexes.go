package indexes

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection struct
type Collection struct {
	Name   string
	Indexs []mongo.IndexModel
}

// Collections struct
type Collections struct {
	Collections []*Collection
}

// GetCollections :
func GetCollections() *Collections {
	return &Collections{
		Collections: []*Collection{
			GetBlockchainCollection(),
			GetAssetCollection(),
		},
	}
}
