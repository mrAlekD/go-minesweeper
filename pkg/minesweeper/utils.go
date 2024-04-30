package minesweeper

type CellCoordinate struct {
	Row int
	Col int
}

func (g Game) getRealIdx(coord CellCoordinate) int {
	return coord.Row*g.params.Cols + coord.Col
}

func (g Game) getCellCoordinate(idx int) CellCoordinate {
	return CellCoordinate{
		Row: idx / g.params.Cols,
		Col: idx % g.params.Cols,
	}
}

func (g Game) crossAdjecentCells(coord CellCoordinate) []CellCoordinate {
	adjCells := make([]CellCoordinate, 0, 4)
	if coord.Row > 0 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row - 1, Col: coord.Col})
	}
	if coord.Row < g.params.Rows-1 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row + 1, Col: coord.Col})
	}
	if coord.Col > 0 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row, Col: coord.Col - 1})
	}
	if coord.Col < g.params.Cols-1 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row, Col: coord.Col + 1})
	}
	return adjCells
}

func (g Game) adjecentCells(coord CellCoordinate) []CellCoordinate {
	adjCells := g.crossAdjecentCells(coord)
	rowGt0 := coord.Row > 0
	colGt0 := coord.Col > 0
	rowLtRows := coord.Row < g.params.Rows-1
	colLtCols := coord.Col < g.params.Cols-1

	if rowGt0 && colGt0 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row - 1, Col: coord.Col - 1})
	}
	if rowGt0 && colLtCols {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row - 1, Col: coord.Col + 1})
	}
	if rowLtRows && colGt0 {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row + 1, Col: coord.Col - 1})
	}
	if rowLtRows && colLtCols {
		adjCells = append(adjCells, CellCoordinate{Row: coord.Row + 1, Col: coord.Col + 1})
	}
	return adjCells
}
