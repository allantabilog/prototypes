package logbook

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAddEntry(t *testing.T) {
	tempDir := t.TempDir()
	SetEntriesDirectory(tempDir)

	tests := []struct {
		name     string
		entry    string
		tags     string
		wantTags []string
		wantErr  bool
	}{
		{
			name:     "entry with tags",
			entry:    "Test entry",
			tags:     "work,important",
			wantTags: []string{"work", "important"},
			wantErr:  false,
		},
		{
			name:     "entry without tags",
			entry:    "Test entry no tags",
			tags:     "",
			wantTags: []string{},
			wantErr:  false,
		},
		{
			name:     "entry with single tag",
			entry:    "Single tag entry",
			tags:     "personal",
			wantTags: []string{"personal"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddEntry(tt.entry, tt.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			entries, err := os.ReadDir(tempDir)
			if err != nil {
				t.Fatalf("Failed to read test directory: %v", err)
			}

			if len(entries) == 0 {
				t.Fatal("No entry file was created")
			}

			lastEntry := entries[len(entries)-1]
			data, err := os.ReadFile(filepath.Join(tempDir, lastEntry.Name()))
			if err != nil {
				t.Fatalf("Failed to read entry file: %v", err)
			}

			var savedEntry Entry
			if err := json.Unmarshal(data, &savedEntry); err != nil {
				t.Fatalf("Failed to unmarshal entry: %v", err)
			}

			if savedEntry.Text != tt.entry {
				t.Errorf("Text = %v, want %v", savedEntry.Text, tt.entry)
			}

			if len(savedEntry.Tags) != len(tt.wantTags) {
				t.Errorf("Tags length = %v, want %v", len(savedEntry.Tags), len(tt.wantTags))
			}

			for i, tag := range savedEntry.Tags {
				if tag != tt.wantTags[i] {
					t.Errorf("Tag[%d] = %v, want %v", i, tag, tt.wantTags[i])
				}
			}
		})
	}
}

func TestSaveEntry(t *testing.T) {
	tempDir := t.TempDir()
	SetEntriesDirectory(tempDir)

	entry := Entry{
		Text:      "Test entry",
		Timestamp: time.Now(),
		Tags:      []string{"test", "example"},
	}

	err := saveEntry(entry)
	if err != nil {
		t.Fatalf("saveEntry() error = %v", err)
	}

	entries, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read test directory: %v", err)
	}

	if len(entries) != 1 {
		t.Errorf("Expected 1 file, got %d", len(entries))
	}

	data, err := os.ReadFile(filepath.Join(tempDir, entries[0].Name()))
	if err != nil {
		t.Fatalf("Failed to read entry file: %v", err)
	}

	var savedEntry Entry
	if err := json.Unmarshal(data, &savedEntry); err != nil {
		t.Fatalf("Failed to unmarshal entry: %v", err)
	}

	if savedEntry.Text != entry.Text {
		t.Errorf("Text = %v, want %v", savedEntry.Text, entry.Text)
	}

	if len(savedEntry.Tags) != len(entry.Tags) {
		t.Errorf("Tags length = %v, want %v", len(savedEntry.Tags), len(entry.Tags))
	}
}

func TestListEntries(t *testing.T) {
	tempDir := t.TempDir()
	SetEntriesDirectory(tempDir)

	testEntries := []Entry{
		{Text: "Entry 1", Timestamp: time.Now(), Tags: []string{"tag1"}},
		{Text: "Entry 2", Timestamp: time.Now(), Tags: []string{}},
	}

	for _, entry := range testEntries {
		if err := saveEntry(entry); err != nil {
			t.Fatalf("Failed to save test entry: %v", err)
		}
	}

	err := ListEntries()
	if err != nil {
		t.Errorf("ListEntries() error = %v", err)
	}
}
