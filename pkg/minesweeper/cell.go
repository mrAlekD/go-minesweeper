package minesweeper

type CellValue int

const (
	Mine CellValue = -1
)

type Cell struct {
	value   CellValue
	pressed bool
}
