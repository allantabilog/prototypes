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
	filename := fmt.Sprintf("entries/entry_%d.json", entry.Timestamp.Unix())
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("Saved to: %s\n", filename)
	return nil
}

func addEntry(entry string) error {
	newEntry := Entry{
		Text:      entry,
		Timestamp: time.Now(),
	}

	err := saveEntry(newEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("New entry: %s created on %s\n", entry, newEntry.Timestamp.Format("Monday, January 2, 2006 at 3:04 PM"))
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

	// Check if we have at least one argument (the subcommand)
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Get the subcommand
	subcommand := os.Args[1]

	// parse the commands
	if subcommand == "add" {
		// Create a new FlagSet for the 'add' subcommand
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		entry := addCmd.String("entry", "", "Create a new logbook entry")

		// Parse flags after the subcommand
		addCmd.Parse(os.Args[2:])

		if *entry == "" {
			fmt.Println("To add a new entry please specify the entry via --entry option")
			return
		}
		fmt.Println("Adding a new entry")
		addEntry(*entry)
	} else if subcommand == "list" {
		fmt.Println("Listing all entrieds")
	} else {
		usage()
		return
	}
}
