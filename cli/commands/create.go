package commands

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Raghav1909/sat_app/db/models"
	"github.com/spf13/cobra"
)

func CreateCommand(db *models.Queries) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new student",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			var name string
			for {
				fmt.Print("Enter name: ")
				name, _ = reader.ReadString('\n')
				name = sanitizeInput(name)

				_, err := db.GetStudentByName(context.Background(), name)

				if err == nil {
					fmt.Printf("A student with the name '%s' already exists. Please choose a different name.\n", name)
				} else if err == sql.ErrNoRows {
					break
				} else {
					fmt.Printf("Error checking student existence: %v\n", err)
					return
				}
			}

			fmt.Print("Enter address: ")
			address, _ := reader.ReadString('\n')
			address = sanitizeInput(address)

			fmt.Print("Enter city: ")
			city, _ := reader.ReadString('\n')
			city = sanitizeInput(city)

			fmt.Print("Enter country: ")
			country, _ := reader.ReadString('\n')
			country = sanitizeInput(country)

			var pincode string
			for {
				fmt.Print("Enter pincode (6 digits): ")
				pincode, _ = reader.ReadString('\n')
				pincode = sanitizeInput(pincode)

				if !isValidPincode(pincode) {
					fmt.Println("Invalid pincode. Pincode must be a 6-digit number.")
				} else {
					break
				}
			}

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

			// Determine if the student has passed based on the SAT score
			passed := (float64(satScore) / float64(1600) * 100) > 30

			// Insert the new student into the database
			err := db.CreateStudent(context.Background(), models.CreateStudentParams{
				Name:     name,
				Address:  address,
				City:     city,
				Country:  country,
				Pincode:  pincode,
				SatScore: int64(satScore),
				Passed:   passed,
			})

			if err != nil {
				fmt.Printf("Error creating student: %v\n", err)
			} else {
				fmt.Println("Student created successfully!")
				if passed {
					fmt.Println("Student has passed SAT")
				} else {
					fmt.Println("Student has failed SAT")
				}
			}
		},
	}
}

// Helper function to sanitize input (removes trailing newlines)
func sanitizeInput(input string) string {
	return strings.TrimSpace(input)
}

// Helper function to check if pincode is a valid 6-digit number
func isValidPincode(pincode string) bool {
	if match, _ := regexp.MatchString(`^\d{6}$`, pincode); !match {
		return false
	}

	_, err := strconv.Atoi(pincode)
	return err == nil
}
