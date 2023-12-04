package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "github.com/c0b2108596/go-app/models"
)

func main() {
	http.HandleFunc("/", getFruits)

	//データベース接続
	db, err := database.Connect()
	fmt.Println("hello")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("データベース接続成功")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tests")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s", id, name)

		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//json.NewEncoder(w).Encode(fruits)
}
