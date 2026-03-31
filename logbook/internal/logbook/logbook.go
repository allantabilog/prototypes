package logbook

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var entriesDirectory string = "./entries"

type Entry struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	Tags      []string  `json:"tags"`
}

func SetEntriesDirectory(dir string) {
	entriesDirectory = dir
}

func saveEntry(entry Entry) error {
	jsonData, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %w", err)
	}

	filename := fmt.Sprintf("%s/entry_%d.json", entriesDirectory, entry.Timestamp.Unix())
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("Saved to: %s\n", filename)
	return nil
}

func AddEntry(entry string, tags string) error {
	var tagSlice []string
	if tags != "" {
		tagSlice = strings.Split(tags, ",")
	} else {
		tagSlice = []string{}
	}

	newEntry := Entry{
		Text:      entry,
		Timestamp: time.Now(),
		Tags:      tagSlice,
	}

	err := saveEntry(newEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("New entry: %s created on %s\n", entry, newEntry.Timestamp.Format("Monday, January 2, 2006 at 3:04 PM"))
	return nil
}

func ListEntries() error {
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

func SearchByTags(tags []string) error {
	fmt.Printf("Searching by tags: %+v", tags)
	return nil
}
