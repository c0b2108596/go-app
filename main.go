package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/c0b2108596/TodoList/go-app/databases/database"
)

type fruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var fruits = []fruit{
	{ID: 1, Name: "apple"},
	{ID: 2, Name: "banana"},
	{ID: 3, Name: "grape"},
}

func main() {
	http.HandleFunc("/", getFruits)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//データベース接続
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(fruits)
}
