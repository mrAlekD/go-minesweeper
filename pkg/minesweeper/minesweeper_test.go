package minesweeper

import (
	"reflect"
	"testing"
)

func Test_NewGame(t *testing.T) {
	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name: "zero",
			params: Params{
				Rows:  0,
				Cols:  0,
				Mines: 0,
			},
			err: ErrInvalidDimensions,
		},
		{
			name: "negative",
			params: Params{
				Rows:  1,
				Cols:  1,
				Mines: -1,
			},
			err: ErrInvalidMinesAmount,
		},
		{
			name: "too many",
			params: Params{
				Rows:  1,
				Cols:  1,
				Mines: 2,
			},
			err: ErrInvalidMinesAmount,
		},
		{
			name: "ok",
			params: Params{
				Rows:  1,
				Cols:  1,
				Mines: 0,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewGame(tt.params); err != tt.err {
				t.Errorf("NewGame() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func Test_Init(t *testing.T) {
	type Test struct {
		name string
		game Game
		want struct {
			boardSize int
			mines     int
		}
	}

	newTest := func(name string, params Params) Test {
		return Test{
			name: name,
			game: Game{
				params: params,
			},
			want: struct {
				boardSize int
				mines     int
			}{
				boardSize: params.Rows * params.Cols,
				mines:     params.Mines,
			},
		}
	}

	tests := []Test{
		newTest("usual", Params{Rows: 5, Cols: 5, Mines: 5}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.game.Init()
			if tt.game.state != Running {
				t.Errorf("state = %v, want %v", tt.game.state, "Init")
			}
			if countPressed(tt.game.board) != 0 {
				t.Errorf("pressed cells = %d, want %d", countPressed(tt.game.board), 0)
			}
			if len(tt.game.board) != tt.want.boardSize {
				t.Errorf("board size = %d, want %d", len(tt.game.board), tt.want.boardSize)
			}
			if countMines(tt.game.board) != tt.want.mines {
				t.Errorf("mines = %d, want %d", countMines(tt.game.board), tt.want.mines)
			}
		})
	}
}

func Test_counts(t *testing.T) {
	type Board struct {
		rows  int
		cols  int
		cells []Cell
	}
	type Test struct {
		name  string
		board Board
		want  []Cell
	}

	newTestBoard := func(rows, cols int, mines []CellCoordinate) Board {
		board := make([]Cell, rows*cols)
		for _, coord := range mines {
			board[coord.Row*cols+coord.Col].value = Mine
		}
		return Board{rows: rows, cols: cols, cells: board}
	}

	tests := []Test{
		{
			name:  "2x2, 1 mine",
			board: newTestBoard(2, 2, []CellCoordinate{{Row: 0, Col: 1}}),
			want: []Cell{
				{value: 1},
				{value: Mine},
				{value: 1},
				{value: 1},
			},
		},
		{
			name:  "3x3, 1 mine",
			board: newTestBoard(3, 3, []CellCoordinate{{Row: 0, Col: 1}}),
			want: []Cell{
				{value: 1},
				{value: Mine},
				{value: 1},
				{value: 1},
				{value: 1},
				{value: 1},
				{value: 0},
				{value: 0},
				{value: 0},
			},
		},
		{
			name:  "3x3, 2 mines",
			board: newTestBoard(3, 3, []CellCoordinate{{Row: 0, Col: 2}, {Row: 1, Col: 0}}),
			want: []Cell{
				{value: 1},
				{value: 2},
				{value: Mine},
				{value: Mine},
				{value: 2},
				{value: 1},
				{value: 1},
				{value: 1},
				{value: 0},
			},
		},
		{
			name:  "3x3, 3 mines",
			board: newTestBoard(3, 3, []CellCoordinate{{Row: 0, Col: 2}, {Row: 1, Col: 0}, {Row: 2, Col: 2}}),
			want: []Cell{
				{value: 1},
				{value: 2},
				{value: Mine},
				{value: Mine},
				{value: 3},
				{value: 2},
				{value: 1},
				{value: 2},
				{value: Mine},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := Game{board: tt.board.cells, params: Params{Rows: tt.board.rows, Cols: tt.board.cols, Mines: countMines(tt.board.cells)}}
			game.applyCounts()
			if !reflect.DeepEqual(game.board, tt.want) {
				t.Errorf("board = %v, want %v", game.board, tt.want)
			}
		})
	}
}

func Test_Press(t *testing.T) {
	type Test struct {
		name      string
		params    Params
		mines     []CellCoordinate
		press     CellCoordinate
		wantBoard []Cell
		wantState GameState
	}

	tests := []Test{
		{
			name: "3x3, press mine",
			params: Params{
				Rows:  3,
				Cols:  3,
				Mines: 1,
			},
			mines: []CellCoordinate{{Row: 2, Col: 2}},
			press: CellCoordinate{Row: 2, Col: 2},
			wantBoard: []Cell{
				{value: 0},
				{value: 0},
				{value: 0},
				{value: 0},
				{value: 1},
				{value: 1},
				{value: 0},
				{value: 1},
				{value: Mine, pressed: true},
			},
			wantState: Lost,
		},
		{
			name: "3x3, press a nonzero number",
			params: Params{
				Rows:  3,
				Cols:  3,
				Mines: 1,
			},
			mines: []CellCoordinate{{Row: 2, Col: 2}},
			press: CellCoordinate{Row: 2, Col: 1},
			wantBoard: []Cell{
				{value: 0},
				{value: 0},
				{value: 0},
				{value: 0},
				{value: 1},
				{value: 1},
				{value: 0},
				{value: 1, pressed: true},
				{value: Mine},
			},
			wantState: Running,
		},
		{
			name: "3x3, press a zero",
			params: Params{
				Rows:  3,
				Cols:  3,
				Mines: 1,
			},
			mines: []CellCoordinate{{Row: 2, Col: 2}},
			press: CellCoordinate{Row: 2, Col: 0},
			wantBoard: []Cell{
				{value: 0, pressed: true},
				{value: 0, pressed: true},
				{value: 0, pressed: true},
				{value: 0, pressed: true},
				{value: 1, pressed: true},
				{value: 1, pressed: true},
				{value: 0, pressed: true},
				{value: 1, pressed: true},
				{value: Mine},
			},
			wantState: Won,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame(tt.params)
			if err != nil {
				t.Error(err)
			}

			game.board = make([]Cell, tt.params.Rows*tt.params.Cols)
			for _, mine := range tt.mines {
				game.board[game.getRealIdx(mine)] = Cell{value: Mine}
			}
			game.applyCounts()
			game.state = Running
			game.Press(tt.press)
			if !reflect.DeepEqual(game.board, tt.wantBoard) {
				t.Errorf("board = %v, want %v", game.board, tt.wantBoard)
			}
			if game.state != tt.wantState {
				t.Errorf("state = %v, want %v", game.state, tt.wantState)
			}
		})
	}
}

func countPressed(board []Cell) int {
	var count int
	for _, cell := range board {
		if cell.pressed {
			count++
		}
	}
	return count
}

func countMines(board []Cell) int {
	var count int
	for _, cell := range board {
		if cell.value == Mine {
			count++
		}
	}
	return count
}
