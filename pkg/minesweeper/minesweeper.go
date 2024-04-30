package minesweeper

import "math/rand"

type GameState int

const (
	Init GameState = iota
	Running
	Won
	Lost
	Error
)

type Game struct {
	params Params
	board  []Cell
	state  GameState
}

func NewGame(params Params) (Game, error) {
	if err := validateParams(params); err != nil {
		return Game{}, err
	}

	return Game{params: params}, nil
}

func (g *Game) Init() {
	g.state = Init
	g.board = make([]Cell, g.params.Rows*g.params.Cols)
	g.addMines()
	g.applyCounts()
	g.state = Running
}

func (g *Game) Press(coord CellCoordinate) {
	if g.state != Running {
		return
	}

	idx := g.getRealIdx(coord)
	if g.board[idx].pressed {
		return
	}

	g.board[idx].pressed = true
	if g.board[idx].value == Mine {
		g.state = Lost
		return
	}

	if g.board[idx].value == 0 {
		adj := g.adjecentCells(coord)
		for _, c := range adj {
			g.Press(c)
		}
	}

	if g.isWon() {
		g.state = Won
	}
}

func (g Game) isWon() bool {
	for i := range g.board {
		if g.board[i].value != Mine && !g.board[i].pressed {
			return false
		}
	}
	return g.state == Running // cuts off Loses and weird configrations, may change later
}

func (g Game) State() GameState {
	return g.state
}

func (g *Game) addMines() {
	if g.state != Init {
		return
	}

	added := 0
	for added < g.params.Mines {
		idx := rand.Intn(g.params.Rows * g.params.Cols)
		if g.board[idx].value != Mine {
			g.board[idx].value = Mine
			added++
		}
	}
}

func (g *Game) applyCounts() {
	if g.state != Init {
		return
	}

	for i := range g.board {
		if g.board[i].value == Mine {
			coord := g.getCellCoordinate(i)
			adj := g.adjecentCells(coord)
			for _, c := range adj {
				if g.board[g.getRealIdx(c)].value != Mine {
					g.board[g.getRealIdx(c)].value++
				}
			}
		}
	}
}
