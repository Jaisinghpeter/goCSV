package mongodb

import (
	"context"
	"log"
	"time"
	"consumer/configurations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func DumpEmployeeList(messages []byte){
    client, err := mongo.NewClient(options.Client().ApplyURI(configurations.MongoDBURL))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    employeeDetailsDatabase := client.Database(configurations.MongoDBDatabase)
    employeeDetailsCollection := employeeDetailsDatabase.Collection(configurations.MongoDBEmployeeCollection)
	

	var doc []interface{}
	bson.UnmarshalExtJSON(messages, true, &doc)
	employeeDetailsCollection.InsertMany(ctx, doc)
}


