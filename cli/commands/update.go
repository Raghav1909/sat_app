package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func UpdateCommand(db *models.Queries) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a student's SAT score and passed status",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Please provide a name using the --name flag.")
				return
			}

			_, err := db.GetStudentByName(context.Background(), name)
			if err == sql.ErrNoRows {
				fmt.Printf("Student with name '%s' does not exist.\n", name)
				return
			} else if err != nil {
				fmt.Printf("Error checking if student exists: %v\n", err)
				return
			}

			var satScore int
			for {
				fmt.Print("Enter SAT score (out of 1600): ")
				_, err := fmt.Scanf("%d", &satScore)
				if err != nil || satScore < 0 || satScore > 1600 {
					fmt.Println("Invalid SAT score. Score should be between 0 and 1600.")
				} else {
					break
				}
			}

			passed := (float64(satScore) / float64(1600) * 100) > 30

			err = db.UpdateStudentScore(context.Background(), models.UpdateStudentScoreParams{
				SatScore: int64(satScore),
				Passed:   passed,
				Name:     name,
			})

			if err != nil {
				fmt.Printf("Error updating student: %v\n", err)
			} else {
				fmt.Println("Student updated successfully!")
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the student to update")

	return cmd
}
