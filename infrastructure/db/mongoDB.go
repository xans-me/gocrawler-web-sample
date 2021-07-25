package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gocrawler-web-sample/infrastructure/configuration"
	"time"
)

// NewMongodb constructor to use mongo db
func NewMongodb(conf *configuration.AppConfig) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.Connection))
	if err != nil {
		log.Panic(err)
	}

	MongoDb := client.Database(conf.Mongodb.DbName)

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error("MongoDB : Error ", err)
		return MongoDb
	}

	log.Info("MongoDB : Ready")
	return MongoDb
}
