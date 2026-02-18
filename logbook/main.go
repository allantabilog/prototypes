package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func saveEntry(entry Entry) error {
	// Convert to JSON
	jsonData, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %w", err)
	}

	// Save to file
	filename := fmt.Sprintf("entry_%d.json", entry.Timestamp.Unix())
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("Saved to: %s\n", filename)
	return nil
}

func addEntry(entry Entry) error {
	return nil
}

func listEntries() error {
	return nil
}

func usage() {
	fmt.Println("Usage text")
}

func main() {
	fmt.Println("Welcome to Logbook!")
	entry := flag.String("entry", "", "Create a new logbook entry")
	flag.Parse()
	fmt.Printf("Command line arguments: %v", flag.Args())

	// parse the commands
	if flag.Arg(0) == "add" {
		fmt.Println("Adding a new entry")
	} else if flag.Arg(0) == "list" {
		fmt.Println("Listing all entrieds")
	} else {
		usage()
		return
	}

	if *entry != "" {
		newEntry := Entry{
			Text:      *entry,
			Timestamp: time.Now(),
		}

		err := saveEntry(newEntry)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("New entry: %s created on %s\n", *entry, newEntry.Timestamp.Format("Monday, January 2, 2006 at 3:04 PM"))
	} else {
		fmt.Println("No new entry created.")
	}

}
