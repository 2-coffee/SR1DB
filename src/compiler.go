package src

import (
	"fmt"
	"strconv"
	"strings"
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
	StatementDelete
)

type Statement struct {
	Type        StatementType
	RowToInsert Row
	TargetID    uint32
}

// For the CLI to use
func PrepareStatement(input string) (Statement, error) {
	var statement Statement
	tokens := tokenize(input)
	if len(tokens) == 0 {
		return statement, fmt.Errorf("empty statement")
	}
	command := strings.ToLower(tokens[0])

	switch command {
	case "insert":
		if len(tokens) != 4 {
			// Strict column struct for now without nulls
			return statement, fmt.Errorf("syntax error: expected 'insert <id> <username> <email>'")
		}
		statement.Type = StatementInsert

		// Parse ID: base 10 and 32 bits integer
		id, err := strconv.ParseUint(tokens[1], 10, 32)
		if err != nil {
			return statement, fmt.Errorf("syntax error: ID must be a number")
		}
		statement.RowToInsert.ID = uint32(id)
		// Insert name
		statement.RowToInsert.Username = tokens[2]
		// Insert email
		statement.RowToInsert.Email = tokens[3]

		return statement, nil
	case "select":
		statement.Type = StatementSelect
		return statement, nil
	case "delete":
		if len(tokens) != 2 {
			return statement, fmt.Errorf("syntax error: expected 'delete <id>'")
		}
		statement.Type = StatementDelete
		// Parse ID: base 10 and 32 bits integer
		id, err := strconv.ParseUint(tokens[1], 10, 32)
		if err != nil {
			return statement, fmt.Errorf("syntax error: ID must be a number")
		}
		// Let the vm handle this
		statement.TargetID = uint32(id)
		return statement, nil
	default:
		return statement, fmt.Errorf("syntax error: unrecognized command '%s'", command)
	}
}

// tokenize breaks a string into arguments; respects single and double quotes.
func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	inQuotes := false

	for _, char := range input {
		switch {
		case char == '\'' || char == '"':
			// Toggle quote states
			inQuotes = !inQuotes
		case char == ' ' && !inQuotes:
			// Empty space and no in quotes, then we are in a new arg
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
		default:
			// add characters as normal
			currentToken.WriteRune(char)
		}
	}
	// Last arg
	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
		currentToken.Reset()
	}
	return tokens
}
