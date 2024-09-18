package main

import (
	"github.com/Raghav1909/sat_app/cli/root"
	"log"
	"os"
)

func main() {
	// Execute the root command
	if err := root.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
