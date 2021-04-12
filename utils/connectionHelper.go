package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client
//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error
//Used to execute client creation procedure only once.
var mongoOnce sync.Once

const connectionURI= "mongodb+srv://trinadhkoya:iP%40ssword77@cluster0.khuaz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func GetMongoClient() (*mongo.Client, error) {

	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(connectionURI)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
