package main

import (
	"flag"
	"minesweeper/internal/tui"
	"minesweeper/pkg/minesweeper"
)

func main() {
	game, err := minesweeper.NewGame(parseParams())
	if err != nil {
		panic(err)
	}
	if err := tui.Run(&game); err != nil {
		panic(err)
	}
}

func parseParams() minesweeper.Params {
	params := minesweeper.Params{
		Rows:  10,
		Cols:  10,
		Mines: 10,
	}

	flag.IntVar(&params.Rows, "rows", params.Rows, "number of rows")
	flag.IntVar(&params.Cols, "cols", params.Cols, "number of columns")
	flag.IntVar(&params.Mines, "mines", params.Mines, "number of mines")
	flag.Parse()
	return params
}
