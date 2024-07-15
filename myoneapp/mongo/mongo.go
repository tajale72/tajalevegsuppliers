package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"myoneapp/model"
)

const (
	mongoURI = "mongodb://localhost:27017"
	dbName   = "tajalevegsuppliers"
	collName = "bill_details"
)

type DBClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func GetDBConnection() (*DBClient, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	db := client.Database(dbName)

	return &DBClient{Client: client, DB: db}, nil
}

func (dbClient *DBClient) Close() {
	if dbClient.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		dbClient.Client.Disconnect(ctx)
	}
}

func (dbClient *DBClient) CreateTable(data []byte) error {
	var req interface{}
	json.Unmarshal(data, &req)

	fmt.Println("Data", string(data))

	// var req model.Request
	// if err := json.Unmarshal(data, &req); err != nil {
	// 	log.Println("Unmarshal", err)
	// 	return fmt.Errorf("error unmarshaling request: %w", err)
	// }

	// cellections, _ := dbClient.DB.ListCollectionNames(context.Background(), bson.M{})
	//log.Println("calling mongo cellections", cellections)

	collection := dbClient.DB.Collection(collName)

	_, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Println("error inserting document", err)
		return fmt.Errorf("error inserting document: %w", err)
	}

	return nil
}

func (dbClient *DBClient) GetProducts() ([]model.Request, error) {
	collection := dbClient.DB.Collection(collName)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer cursor.Close(context.TODO())

	var listOfBills []model.Request
	for cursor.Next(context.TODO()) {
		var bill model.Request
		if err := cursor.Decode(&bill); err != nil {
			log.Println("Error decoding document:", err)
			return nil, fmt.Errorf("error decoding document: %w", err)
		}
		listOfBills = append(listOfBills, bill)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, fmt.Errorf("error iterating over cursor: %w", err)
	}

	return listOfBills, nil
}

func (dbClient *DBClient) GetProductByID(id int) (model.Request, error) {
	collection := dbClient.DB.Collection(collName)
	var bill model.Request

	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&bill)
	if err != nil {
		log.Println("Error getting data from the database:", err)
		return bill, fmt.Errorf("error getting data from the database: %w", err)
	}

	return bill, nil
}

func (dbClient *DBClient) UpdateUser(id int, newUsername, newEmail string) error {
	collection := dbClient.DB.Collection(collName)

	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"username", newUsername}, {"email", newEmail}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating user:", err)
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (dbClient *DBClient) DeleteUser(id int) error {
	collection := dbClient.DB.Collection(collName)

	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting user:", err)
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

type Result struct {
	BillNumber int64 `json:"billNumber"`
}

func (dbClient *DBClient) GetLastBillNumber() (Result, error) {
	collection := dbClient.DB.Collection(collName)
	billnumber, _ := collection.CountDocuments(context.Background(), bson.M{})
	// opts := options.FindOne().SetSort(bson.D{{"billNumber", -1}})
	var result Result
	result.BillNumber = billnumber + 1
	return result, nil
}

type User struct {
	ID       int
	Username string
	Email    string
}
