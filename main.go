package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	router "rest-api/http"
	"time"
)

var httpRouter router.Router = router.NewMuxRouter()

func main() {
	const port string = "8000"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://trinadhkoya:iP%40ssword77@cluster0.khuaz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//which will list out all the database names [admin config local website]
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)

	websiteDatabase := client.Database("website")
	blogCollection := websiteDatabase.Collection("blogs")

	cursor, err := blogCollection.Find(ctx, bson.M{})
	var posts []bson.M
	if err != nil {
		fmt.Println("No record found")
	}
	if err = cursor.All(ctx, &posts); err != nil {
		fmt.Println("Boom")
	}

	for _, post := range posts {
		fmt.Println(post)
	}

	httpRouter.SERVE(port)
}
