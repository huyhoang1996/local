package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name       string
	Age        int
	City       string
	DateJoined time.Time
}

// func main() {
// 	// Rest of the code will go here
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	collection := client.Database("test").Collection("trainers")
// 	ash := Trainer{"Ash", 10, "Pallet Town"}
// 	// misty := Trainer{"Misty", 10, "Cerulean City"}
// 	// brock := Trainer{"Brock", 15, "Pewter City"}
// 	insertResult, err := collection.InsertOne(context.TODO(), ash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
// }

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	fmt.Println("Connected to MongoDB!")
	collection := c.Database("test").Collection("trainers")
	ash := Trainer{"Ash", 10, "Huy hoang", time.Now()}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	toDate, _ := time.Parse("02/01/2006 15:04:05", "01/10/2021 11:55:20")
	// toDate := time.Date(2014, time.November, 5, 0, 0, 0, 0, time.UTC)
	// filter := bson.D{
	// 	{"dateJoined", bson.D{
	// 		{"$gt", initDate},
	// 	}},
	// }
	fmt.Println("Inserted a single document: ", toDate)

	// filter := bson.M{
	// 	"dateJoined": bson.M{
	// 		"$lt": toDate,
	// 	},
	// }
	filter := bson.M{"city": "Huy hoang"}
	heroes := ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Age, hero.City, hero.DateJoined)
	}
}

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Trainer {
	var trainers []*Trainer
	collection := client.Database("test").Collection("trainers")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var trainer Trainer
		err = cur.Decode(&trainer)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		trainers = append(trainers, &trainer)
	}
	return trainers
}
