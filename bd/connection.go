package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@twittercluster.rofavhb.mongodb.net/?retryWrites=true&w=majority")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("DB Connected")
	return client
}

func CheckConnetion() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
