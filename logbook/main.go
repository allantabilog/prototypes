package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var entriesDirectory string = "./entries"

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
	filename := fmt.Sprintf("%s/entry_%d.json", entriesDirectory, entry.Timestamp.Unix())
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
	// open the entries directory and print out the filename and contents of each file
	entries, err := os.ReadDir(entriesDirectory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(entriesDirectory, entry.Name())

		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", entry.Name(), err)
			continue
		}

		fmt.Printf("\n=== %s ===\n%s\n", entry.Name(), string(content))
	}

	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Welcome to Logbook!\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  add   Add a new logbook entry\n")
		fmt.Fprintf(os.Stderr, "  list  List all entries\n")
	}
	// Check if we have at least one argument (the subcommand)
	if len(os.Args) < 2 {
		flag.Usage()
		// flag.PrintDefaults()
		os.Exit(1)
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
		fmt.Println("Listing all entries")
		listEntries()
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
