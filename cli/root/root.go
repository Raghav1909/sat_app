package root

import (
	"database/sql"
	"log"

	"github.com/Raghav1909/sat_app/cli/commands"
	"github.com/Raghav1909/sat_app/db/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var db *models.Queries

// rootCmd is the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sat-cli",
	Short: "CLI for adding SAT scores",
	Long:  `A simple CLI for managing sat scores in the database`,
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	conn, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = models.New(conn)

	// Disable completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Register subcommands to the root command
	rootCmd.AddCommand(commands.CreateCommand(db))
	rootCmd.AddCommand(commands.GetAllCommand(db))
	rootCmd.AddCommand(commands.GetRankCommand(db))
	rootCmd.AddCommand(commands.UpdateCommand(db))
	rootCmd.AddCommand(commands.DeleteCommand(db))
}
