package connections

import (
	"context"
	"docker/config"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongodb *mongo.Database

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://192.168.40.131:27017"))
	if err != nil {
		log.Println("errors connection mongo: ", err.Error())
		panic(err)
	}

	if client == nil {
		panic(errors.New("client is nil"))
	}
	Mongodb = client.Database(config.MONGO_DATABASE_USER)
}
