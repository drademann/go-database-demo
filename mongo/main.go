package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	godb := client.Database("godb")
	daysCollection := godb.Collection("days")

	count, err := daysCollection.CountDocuments(ctx, bson.D{})
	fmt.Printf("handling days collection with %d days\n", count)

	newDay := Day{
		Date: time.Now().Truncate(24 * time.Hour),
		Tasks: []Task{
			{
				Start: time.Date(2024, time.June, 3, 10, 0, 0, 0, time.Local),
				Text:  "Test Insert",
			},
		},
	}

	insertResult, err := daysCollection.InsertOne(ctx, newDay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted a new day with id %d\n", insertResult.InsertedID)

	count, err = daysCollection.CountDocuments(ctx, bson.D{})
	fmt.Printf("handling days collection with now %d days\n", count)

	var day Day
	hex, err := primitive.ObjectIDFromHex("6679356f8600d9946e82dd11")
	if err != nil {
		log.Fatal(err)
	}
	err = daysCollection.FindOne(ctx, bson.M{"_id": hex}).Decode(&day)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("found day with %d tasks\n", len(day.Tasks))
}

type Day struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Date     time.Time          `bson:"Date"`
	Finished time.Time          `bson:"Finished"`
	Tasks    []Task
}

type Task struct {
	Start   time.Time `bson:"Start"`
	Text    string    `bson:"Text"`
	IsPause bool      `bson:"IsPause"`
	Tags    []string  `bson:"Tags"`
}
