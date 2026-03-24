package src

import (
	"encoding/binary"
	"fmt"
)

const (
	PageHeaderSize = 4
	SlotSize       = 4
)

// InitPage sets up a new page of 4KB
func InitPage(page []byte) {
	// Number of cells
	binary.LittleEndian.PutUint16(page[0:2], 0)

	// The convention for free space starts at the end of the page
	// Data of each row grows up until it meets cells slots
	binary.LittleEndian.PutUint16(page[2:4], uint16(PageSize))
}

// Insert a new row to an existing page
// Each page has a header of 4 bytes which contains a counter for number of cells and a freeSpacePtr,
// using 2 bytes each.
// Slots begin immediately after the header, and slots are also 4 bytes each.
// Gives row ptr and length of row.
// Actual row data begins from the end of the page.
func InsertRow(page []byte, rowData []byte) error {
	numCells := binary.LittleEndian.Uint16(page[0:2])
	freeSpacePtr := binary.LittleEndian.Uint16(page[2:4])
	rowLen := uint16(len(rowData))

	// Check if there is enough space in this page
	slotsEnd := PageHeaderSize + numCells*SlotSize
	spaceNeeded := rowLen + SlotSize

	if freeSpacePtr-slotsEnd < spaceNeeded {
		return fmt.Errorf("page is full")
	}

	// Add row and meta data to page
	newFreeSpacePtr := freeSpacePtr - rowLen
	numberCopied := copy(page[newFreeSpacePtr:freeSpacePtr], rowData)
	if numberCopied != int(rowLen) { // visit this again
		return fmt.Errorf("Did not successfully copy data")
	}
	// Slots update
	binary.LittleEndian.PutUint16(page[slotsEnd:slotsEnd+2], newFreeSpacePtr)
	binary.LittleEndian.PutUint16(page[slotsEnd+2:slotsEnd+4], rowLen)
	// Header update
	binary.LittleEndian.PutUint16(page[0:2], numCells+1)
	binary.LittleEndian.PutUint16(page[2:4], freeSpacePtr)

	return nil
}

func GetRow(page []byte, slotNum uint16) ([]byte, error) {
	numCells := binary.LittleEndian.Uint16(page[0:2])
	if slotNum >= numCells {
		return nil, fmt.Errorf("slot number: %d is out of bounds", slotNum)
	}

	// Find the slot
	cellOffset := PageHeaderSize + slotNum*SlotSize
	rowPtr := binary.LittleEndian.Uint16(page[cellOffset : cellOffset+2])
	rowLen := binary.LittleEndian.Uint16(page[cellOffset+2 : cellOffset+4])

	rowData := make([]byte, rowLen)
	copy(rowData, page[rowPtr:rowPtr+rowLen])

	return rowData, nil
}
