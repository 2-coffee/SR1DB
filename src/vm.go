package src

import "fmt"

// Used in the CLI to trigger the VM
func ExecuteStatement(statement Statement, table *Table) {
	switch statement.Type {
	case StatementInsert:
		executeInsert(statement, table)
	case StatementSelect:
		// currently selects all
		executeSelect(table)
	case StatementDelete:
		executeDelete(statement, table)
	}
}

func executeInsert(statement Statement, table *Table) {
	table.Rows = append(table.Rows, statement.RowToInsert)
	fmt.Println("Executed.")
}

func executeSelect(table *Table) {
	// TODO: filter
	// Current selects all
	for _, row := range table.Rows {
		fmt.Printf("(%d, %s, %s)\n", row.ID, row.Username, row.Email)
	}
}

func executeDelete(statement Statement, table *Table) {
	targetID := statement.TargetID
	indexToDelete := -1

	for i, row := range table.Rows {
		if row.ID == targetID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete == -1 {
		fmt.Printf("Error: Row with ID %d not found.", targetID)
		return
	}

	// Preserve table order deletion
	table.Rows = append(table.Rows[:indexToDelete], table.Rows[indexToDelete+1:]...)
	fmt.Printf("Deleted row with ID %d.\n", targetID)
}
