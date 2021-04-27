package dbs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/quangdangfit/go-backend/app/interfaces"
	"github.com/quangdangfit/go-backend/config"
	"github.com/quangdangfit/go-backend/pkg/errors"
)

// Database struct
type Database struct {
	dbInstance *mongo.Database
}

// NewDatabase new connect to database and return database interface
func NewDatabase() interfaces.IDatabase {
	dbConf := config.Config.Database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionStr := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=SCRAM-SHA-1%s",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name, dbConf.Options)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.Fatal(errors.New(errors.ECMongoConnection, err.Error(), "mongodb.Init"))
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(errors.New(errors.ECMongoConnection, err.Error(), "mongodb.Init"))
		return nil
	}

	database := Database{dbInstance: client.Database(dbConf.Name)}
	return &database
}

// GetInstance get mongodb instance
func (d *Database) GetInstance() *mongo.Database {
	return d.dbInstance
}
