package minesweeper

import "errors"

var (
	ErrInvalidMinesAmount = errors.New("invalid amount of mines")
	ErrInvalidDimensions  = errors.New("invalid dimensions")
)

type Params struct {
	Rows  int
	Cols  int
	Mines int
}

func validateParams(params Params) error {
	if params.Rows < 1 || params.Cols < 1 {
		return ErrInvalidDimensions
	}
	size := params.Rows * params.Cols
	if params.Mines > size || params.Mines < 0 {
		return ErrInvalidMinesAmount
	}
	return nil
}
