package minesweeper

import (
	"cmp"
	"slices"
	"testing"
)

func Test_adjecentCells(t *testing.T) {
	type Test struct {
		name  string
		game  Game
		coord CellCoordinate
		want  []CellCoordinate
	}

	newTest := func(name string, params Params, coord CellCoordinate, want []CellCoordinate) Test {
		return Test{
			name: name,
			game: Game{
				params: params,
			},
			coord: coord,
			want:  want,
		}
	}

	tests := []Test{
		newTest("zero", Params{Rows: 1, Cols: 1}, CellCoordinate{Row: 0, Col: 0}, []CellCoordinate{}),
		newTest("bottom right corner", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 4, Col: 4}, []CellCoordinate{{Row: 3, Col: 4}, {Row: 4, Col: 3}, {Row: 3, Col: 3}}),
		newTest("top left corner", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 0, Col: 0}, []CellCoordinate{{Row: 1, Col: 1}, {Row: 1, Col: 0}, {Row: 0, Col: 1}}),
		newTest("bottom middle", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 0, Col: 2}, []CellCoordinate{{Row: 1, Col: 1}, {Row: 1, Col: 2}, {Row: 1, Col: 3}, {Row: 0, Col: 1}, {Row: 0, Col: 3}}),
		newTest("top middle", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 4, Col: 2}, []CellCoordinate{{Row: 3, Col: 3}, {Row: 3, Col: 2}, {Row: 3, Col: 1}, {Row: 4, Col: 3}, {Row: 4, Col: 1}}),
		newTest("left middle", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 2, Col: 0}, []CellCoordinate{{Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}, {Row: 1, Col: 0}, {Row: 3, Col: 0}}),
		newTest("right middle", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 2, Col: 4}, []CellCoordinate{{Row: 1, Col: 3}, {Row: 2, Col: 3}, {Row: 3, Col: 3}, {Row: 1, Col: 4}, {Row: 3, Col: 4}}),
		newTest("middle", Params{Rows: 5, Cols: 5}, CellCoordinate{Row: 2, Col: 2}, []CellCoordinate{
			{Row: 1, Col: 1},
			{Row: 1, Col: 2},
			{Row: 1, Col: 3},
			{Row: 2, Col: 1},
			{Row: 2, Col: 3},
			{Row: 3, Col: 1},
			{Row: 3, Col: 2},
			{Row: 3, Col: 3},
		}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.game.adjecentCells(tt.coord); !isEqual(got, tt.want) {
				t.Errorf("crossAdjecentCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b []CellCoordinate) bool {
	if len(a) != len(b) {
		return false
	}
	sort := func(a, b CellCoordinate) int {
		if res := cmp.Compare(a.Row, b.Row); res != 0 {
			return res
		}

		return cmp.Compare(a.Col, b.Col)
	}
	slices.SortFunc(a, sort)
	slices.SortFunc(b, sort)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
