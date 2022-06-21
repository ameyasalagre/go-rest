package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/****************************************.
*	MongoDB getConnection
* 	MongoDb getCollection
* 	MongoDb  closeConnection
*  	TODO: Logging,Refactoring
******************************************/

var client *mongo.Client
var Version="1.0"
var log = New().With(nil, "version", Version)
/********
*	connect To MongoDb and get Collection by passing documentName,tableName as a String
* 	and Return the collection pointer
*********/
func ConnectAndGetMongoDbCollection(documentName string, tableName string) *mongo.Collection {
	// get a MongoDBClient
	client := getClient()
	log.Info("Connected to MongoDb")
	// Check the connection
	collection := client.Database(documentName).Collection(tableName)
	return collection
}

/********
*	Close  MongoDb connection
*	TODO:
*	~Stop MongoClient when its ideal for more than 2 mins
*********/
func CloseMongoDbClient() {
	client := getClient()
	client.Disconnect(context.TODO())
	log.Info("Connected to MongoDb")
}


/********
*	get MongoDb client
*********/
func getClient() *mongo.Client {
	url:= GetEnv("MONGO_URL","localhost:12354")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(err)
		}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error("Error While Connection", err)
	}
	return client 
}




