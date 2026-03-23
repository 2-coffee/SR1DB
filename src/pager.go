package src

import (
	"fmt"
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

}
