package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atabilog/logbook/internal/logbook"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Welcome to Logbook!\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  add   Add a new logbook entry\n")
		fmt.Fprintf(os.Stderr, "  list  List all entries\n")
		fmt.Fprintf(os.Stderr, "  search-tags  Find entries by tags\n")
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	subcommand := os.Args[1]

	if subcommand == "add" {
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		entry := addCmd.String("entry", "", "Create a new logbook entry")
		tags := addCmd.String("tags", "", "Optional comma-separated list of tags for the entry")

		addCmd.Parse(os.Args[2:])

		if *entry == "" {
			fmt.Println("To add a new entry please specify the entry via --entry option")
			return
		}
		fmt.Println("Adding a new entry")
		logbook.AddEntry(*entry, *tags)
	} else if subcommand == "list" {
		fmt.Println("Listing all entries")
		logbook.ListEntries()
	} else if subcommand == "search-tags" {
		logbook.SearchByTags(os.Args[2:])
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
