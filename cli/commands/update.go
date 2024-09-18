package commands

import (
	"context"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func UpdateCommand(db *models.Queries) *cobra.Command {
	return &cobra.Command{
		Use:   "update [name]",
		Short: "Update a student's SAT score and passed status",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			var satScore int
			for {
				fmt.Print("Enter SAT score (out of 1600): ")
				_, err := fmt.Scanf("%d", &satScore)
				if err != nil || satScore < 0 || satScore > 1600 {
					fmt.Println("Invalid SAT score. Score should be between between 0 and 1600.")
				} else {
					break
				}
			}

			passed := (float64(satScore) / float64(1600) * 100) > 30

			err := db.UpdateStudentScore(context.Background(), models.UpdateStudentScoreParams{
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
}
