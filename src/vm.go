package src

import (
	"encoding/binary"
	"fmt"
)

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

// Get the first available page with space.
func executeInsert(statement Statement, table *Table) {
	rowData := statement.RowToInsert.Serialize()

	// First data
	if table.Pager.numPages == 0 {
		page, err := table.Pager.GetPage(0)
		if err != nil {
			fmt.Printf("Error fetching page: %v\n", err)
			return
		}
		InitPage(page)
		table.Pager.numPages = 1
	}

	lastPageNum := table.Pager.numPages - 1 // 0 indexed array
	page, err := table.Pager.GetPage(lastPageNum)
	if err != nil {
		fmt.Printf("Error fetching page: %v\n", err)
		return
	}

	err = InsertRow(page, rowData)
	// Full page
	if err != nil && err.Error() == "page is full" {
		newPageNum := table.Pager.numPages
		newPage, err := table.Pager.GetPage(newPageNum)
		if err != nil {
			fmt.Printf("Error fetching page: %v\n", err)
			return
		}
		InitPage(newPage)
		table.Pager.numPages++
		err = InsertRow(newPage, rowData)
		if err != nil {
			fmt.Printf("Error inserting into new page: %v\n", err)
			return
		}
	} else if nil != err {
		fmt.Printf("Error inserting row: %v\n", err)
	}
	fmt.Println("Executed.")

}

// Current implementation iterates through every page.
func executeSelect(table *Table) {
	for currentPage := uint32(0); currentPage < table.Pager.numPages; currentPage++ {
		page, err := table.Pager.GetPage(currentPage)
		if err != nil {
			fmt.Printf("Error fetching page %d: %v", currentPage, err)
			continue
		}

		numCells := binary.LittleEndian.Uint16(page[0:2])
		for slotNum := uint16(0); slotNum < numCells; slotNum++ {
			rowData, err := GetRow(page, slotNum)
			if err != nil {
				fmt.Printf("Error reading row %d on page %d: %v\n", slotNum, currentPage, err)
				continue
			}
			row := Deserialize(rowData)
			fmt.Printf("(%d, %s, %s)\n", row.ID, row.Name, row.Email)
		}
	}
}

func executeDelete(statement Statement, table *Table) {
	// targetID := statement.TargetID
	// indexToDelete := -1

	// for i, row := range table.Rows {
	// 	if row.ID == targetID {
	// 		indexToDelete = i
	// 		break
	// 	}
	// }
	// if indexToDelete == -1 {
	// 	fmt.Printf("Error: Row with ID %d not found.", targetID)
	// 	return
	// }

	// // Preserve table order deletion
	// table.Rows = append(table.Rows[:indexToDelete], table.Rows[indexToDelete+1:]...)
	// fmt.Printf("Deleted row with ID %d.\n", targetID)
}
