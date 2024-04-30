package main

import (
	"minesweeper/internal/tui"
	"minesweeper/pkg/minesweeper"
)

func main() {
	game, err := minesweeper.NewGame(minesweeper.Params{
		Rows:  10,
		Cols:  10,
		Mines: 10,
	})
	if err != nil {
		panic(err)
	}
	if err := tui.Run(&game); err != nil {
		panic(err)
	}
}
