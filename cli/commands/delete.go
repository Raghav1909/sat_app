package commands

import (
	"context"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func DeleteCommand(db *models.Queries) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a student by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			err := db.DeleteStudent(context.Background(), name)
			if err != nil {
				fmt.Printf("Error deleting student: %v\n", err)
			} else {
				fmt.Println("Student deleted successfully!")
			}
		},
	}
}
