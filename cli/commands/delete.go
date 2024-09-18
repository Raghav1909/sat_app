package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func DeleteCommand(db *models.Queries) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a student by name",
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

			err = db.DeleteStudent(context.Background(), name)
			if err != nil {
				fmt.Printf("Error deleting student: %v\n", err)
			} else {
				fmt.Printf("Student '%s' deleted successfully!\n", name)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the student to delete")

	return cmd
}
