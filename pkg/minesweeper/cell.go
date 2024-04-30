package minesweeper

type CellValue int

const (
	Hidden CellValue = -2 + iota
	Mine
)

type Cell struct {
	value   CellValue
	pressed bool
}
