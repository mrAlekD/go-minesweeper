package minesweeper

type CellValue int

const (
	Hidden CellValue = -3 + iota
	Marked
	Mine
)

type Cell struct {
	value   CellValue
	pressed bool
	marked  bool
}
