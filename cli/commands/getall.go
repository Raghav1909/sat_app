package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

type Student struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Pincode  string `json:"pincode"`
	SatScore int64  `json:"sat_score"`
	Result   string `json:"result"`
}

func GetAllCommand(db *models.Queries) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all students",
		Run: func(cmd *cobra.Command, args []string) {
			students, err := db.GetAllStudents(context.Background())
			if err != nil {
				fmt.Printf("Error fetching students: %v\n", err)
				return
			}

			if len(students) == 0 {
				fmt.Println("No students in the database")
				return
			}

			var studentsDisplay []Student

			for _, student := range students {
				result := "FAIL"
				if student.Passed == 1 {
					result = "PASS"
				}
				studentsDisplay = append(studentsDisplay, Student{
					Name:     student.Name,
					Address:  student.Address,
					City:     student.City,
					Country:  student.Country,
					Pincode:  student.Pincode,
					SatScore: student.SatScore,
					Result:   result,
				})
			}
			studentJSON, err := json.MarshalIndent(studentsDisplay, "", "  ")
			if err != nil {
				fmt.Printf("Error converting students to JSON: %v\n", err)
				return
			}

			fmt.Println(string(studentJSON))
		},
	}
}
