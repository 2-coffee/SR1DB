package main

import (
	"SR1DB/src"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	table, err := src.DbOpen("SR1DB.db")
	if err != nil {
		fmt.Printf("Failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer table.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("SR1DB > ")

		// Read input
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Handle meta commands
		if strings.HasPrefix(input, ".") {
			handleMetaCommand(input)
			continue
		}

		// Compile command
		statement, err := src.PrepareStatement(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Send command to the backend for execution
		src.ExecuteStatement(statement, table)

	}
}

func handleMetaCommand(input string) {
	if input == ".exit" {
		fmt.Println("Exiting database.")
		os.Exit(0)
	} else {
		fmt.Printf("Unrecognized command: '%s'\n", input)
	}
}
