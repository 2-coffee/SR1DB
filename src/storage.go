package src

import (
	"bytes"
	"encoding/binary"
)

// Column size limits
const (
	ColumnUserNameSize = 32
	ColumnEmailSize    = 255
	PageSize           = 4096 // 4KB OS page
)

type Row struct {
	ID    uint32
	Name  string
	Email string
}

type Table struct {
	Rows []Row
}

func (r *Row) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Write ID (4 bytes)
	// Might be useful to use/research more on nativeEndian
	binary.Write(buf, binary.LittleEndian, r.ID)

	// Write name
	nameLen := uint8(len(r.Name))
	binary.Write(buf, binary.LittleEndian, nameLen)
	buf.WriteString(r.Name)

	// Write email
	emailLen := uint8(len(r.Email))
	binary.Write(buf, binary.LittleEndian, emailLen)
	buf.WriteString(r.Email)
	return buf.Bytes()
}

func Deserialize(data []byte) Row {
	var row Row
	buf := bytes.NewReader(data)

	// Read ID
	binary.Read(buf, binary.LittleEndian, &row.ID) // 3rd argument must a be a pointer to a fixed size value

	// Read Name
	var nameLen uint8
	binary.Read(buf, binary.LittleEndian, &nameLen)
	nameBytes := make([]byte, nameLen)
	buf.Read(nameBytes)
	row.Name = string(nameBytes) // turn the bytes in to a string

	// Read email
	var emailLen uint8
	binary.Read(buf, binary.LittleEndian, &emailLen)
	emailBytes := make([]byte, emailLen)
	buf.Read(emailBytes)
	row.Email = string(emailBytes)

	return row
}
