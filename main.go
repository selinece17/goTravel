package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
Travel Planner Console Application
Author: Selin Ece Birgül

Description:
This program is a console-based travel itinerary planner.
Users can add destinations, assign a number of days and a
budget to each destination, and store activities for each
location. The program includes full input validation and
stores trip data in a JSON file so the itinerary persists
between program runs.

Major features:
- Add destinations
- Add activities to destinations
- View full itinerary
- View total trip budget
- Remove destinations
- Automatically save/load trip data using JSON
*/

type Destination struct {
	Name       string   `json:"name"`       // Destination name
	Days       int      `json:"days"`       // Number of days staying
	Budget     float64  `json:"budget"`     // Budget allocated for this destination
	Activities []string `json:"activities"` // List of planned activities
}

/*
Trip represents the entire itinerary and stores
all destinations the user adds.
*/
type Trip struct {
	Destinations []Destination `json:"destinations"` // Slice storing all destinations
}

// Global variable storing the current trip data
var trip Trip

// Reader used to capture user input from the terminal
var reader = bufio.NewReader(os.Stdin)

// JSON file where trip data will be saved
const fileName = "trip_data.json"

/*
Program entry point.
Loads saved trip data and repeatedly runs the menu loop.
*/
func main() {

	// Load trip data from JSON file if it exists
	loadTrip()

	// Infinite loop keeps the program running until user exits
	for {

		printMenu()

		// Get menu choice from the user
		choice := getIntInput("Choose option: ")

		// Perform action based on user choice
		switch choice {

		case 1:
			addDestination()

		case 2:
			addActivity()

		case 3:
			viewItinerary()

		case 4:
			showBudget()

		case 5:
			removeDestination()

		case 6:
			fmt.Println("Goodbye!")
			saveTrip() // Save trip before exiting
			return

		default:
			fmt.Println("Invalid menu choice. Please select 1-6.")
		}
	}
}

/*
Displays the main menu options.
*/
func printMenu() {

	fmt.Println("\n====== Travel Planner ======")
	fmt.Println("1. Add Destination")
	fmt.Println("2. Add Activity")
	fmt.Println("3. View Itinerary")
	fmt.Println("4. View Budget Summary")
	fmt.Println("5. Remove Destination")
	fmt.Println("6. Exit")
	fmt.Println("============================")
}

/*
Reads and validates string input from the user.
*/
func getStringInput(prompt string) string {

	for {

		fmt.Print(prompt)

		// Read input until newline
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		// Remove newline and extra whitespace
		input = strings.TrimSpace(input)

		// Prevent empty input
		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		return input
	}
}

/*
Reads integer input and ensures it is valid.
*/
func getIntInput(prompt string) int {

	for {

		fmt.Print(prompt)

		// Read input
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		// Convert string to integer
		value, err := strconv.Atoi(input)

		// If conversion fails, ask again
		if err != nil {
			fmt.Println("Please enter a valid whole number.")
			continue
		}

		return value
	}
}

/*
Reads floating point numbers safely.
*/
func getFloatInput(prompt string) float64 {

	for {

		fmt.Print(prompt)

		// Read user input
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		// Convert string to float
		value, err := strconv.ParseFloat(input, 64)

		// If conversion fails, re-prompt
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		return value
	}
}

/*
Ensures destination names contain letters only.
*/
func isValidDestinationName(name string) bool {

	// Reject empty names
	if name == "" {
		return false
	}

	// Check every character in the string
	for _, char := range name {

		// If any character is not a letter, reject
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}

/*
Checks if the destination already exists.
*/
func destinationExists(name string) bool {

	// Loop through all destinations
	for _, d := range trip.Destinations {

		// Case-insensitive comparison
		if strings.EqualFold(d.Name, name) {
			return true
		}
	}

	return false
}

/*
Adds a new destination to the itinerary.
*/
func addDestination() {

	var name string

	for {

		name = getStringInput("Destination name (letters only): ")

		// Validate name format
		if !isValidDestinationName(name) {
			fmt.Println("Invalid name. Only letters allowed.")
			continue
		}

		// Prevent duplicate destinations
		if destinationExists(name) {
			fmt.Println("Destination already exists.")
			continue
		}

		break
	}

	var days int

	for {

		days = getIntInput("Number of days: ")

		// Ensure days are reasonable
		if days <= 0 || days > 365 {
			fmt.Println("Days must be between 1 and 365.")
			continue
		}

		break
	}

	var budget float64

	for {

		budget = getFloatInput("Budget for this destination: ")

		// Budget validation
		if budget <= 0 {
			fmt.Println("Budget must be greater than 0.")
			continue
		}

		// Prevent unrealistic values
		if budget > 1000000 {
			fmt.Println("Budget too large.")
			continue
		}

		break
	}

	// Create destination struct
	dest := Destination{
		Name:       name,
		Days:       days,
		Budget:     budget,
		Activities: []string{}, // initialize empty activity list
	}

	// Append destination to trip slice
	trip.Destinations = append(trip.Destinations, dest)

	fmt.Println("Destination added.")
}

/*
Adds an activity to a selected destination.
*/
func addActivity() {

	// Ensure at least one destination exists
	if len(trip.Destinations) == 0 {
		fmt.Println("Add a destination first.")
		return
	}

	viewDestinations()

	var index int

	for {

		// Convert user input to slice index
		index = getIntInput("Select destination number: ") - 1

		// Ensure index is valid
		if index < 0 || index >= len(trip.Destinations) {
			fmt.Println("Invalid destination number.")
			continue
		}

		break
	}

	var activity string

	for {

		activity = getStringInput("Activity name: ")

		// Minimum length validation
		if len(activity) < 2 {
			fmt.Println("Activity name too short.")
			continue
		}

		// Prevent excessively long activity names
		if len(activity) > 50 {
			fmt.Println("Activity name too long.")
			continue
		}

		break
	}

	// Append activity to destination activity list
	trip.Destinations[index].Activities =
		append(trip.Destinations[index].Activities, activity)

	fmt.Println("Activity added.")
}

/*
Displays all destinations with numbering.
*/
func viewDestinations() {

	fmt.Println("\nDestinations:")

	for i, d := range trip.Destinations {

		// i+1 because users expect numbering starting at 1
		fmt.Printf("%d. %s\n", i+1, d.Name)
	}
}

/*
Prints the full trip itinerary.
*/
func viewItinerary() {

	// If no destinations exist, inform the user
	if len(trip.Destinations) == 0 {
		fmt.Println("No destinations planned.")
		return
	}

	fmt.Println("\n====== Trip Itinerary ======")

	for _, d := range trip.Destinations {

		fmt.Println("\nDestination:", d.Name)
		fmt.Println("Days:", d.Days)
		fmt.Printf("Budget: %.2f\n", d.Budget)

		// If destination has no activities
		if len(d.Activities) == 0 {

			fmt.Println("Activities: None")

		} else {

			fmt.Println("Activities:")

			// Print each activity
			for _, a := range d.Activities {

				fmt.Println("-", a)
			}
		}
	}
}

/*
Calculates total trip budget.
*/
func showBudget() {

	total := 0.0

	// Sum budgets from all destinations
	for _, d := range trip.Destinations {

		total += d.Budget
	}

	fmt.Printf("\nTotal Trip Budget: %.2f\n", total)
}

/*
Removes a destination from the itinerary.
*/
func removeDestination() {

	if len(trip.Destinations) == 0 {
		fmt.Println("No destinations to remove.")
		return
	}

	viewDestinations()

	var index int

	for {

		// Convert user input to slice index
		index = getIntInput("Select destination number to remove: ") - 1

		if index < 0 || index >= len(trip.Destinations) {
			fmt.Println("Invalid destination number.")
			continue
		}

		break
	}

	// Store name of removed destination for confirmation
	removed := trip.Destinations[index].Name

	// Remove item from slice using Go slice manipulation
	trip.Destinations = append(
		trip.Destinations[:index],
		trip.Destinations[index+1:]...,
	)

	fmt.Println("Destination removed:", removed)
}

/*
Saves trip data to JSON file.
*/
func saveTrip() {

	// Open file (create if it doesn't exist, overwrite if it does)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Error saving file.")
		return
	}

	// Ensure file closes after function finishes
	defer file.Close()

	// Create JSON encoder
	encoder := json.NewEncoder(file)

	// Convert trip struct to JSON and write to file
	err = encoder.Encode(trip)

	if err != nil {
		fmt.Println("Error encoding data.")
		return
	}

	fmt.Println("Trip saved successfully.")
}

/*
Loads trip data from JSON file.
*/
func loadTrip() {

	// Attempt to open existing JSON file
	file, err := os.Open(fileName)

	// If file does not exist, start with empty trip
	if err != nil {
		return
	}

	defer file.Close()

	// Create JSON decoder
	decoder := json.NewDecoder(file)

	// Decode JSON into trip struct
	err = decoder.Decode(&trip)

	// If JSON file is corrupted
	if err != nil {
		fmt.Println("Saved file corrupted. Starting new trip.")
		trip = Trip{} // reset trip data
	}
}
