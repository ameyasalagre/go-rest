package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"modules"
	"utils"

	"go.mongodb.org/mongo-driver/bson"
)

/****************************************.
*	path : "/"
*	purpose : Test Page
*	TO DO : ------
*
******************************************/
func Home(w http.ResponseWriter, r *http.Request) {
	collection := utils.ConnectAndGetMongoDbCollection("LoginApp", "users")
	filter := bson.M{"ip": "192.168.0.103"}
	result := modules.User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
	go utils.CloseMongoDbClient()

}
