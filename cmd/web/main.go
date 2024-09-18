package main

import (
	"database/sql"
	"github.com/Raghav1909/sat_app/db/models"
	"github.com/Raghav1909/sat_app/web/handlers"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetStudents(w, r, queries)
	})

	// http.HandleFunc("/create-student", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.CreateStudent(w, r, queries)
	// })

	// http.HandleFunc("/student-rank", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetStudentRank(w, r, queries)
	// })
	//
	// http.HandleFunc("/update-student", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.UpdateStudent(w, r, queries)
	// })
	//
	// http.HandleFunc("/delete-student", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.DeleteStudent(w, r, queries)
	// })
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
