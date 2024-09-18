package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Raghav1909/sat_app/internal/db/models"
	"net/http"
)

func CreateStudentHandler(w http.ResponseWriter, r *http.Request, queries *models.Queries) {
	var request StudentRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		ClientError(w, http.StatusBadRequest)
		return
	}

	// Convert the "PASS"/"FAIL" value to a boolean (true for "PASS", false for "FAIL")
	passed, err := request.ConvertPassed()
	if err != nil {
		ClientError(w, http.StatusBadRequest)
		return
	}

	// Create an instance of `CreateStudentParams` with proper SQL types
	params := models.CreateStudentParams{
		Name:     request.Name,
		Address:  sql.NullString{String: request.Address, Valid: request.Address != ""},
		City:     sql.NullString{String: request.City, Valid: request.City != ""},
		Country:  sql.NullString{String: request.Country, Valid: request.Country != ""},
		Pincode:  sql.NullString{String: request.Pincode, Valid: request.Pincode != ""},
		SatScore: sql.NullInt64{Int64: request.SatScore, Valid: true},
		Passed:   passed, // Use the converted boolean value for "passed"
	}

	// Pass the params struct to the generated `CreateStudent` method
	err = queries.CreateStudent(context.Background(), params)
	if err != nil {
		ServerError(w, err)
		return
	}

	// Return success response
	response := Response{Message: "Student created successfully"}
	JsonResponse(w, response)
}

func GetAllStudentsHandler(w http.ResponseWriter, queries *models.Queries) {
	// Retrieve all students using SQLC-generated method
	students, err := queries.GetAllStudents(context.Background())
	if err != nil {
		ServerError(w, err)
		return
	}

	// Prepare and send the response as JSON
	response := Response{
		Message: "List of all students",
		Data:    students,
	}
	JsonResponse(w, response)
}

func GetStudentRankHandler(w http.ResponseWriter, r *http.Request, queries *models.Queries) {
	name := r.URL.Query().Get("name")
	if name == "" {
		ClientError(w, http.StatusBadRequest)
		return
	}

	rank, err := queries.GetStudentRank(context.Background(), name)
	if err != nil {
		ServerError(w, err)
		return
	}

	// Prepare and send the response as JSON
	response := Response{
		Message: "Student rank retrieved",
		Data:    map[string]interface{}{"rank": rank},
	}
	JsonResponse(w, response)
}
