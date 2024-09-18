package handlers

import (
	"context"
	"github.com/Raghav1909/sat_app/db/models"
	"html/template"
	"net/http"
)

func GetStudents(w http.ResponseWriter, r *http.Request, queries *models.Queries) {
	students, err := queries.GetAllStudents(context.Background())

	if err != nil {
		http.Error(w, "Failed to retrieve students", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html")

	if err != nil {
		ServerError(w, err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", students)

	if err != nil {
		ServerError(w, err)
	}
}
