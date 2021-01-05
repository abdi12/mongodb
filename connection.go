package main

import (
   "context"
   "fmt"
   "log"
//   "time"

//   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
//   "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://admin:Karimun!%4034@mongos-nww.btpns.com:27017/")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    // Insert a document
    collection := client.Database("mydb").Collection("persons")

    ruan := Person{"Abdi", 26, "Balige City"}

    insertResult, err := collection.InsertOne(context.TODO(), ruan)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}

