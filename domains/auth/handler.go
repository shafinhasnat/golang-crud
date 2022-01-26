package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shafinhasnat/kloudlab/db"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
type LoginStruct struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

var col = db.ConnectDB().Database("goTest").Collection("users")

func SomeColl() {
	doc := bson.M{"title": "bongo bangladesh"}
	res, err := col.InsertOne(context.TODO(), doc)
	fmt.Println(res, err, doc)
	// db.
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	insertedData, err := col.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertedData)
	json.NewEncoder(w).Encode(insertedData.InsertedID)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// var result bson.M
	ctx := context.TODO()
	var data []Data
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var res Data
		cur.Decode(&res)
		data = append(data, res)
	}
	json.NewEncoder(w).Encode(data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginStruct
	var result Data
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("banglas")
		panic(err)
	}
	e := col.FindOne(ctx, LoginStruct{Email: payload.Email, Password: payload.Password}).Decode(&result)
	if e != nil {
		// json.NewEncoder(w).Encode()
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", Register).Methods("POST")
	r.HandleFunc("/auth/users", GetUsers).Methods("GET")
	r.HandleFunc("/auth/login", Login).Methods("Post")
}
