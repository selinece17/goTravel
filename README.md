# Go Travel Planner

## Overview

**Go Travel Planner** is a console-based travel itinerary planner
written in **Go (Golang)**. The program allows users to build and manage
a travel itinerary directly from the command line. Users can add
destinations, assign a number of days and a budget to each destination,
and attach activities to each location.

This project demonstrates core programming concepts such as:

-   Structs and structured data
-   Input validation
-   Dynamic lists using slices
-   File persistence using JSON
-   Command line interfaces (CLI)
-   Modular program design
-   CRUD-style operations (Create, Read, Update, Delete)

The program saves trip data in a JSON file (`trip_data.json`) so that
the itinerary persists between program runs.

Even if the program is closed, it will automatically reload the saved
trip data the next time it starts.

------------------------------------------------------------------------

# Features

### Add Destination

Users can create destinations in their itinerary by specifying:

-   Destination name (letters only)
-   Number of days staying
-   Budget allocated for that destination

The program validates all inputs to prevent invalid data.

------------------------------------------------------------------------

### Add Activity

Users can add activities to a destination such as:

-   Museums
-   Restaurants
-   Tours
-   Hiking
-   Sightseeing

Each destination maintains its own list of activities.

------------------------------------------------------------------------

### View Itinerary

Displays the complete travel plan including:

-   Destination names
-   Days at each destination
-   Budget allocation
-   All activities

------------------------------------------------------------------------

### View Budget Summary

Calculates the **total budget for the entire trip**.

------------------------------------------------------------------------

### Remove Destination

Allows users to delete destinations safely from the itinerary.

------------------------------------------------------------------------

### Persistent Storage

Trip data is stored in:

    trip_data.json

The file automatically saves the trip when exiting the program and loads
it when the program starts.

------------------------------------------------------------------------

# Project Structure

    goTravel/
    │
    ├── main.go
    ├── go.mod
    ├── trip_data.json
    └── README.md

### File Descriptions

**main.go**

Contains the entire application including:

-   Menu system
-   Input validation
-   Trip management logic
-   JSON save/load functionality
-   Struct definitions

------------------------------------------------------------------------

**go.mod**

Defines the Go module used by the project and manages dependencies.

------------------------------------------------------------------------

**trip_data.json**

Stores saved trip data.

Example:

``` json
{
  "destinations": [
    {
      "name": "Paris",
      "days": 4,
      "budget": 1200,
      "activities": ["Museum", "Cafe", "Walking Tour"]
    }
  ]
}
```

------------------------------------------------------------------------

# Requirements

You must have **Go installed** to run this project.

### Check if Go is Installed

Open a terminal and run:

    go version

Example output:

    go version go1.22 darwin/amd64

If Go is not installed, download it here:

https://go.dev/dl/

Installation instructions are provided for:

-   macOS
-   Windows
-   Linux

------------------------------------------------------------------------

# How to Download (Clone) the Project from GitHub

### Step 1 --- Install Git (if needed)

Check if Git is installed:

    git --version

Example output:

    git version 2.43.0

If Git is not installed:

Download it here:

https://git-scm.com/downloads

------------------------------------------------------------------------

### Step 2 --- Copy the Repository URL

On GitHub, click the **Code** button and copy the repository URL.

Example:

    https://github.com/selinece17/goTravel.git

------------------------------------------------------------------------

### Step 3 --- Open Terminal

**macOS** - Press Command + Space - Type Terminal - Press Enter

**Windows** - Open Command Prompt or PowerShell

**Linux** - Open your terminal application

------------------------------------------------------------------------

### Step 4 --- Navigate to Where You Want the Project

Example:

    cd Desktop

------------------------------------------------------------------------

### Step 5 --- Clone the Repository

    git clone https://github.com/selinece17/goTravel.git

This will create a new folder called:

    goTravel

------------------------------------------------------------------------

### Step 6 --- Move Into the Project Directory

    cd goTravel

------------------------------------------------------------------------

### Step 7 --- Verify the Files

    ls

You should see:

    go.mod
    main.go
    trip_data.json
    README.md

------------------------------------------------------------------------

# Running the Program

Once inside the project folder, run:

    go run main.go

You should see:

    ====== Travel Planner ======
    1. Add Destination
    2. Add Activity
    3. View Itinerary
    4. View Budget Summary
    5. Remove Destination
    6. Exit
    ============================

------------------------------------------------------------------------

# Example Usage

### Adding a Destination

    Choose option: 1
    Destination name (letters only): Paris
    Number of days: 4
    Budget for this destination: 1500
    Destination added.

------------------------------------------------------------------------

### Adding an Activity

    Choose option: 2
    Select destination number: 1
    Activity name: Louvre Museum
    Activity added.

------------------------------------------------------------------------

### Viewing the Itinerary

    ====== Trip Itinerary ======

    Destination: Paris
    Days: 4
    Budget: 1500.00
    Activities:
    - Louvre Museum

------------------------------------------------------------------------

### Viewing Budget

    Total Trip Budget: 1500.00

------------------------------------------------------------------------

# Programming Concepts Demonstrated

This project demonstrates several important Go concepts:

### Structs

Structured data types representing destinations and the trip.

### Slices

Dynamic lists storing destinations and activities.

### JSON Encoding/Decoding

Using the `encoding/json` package to persist data.

### Input Validation

Ensures users enter valid information.

### Command Line Interfaces

The entire application runs in the terminal.


------------------------------------------------------------------------

# License

This project is intended for educational purposes.
