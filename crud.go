package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Inserting the record
func (apiCfg apiConifg) insertRecord(record any) (interface{}, error) {
	// Maybe extract the dbName and the collectionName from .env file?
	collection := apiCfg.dbClient.Database("munch").Collection("Testing")
	// Do i wanna return the ObjectId?
	result, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted record with _id: %v", result.InsertedID)
	return result.InsertedID, nil
}

// Reading a record, given an Object ID
func (apiCfg apiConifg) readRecord(objectId any) *mongo.SingleResult {
	collection := apiCfg.dbClient.Database("munch").Collection("Testing")
	// Just finding the first record
	filter := bson.D{{"_id", objectId}}
	resultCursor := collection.FindOne(context.TODO(), filter)
	return resultCursor
}

// Just deleting the first record as of now
// I have tested it and it works, just don't know where this will fit in as of now
func (apiCfg apiConifg) deleteRecord() error {
	collection := apiCfg.dbClient.Database("munch").Collection("Testing")
	filter := bson.D{{}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Println("Deleted Record")
	return nil
}

// Updating the first recorded as of now
// I have tested it and it works, just don't know where this will fit in as of now
func (apiCfg apiConifg) updateRecord() error {
	collection := apiCfg.dbClient.Database("munch").Collection("Testing")
	filter := bson.D{{}}
	update := bson.D{{"$set", bson.D{{"dataday.time", time.Now().UTC()}, {"updated", true}}}}
	updatedDoc, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Printf("Updated %v records", updatedDoc.ModifiedCount)
	return nil

}
