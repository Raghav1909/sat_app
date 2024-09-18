package main

import (
	"database/sql"
	"github.com/Raghav1909/sat_app/internal/db/models"
	"github.com/Raghav1909/sat_app/internal/handler"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the database
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the SQLC queries struct
	queries := models.New(db)

	// Set up the routes
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handler.CreateStudentHandler(w, r, queries)
		case "GET":
			handler.GetAllStudentsHandler(w, queries)
			// case "PUT":
			// 	handler.UpdateStudentScoreHandler(w, r, queries)
			// case "DELETE":
			// 	handler.DeleteStudentHandler(w, r, queries)
		}
	})

	// http.HandleFunc("/student/rank", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		handler.GetStudentRankHandler(w, r, queries)
	// 	}
	// })

	// Start the server
	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
