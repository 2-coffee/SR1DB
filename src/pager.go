package src

import (
	"fmt"
	"io"
	"os"
)

const (
	TableMaxPages = 100 // arbitrary
)

type Pager struct {
	file       *os.File
	fileLength int64
	pages      [TableMaxPages][]byte
	numPages   uint32 // current number of pages
}

// Initializes a connection to the db on disk
// This function is used in storage.go by DbOpen
func PagerOpen(filename string) (*Pager, error) {
	// Open the file. If it doesn't exist, create it.
	// 0600 means only the current user can read/write to the file
	file, err := os.OpenFile(filename, 0600, os.FileMode(os.O_RDWR))
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("unable to stat file: %v", err)
	}

	fileLength := stat.Size()
	numPages := uint32(fileLength / PageSize)

	return &Pager{
		file:       file,
		fileLength: fileLength,
		numPages:   numPages,
	}, nil
}

// GetPage fetches a 4KB page. If it is not already in RAM, then we fetch it from disk
func (p *Pager) GetPage(pageNumber uint32) ([]byte, error) {
	if pageNumber >= TableMaxPages {
		return nil, fmt.Errorf("page number %d is out of bounds", pageNumber)
	}

	// In RAM
	if p.pages[pageNumber] != nil {
		return p.pages[pageNumber], nil
	}

	// Cache miss; allocate 4KB and fetch from disk
	page := make([]byte, PageSize)

	if pageNumber < p.numPages {
		// ReadAt func needs an offset in int64
		offset := int64(pageNumber) * PageSize
		_, err := p.file.ReadAt(page, offset)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading file at offset %d: %v", offset, err)
		}
	}
	// Not sure if this pageNumber offset is correct due to 0-indexed
	// Need to revisit this
	p.pages[pageNumber] = page

	return page, nil
}

func (p *Pager) Flush(pageNumber uint32) error {
	// Page not in RAM
	if p.pages[pageNumber] == nil {
		return nil
	}
	// Calculate the offset of the page for the file
	offset := int64(pageNumber * p.numPages)
	_, err := p.file.WriteAt(p.pages[pageNumber], offset)
	return err
}
