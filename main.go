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

/*
Destination represents a single travel location in the trip.

Each destination contains:
- Name of the destination
- Number of days the traveler will stay
- Budget allocated for that destination
- A list of activities planned for that location
*/
type Destination struct {
	Name       string   `json:"name"`
	Days       int      `json:"days"`
	Budget     float64  `json:"budget"`
	Activities []string `json:"activities"`
}

/*
Trip represents the entire travel itinerary.

It stores a slice of Destination structs which together
represent all locations included in the trip.
*/
type Trip struct {
	Destinations []Destination `json:"destinations"`
}

/*
Global trip variable that stores the current itinerary
during program execution.
*/
var trip Trip

/*
Buffered reader used to safely read user input
from the command line.
*/
var reader = bufio.NewReader(os.Stdin)

/*
Name of the JSON file used for persistent storage.
The program will automatically load from and save to this file.
*/
const fileName = "trip_data.json"

/*
main is the program entry point.

The program first loads any previously saved trip data
from the JSON file. After loading the data, it repeatedly
displays a menu and performs actions based on the user's
selection until the user chooses to exit.
*/
func main() {

	loadTrip()

	for {

		printMenu()

		choice := getIntInput("Choose option: ")

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
			saveTrip()
			return

		default:
			fmt.Println("Invalid menu choice. Please select 1-6.")
		}
	}
}

/*
printMenu displays the available menu options to the user.

This menu acts as the main navigation system for the program.
Users select a number corresponding to the feature they want
to use.
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
getStringInput safely reads a string from the user.

Parameters:
prompt - The message shown to the user before input.

Returns:
A trimmed string entered by the user.

Validation:
- Ensures the input is not empty
- Re-prompts the user if invalid input is provided
*/
func getStringInput(prompt string) string {

	for {

		fmt.Print(prompt)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		// Remove newline and surrounding whitespace
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
getIntInput reads and validates integer input from the user.

Parameters:
prompt - The message displayed to the user.

Returns:
A valid integer value.

Validation:
- Ensures the input is a valid integer
- Prompts the user again if conversion fails
*/
func getIntInput(prompt string) int {

	for {

		fmt.Print(prompt)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid whole number.")
			continue
		}

		return value
	}
}

/*
getFloatInput reads and validates floating-point input.

Parameters:
prompt - Message displayed before input.

Returns:
A valid float64 value.

Validation:
- Ensures the input can be converted into a floating-point number
*/
func getFloatInput(prompt string) float64 {

	for {

		fmt.Print(prompt)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		return value
	}
}

/*
isValidDestinationName checks whether a destination name
contains only alphabetic characters.

Parameters:
name - The destination name entered by the user.

Returns:
true if the name contains only letters, false otherwise.

This prevents invalid names that include numbers,
spaces, or special characters.
*/
func isValidDestinationName(name string) bool {

	if name == "" {
		return false
	}

	for _, char := range name {

		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}

/*
destinationExists checks if a destination with the same
name already exists in the itinerary.

Parameters:
name - Destination name to check.

Returns:
true if the destination already exists, otherwise false.

Case-insensitive comparison is used to avoid duplicates
like "Paris" and "paris".
*/
func destinationExists(name string) bool {

	for _, d := range trip.Destinations {

		if strings.EqualFold(d.Name, name) {
			return true
		}
	}

	return false
}

/*
addDestination allows the user to add a new destination
to the trip itinerary.

The function performs several validations:
- Destination name must contain only letters
- Destination must not already exist
- Days must be between 1 and 365
- Budget must be greater than zero and reasonable

After validation, the destination is appended to
the trip's destination list.
*/
func addDestination() {

	var name string

	for {

		name = getStringInput("Destination name (letters only): ")

		if !isValidDestinationName(name) {
			fmt.Println("Invalid name. Only letters allowed.")
			continue
		}

		if destinationExists(name) {
			fmt.Println("Destination already exists.")
			continue
		}

		break
	}

	var days int

	for {

		days = getIntInput("Number of days: ")

		if days <= 0 || days > 365 {
			fmt.Println("Days must be between 1 and 365.")
			continue
		}

		break
	}

	var budget float64

	for {

		budget = getFloatInput("Budget for this destination: ")

		if budget <= 0 {
			fmt.Println("Budget must be greater than 0.")
			continue
		}

		if budget > 1000000 {
			fmt.Println("Budget too large.")
			continue
		}

		break
	}

	// Create a new Destination struct
	dest := Destination{
		Name:       name,
		Days:       days,
		Budget:     budget,
		Activities: []string{},
	}

	// Append the destination to the trip slice
	trip.Destinations = append(trip.Destinations, dest)

	fmt.Println("Destination added.")
}

/*
addActivity allows the user to add an activity to an
existing destination.

Steps:
1. Verify at least one destination exists.
2. Display the destination list.
3. Ask the user which destination to modify.
4. Validate the activity name length.
5. Append the activity to the destination.
*/
func addActivity() {

	if len(trip.Destinations) == 0 {
		fmt.Println("Add a destination first.")
		return
	}

	viewDestinations()

	var index int

	for {

		index = getIntInput("Select destination number: ") - 1

		if index < 0 || index >= len(trip.Destinations) {
			fmt.Println("Invalid destination number.")
			continue
		}

		break
	}

	var activity string

	for {

		activity = getStringInput("Activity name: ")

		if len(activity) < 2 {
			fmt.Println("Activity name too short.")
			continue
		}

		if len(activity) > 50 {
			fmt.Println("Activity name too long.")
			continue
		}

		break
	}

	// Add the activity to the selected destination
	trip.Destinations[index].Activities =
		append(trip.Destinations[index].Activities, activity)

	fmt.Println("Activity added.")
}

/*
viewDestinations prints a numbered list of all
destinations currently in the trip.
*/
func viewDestinations() {

	fmt.Println("\nDestinations:")

	for i, d := range trip.Destinations {

		fmt.Printf("%d. %s\n", i+1, d.Name)
	}
}

/*
viewItinerary displays the complete trip plan.

For each destination, it prints:
- Destination name
- Number of days
- Budget
- All planned activities
*/
func viewItinerary() {

	if len(trip.Destinations) == 0 {
		fmt.Println("No destinations planned.")
		return
	}

	fmt.Println("\n====== Trip Itinerary ======")

	for _, d := range trip.Destinations {

		fmt.Println("\nDestination:", d.Name)
		fmt.Println("Days:", d.Days)
		fmt.Printf("Budget: %.2f\n", d.Budget)

		if len(d.Activities) == 0 {

			fmt.Println("Activities: None")

		} else {

			fmt.Println("Activities:")

			for _, a := range d.Activities {

				fmt.Println("-", a)
			}
		}
	}
}

/*
showBudget calculates and prints the total
budget for the entire trip.

The function iterates through all destinations
and sums their individual budgets.
*/
func showBudget() {

	total := 0.0

	for _, d := range trip.Destinations {

		total += d.Budget
	}

	fmt.Printf("\nTotal Trip Budget: %.2f\n", total)
}

/*
removeDestination deletes a destination from
the itinerary.

The user selects the destination by number.
The slice is updated to exclude the selected entry.
*/
func removeDestination() {

	if len(trip.Destinations) == 0 {
		fmt.Println("No destinations to remove.")
		return
	}

	viewDestinations()

	var index int

	for {

		index = getIntInput("Select destination number to remove: ") - 1

		if index < 0 || index >= len(trip.Destinations) {
			fmt.Println("Invalid destination number.")
			continue
		}

		break
	}

	removed := trip.Destinations[index].Name

	// Remove the destination using slice manipulation
	trip.Destinations = append(
		trip.Destinations[:index],
		trip.Destinations[index+1:]...,
	)

	fmt.Println("Destination removed:", removed)
}

/*
saveTrip writes the current trip data to the JSON file.

Steps:
1. Create or overwrite the JSON file.
2. Encode the Trip struct into JSON format.
3. Save the encoded data to the file.
*/
func saveTrip() {

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Error saving file.")
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(trip)

	if err != nil {
		fmt.Println("Error encoding data.")
		return
	}

	fmt.Println("Trip saved successfully.")
}

/*
loadTrip loads trip data from the JSON file if it exists.

If the file is corrupted or cannot be decoded,
the program safely resets the trip to an empty state
instead of crashing.
*/
func loadTrip() {

	file, err := os.Open(fileName)

	if err != nil {
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&trip)

	if err != nil {
		fmt.Println("Saved file corrupted. Starting new trip.")
		trip = Trip{}
	}
}
