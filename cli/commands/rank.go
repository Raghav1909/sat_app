package commands

import (
	"context"
	"fmt"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func GetRankCommand(db *models.Queries) *cobra.Command {
	return &cobra.Command{
		Use:   "rank [name]",
		Short: "Get the rank of a student by SAT score",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			rank, err := db.GetStudentRank(context.Background(), name)
			if err != nil {
				fmt.Printf("Error fetching rank: %v\n", err)
				return
			}
			fmt.Printf("Student Rank: %d\n", rank)
		},
	}
}
