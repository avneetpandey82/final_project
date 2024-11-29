package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type SignupData struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
}

func handleUserRegister(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusBadRequest)
		return
	}
	var data SignupData
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	collection := mongoClient.Database("UsersData").Collection("users")
	_, err = collection.InsertOne(context.TODO(), data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := "Signup successful! Welcome, " + r.FormValue("fullName") + "."

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
func main() {
	var err error
	// username := os.Getenv("MONGODB_USERNAME")
	username := "avneetpandey82"
	// password := os.Getenv("MONGODB_PASSWORD")
	password := "1tf74EVaWYVNkWvC"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.orwve.mongodb.net/", username, password)
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.HandleFunc("/api/registerUser", handleUserRegister)
	log.Fatal(http.ListenAndServe(":80", nil))
}
