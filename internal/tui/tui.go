package tui

import (
	"fmt"
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
	game *minesweeper.Game
}

func Run(game *minesweeper.Game) error {
	p := tea.NewProgram(NewModel(game), tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err := p.Run()
	return err
}

func NewModel(game *minesweeper.Game) *Model {
	return &Model{game}
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
			return m, nil
		}
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionRelease {
			row := (msg.Y - 1) / cellHeight
			col := msg.X / cellWidth
			coord := minesweeper.CellCoordinate{Row: row, Col: col}
			switch msg.Button {
			case tea.MouseButtonLeft:
				m.game.Press(coord)
			case tea.MouseButtonRight:
				m.game.ToggleMark(coord)
			}
		}
		return m, nil
	}
	return m, nil
}

var states = []string{"init", "running", "won", "lost", "error"}

func (m *Model) View() string {
	params := m.game.Params()
	marked := m.game.Marked()
	rows := make([]string, 0, params.Rows+1)
	rows = append(rows, fmt.Sprintf("State: %s, marked: %d, mines: %d", states[m.game.State()], marked, params.Mines))
	cells := m.game.Cells()
	for _, row := range cells {
		items := make([]string, params.Cols)
		for c, cell := range row {
			switch cell {
			case minesweeper.Mine:
				items[c] = mineCellStyle.Render("M")
			case minesweeper.Marked:
				items[c] = markedCellStyle.Render("X")
			case minesweeper.Hidden:
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
