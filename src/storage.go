package src

// Column size limits
const (
	ColumnUserNameSize = 32
	ColumnEmailSize    = 255
)

type Row struct {
	ID       uint32
	Username string
	Email    string
}

type Table struct {
	Rows []Row
}
