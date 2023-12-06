package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Delnyavor/go-fiber-mongo-hrms/configs"

	"fmt"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg *MongoInstance

const dbName = "hrms"

// const mongoURI = "mongodb://localhost:27017/" + dbName
// const mongoURI = "mongodb+srv://delnyavor:npckaOvvwJC8IZF0@cluster0.yhpupim.mongodb.net/?retryWrites=true&w=majority"

func Connect() error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	Mg = &MongoInstance{
		Client: client,
		Db:     client.Database(dbName),
	}

	return nil

}

func EnvMongoURI() {
	panic("unimplemented")
}

func Disconnect() {
	if err := Mg.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
