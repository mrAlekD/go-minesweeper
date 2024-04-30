package tui

import (
	"minesweeper/pkg/minesweeper"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var defaultColor = lipgloss.Color("247")
var markedColor = lipgloss.Color("47")
var shownColor = lipgloss.Color("39")
var emptyColor = lipgloss.Color("241")
var mineColor = lipgloss.Color("160")
var cellStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(0, 1).Foreground(defaultColor).BorderForeground(defaultColor)
var markedCellStyle = cellStyle.Copy().Foreground(markedColor).BorderForeground(markedColor)
var shownCellStyle = cellStyle.Copy().Foreground(shownColor).BorderForeground(shownColor)
var emptyCellStyle = cellStyle.Copy().Foreground(emptyColor).BorderForeground(emptyColor)
var mineCellStyle = cellStyle.Copy().Foreground(mineColor).BorderForeground(mineColor)
var cellWidth = cellStyle.GetHorizontalFrameSize() + 1
var cellHeight = cellStyle.GetVerticalFrameSize() + 1

type Model struct {
	game   *minesweeper.Game
	marked map[minesweeper.CellCoordinate]bool
}

func Run(game *minesweeper.Game) error {
	p := tea.NewProgram(NewModel(game), tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err := p.Run()
	return err
}

func NewModel(game *minesweeper.Game) *Model {
	return &Model{game, make(map[minesweeper.CellCoordinate]bool)}
}

func (m *Model) Init() tea.Cmd {
	m.game.Init()
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			m.game.Init()
			m.marked = make(map[minesweeper.CellCoordinate]bool)
			return m, nil
		}
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionRelease {
			row := (msg.Y - 1) / cellHeight
			col := msg.X / cellWidth
			switch msg.Button {
			case tea.MouseButtonLeft:
				m.game.Press(minesweeper.CellCoordinate{Row: row, Col: col})
			case tea.MouseButtonRight:
				coord := minesweeper.CellCoordinate{Row: row, Col: col}
				m.marked[coord] = !m.marked[coord]
			}
		}
		return m, nil
	}
	return m, nil
}

var states = []string{"init", "running", "won", "lost", "error"}

func (m *Model) View() string {
	rows := make([]string, 0, m.game.Params().Rows+1)
	rows = append(rows, "State: "+states[m.game.State()])
	cells := m.game.Cells()
	for r, row := range cells {
		items := make([]string, m.game.Params().Cols)
		for c, cell := range row {
			switch cell {
			case minesweeper.Mine:
				items[c] = mineCellStyle.Render("M")
			case minesweeper.Hidden:
				coord := minesweeper.CellCoordinate{Row: r, Col: c}
				if m.marked[coord] {
					items[c] = markedCellStyle.Render("X")
					continue
				}
				items[c] = cellStyle.Render("?")
			case 0:
				items[c] = emptyCellStyle.Render(" ")
			default:
				items[c] = shownCellStyle.Render(strconv.Itoa(int(cell)))
			}
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, items...))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
