package game

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Tick message for game updates
type tickMsg time.Time

type Model struct {
	game     *Game
	paused   bool
	tickRate time.Duration
}

func NewModel(game *Game) Model {
	return Model{
		game:     game,
		paused:   false,
		tickRate: time.Millisecond * 1000, // Adjust this value to control game speed
	}
}

func (m Model) Init() tea.Cmd {
	return tick(m.tickRate) // Start the game loop immediately
}

// Helper function to create a tick command
func tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "a":
			m.tickRate = m.tickRate / 2
		case "d":
			m.tickRate = m.tickRate * 2
		case "p": // Add pause functionality
			m.paused = !m.paused
			if m.paused {
				m.game.Pause()
			} else {
				m.game.Resume()
			}

			if !m.paused {
				return m, tick(m.tickRate)
			}
			return m, nil
		}

	case tickMsg:
		if !m.paused {
			// Run your game logic here
			m.game.Update() // You'll need to implement this
			return m, tick(m.tickRate)
		}
	}
	return m, nil
}

func (m Model) View() string {
	render := m.game.Render()
	speed := fmt.Sprintf("Speed: %s", m.tickRate)
	instructions := "Press 'a' to speed up, 'd' to slow down, 'p' to pause"

	return fmt.Sprintf("%s\n%s\n%s\n", render, speed, instructions)
}
