package dbs

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/trongdth/token-checker/m/v2/app/dbs/indexes"
	"github.com/trongdth/token-checker/m/v2/app/interfaces"
	"github.com/trongdth/token-checker/m/v2/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database struct
type Database struct {
	dbInstance *mongo.Database
}

// NewDatabase new connect to database and return database interface
func NewDatabase() interfaces.IDatabase {
	conf := config.GetConfig()

	// set client options
	clientOptions := options.Client().SetMaxConnIdleTime(5 * time.Second).SetMaxPoolSize(250).ApplyURI(conf.MongoDBURL)

	// connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	// check connection
	ctxCheckConn, cancelCheckConn := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelCheckConn()

	err = client.Ping(ctxCheckConn, nil)
	if err != nil {
		return nil
	}
	db := client.Database(conf.MongoDBName)
	database := &Database{dbInstance: db}

	// create indexes
	database.EnsureIndexes()

	return database
}

// GetInstance get mongodb instance
func (d *Database) GetInstance() *mongo.Database {
	return d.dbInstance
}

func (d *Database) EnsureIndexes() {
	type MongoIndex struct {
		Name string
		Keys interface{}
	}

	colls := indexes.GetCollections()
	for _, coll := range colls.Collections {

		var dropIndexs = []string{}

		c := d.dbInstance.Collection(coll.Name)
		duration := 10 * time.Second
		batchSize := int32(100)

		cur, err := c.Indexes().List(context.Background(), &options.ListIndexesOptions{
			BatchSize: &batchSize,
			MaxTime:   &duration,
		})
		if err != nil {
			log.Fatalf("Something went wrong listing %v", err)
		}

		sort.Slice(coll.Indexs, func(i, j int) bool {
			return *coll.Indexs[i].Options.Name <= *coll.Indexs[j].Options.Name
		})

		for cur.Next(context.Background()) {
			index := MongoIndex{}
			cur.Decode(&index)

			if index.Name == "_id_" {
				continue
			}

			idDrop := -1
			for i, v := range coll.Indexs {
				if *v.Options.Name == index.Name {
					idDrop = i
				}
			}

			// Drop all index is not found on which defined
			if idDrop == -1 {
				dropIndexs = append(dropIndexs, index.Name)
			}
		}
		d.dropIndexes(dropIndexs, c)
		d.createIndexes(coll.Indexs, c)
	}
}

func (d *Database) dropIndexes(dropIndexs []string, collection *mongo.Collection) {
	for _, indexStr := range dropIndexs {
		opts := options.DropIndexes().SetMaxTime(10 * time.Second)
		collection.Indexes().DropOne(context.Background(), indexStr, opts)
	}
}

func (d *Database) createIndexes(createIndexs []mongo.IndexModel, collection *mongo.Collection) {
	for _, index := range createIndexs {
		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		collection.Indexes().CreateOne(context.Background(), index, opts)
	}
}
