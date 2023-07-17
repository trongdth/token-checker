package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatabase interface {
	GetInstance() *mongo.Database
	EnsureIndexes()
}
