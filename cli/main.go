package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
	StatementDelete
)

type Statement struct {
	Type StatementType
}

func main() {
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

		// Handle exit command
		if strings.HasPrefix(input, ".") {
			if input == ".exit" {
				fmt.Println("Exiting database.")
				os.Exit(0)
			} else {
				fmt.Printf("Unrecognized command: '%s'\n", input)
			}
			continue
		}

		// Handle basic CRUD
		var statement Statement
		isPrepared := false

		lowerInput := strings.ToLower(input)
		if strings.HasPrefix(lowerInput, "insert") {
			statement.Type = StatementInsert
			isPrepared = true
		} else if strings.HasPrefix(lowerInput, "select") {
			statement.Type = StatementSelect
			isPrepared = true
		} else if strings.HasPrefix(lowerInput, "delete") {
			statement.Type = StatementDelete
			isPrepared = true
		} else {
			fmt.Printf("Syntax error: Unrecognized keyword at the start of '%s'.\n", lowerInput)
		}
		if isPrepared {
			executeStatement(statement)
		}
	}
	// Send it to execute in the backend

}

// Dummy backend function
func executeStatement(statement Statement) {
	switch statement.Type {
	case StatementInsert:
		fmt.Println("Executing an INSERT statement...")
	case StatementSelect:
		fmt.Println("Executing a SELECT statement...")
	case StatementDelete:
		fmt.Println("Executing a DELETE statement...")
	}
}
