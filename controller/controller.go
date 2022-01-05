package controller

import (
	"assignment1/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//this file contains database connection,database helper,controllers
const connectionString = "mongodb+srv://Sourav:testUser@cluster0.rqxgg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "users"
const colName = "data"

//Most important as connection
var collection *mongo.Collection

//connect with mongo db
//will be called automatically
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo DB
	client, err := mongo.Connect(context.TODO(), clientOption)

	errHandler(err)

	fmt.Println("Mongo DB connect success")

	//create db and collection
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//***MongoDB Helper functions start***//

//will Insert one user
func insertUser(user models.User) primitive.ObjectID {
	inserted, err := collection.InsertOne(context.Background(), user)
	errHandler(err)
	return inserted.InsertedID.(primitive.ObjectID)
}

// will delete one user
func deleteUser(id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	errHandler(err)
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	errHandler(err)
}

///will return all users
func getAllUsers() []models.User {
	var users []models.User
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	errHandler(err)
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		errHandler(err)
		users = append(users, user)
	}
	return users
}

//will update single user
func updateUser(id string, user models.User) {
	objectID, err := primitive.ObjectIDFromHex(id)
	errHandler(err)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"name": user.Name, "dob": user.DOB, "address": user.Address, "description": user.Description}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	errHandler(err)
}

//get single user
func getUser(id string) models.User {
	objectID, err := primitive.ObjectIDFromHex(id)
	errHandler(err)
	filter := bson.M{"_id": objectID}
	var user models.User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	errHandler(err)
	return user
}

//delete all users
func deleteAllUsers() {
	_, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	errHandler(err)
}

//***MongoDB Helper functions end***//

//***Controller functions start***//

//create user
func InsertUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user models.User
	var err = json.NewDecoder(r.Body).Decode(&user)
	errHandler(err)
	user.CreatedAt = time.Now()
	id := insertUser(user)
	user.ID = id
	json.NewEncoder(w).Encode(user)
}
//Delete one user
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	id := params["id"]
	deleteUser(id)
	json.NewEncoder(w).Encode(id)
}
// get all users
func GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	var users = getAllUsers()
	json.NewEncoder(w).Encode(users)
}
//update user
func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	var user models.User
	//requires all params of user(model) in body except id,createdAt
	var err = json.NewDecoder(r.Body).Decode(&user)
	errHandler(err)
	params := mux.Vars(r)
	id := params["id"]
	updateUser(id, user)
	user.ID, _ = primitive.ObjectIDFromHex(id)
	json.NewEncoder(w).Encode(user)
}
//get single user
func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	id := params["id"]
	var user = getUser(id)
	json.NewEncoder(w).Encode(user)
}
//delete all users
func DeleteAllUsersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllUsers()
	json.NewEncoder(w).Encode("All users deleted")
}

//***Controller functions end***//

//handle  errors
func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
