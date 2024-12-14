package game

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/rand"
)

// Color palette - from cool to warm colors
var colors = []string{
	"#0000FF", // Deep Blue
	"#0033CC", // Royal Blue
	"#0066FF", // Bright Blue
	"#0099FF", // Sky Blue
	"#00CCFF", // Cyan
	"#00FFCC", // Turquoise
	"#00FF99", // Sea Green
	"#00FF66", // Spring Green
	"#00FF00", // Pure Green
	"#33FF00", // Lime Green
	"#66FF00", // Yellow Green
	"#99FF00", // Chartreuse
	"#CCFF00", // Bright Yellow Green
	"#FFFF00", // Yellow
	"#FFCC00", // Golden Yellow
	"#FF9900", // Orange
	"#FF6600", // Dark Orange
	"#FF3300", // Orange Red
	"#FF0000", // Pure Red
	"#CC0000", // Dark Red
}

func init() {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	// Shuffle the colors slice
	for i := len(colors) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		colors[i], colors[j] = colors[j], colors[i]
	}
}

type Entity interface {
	GetX() int
	GetY() int
	IsObstacle() bool
	GetSymbol() string
}

type Empty struct {
	X          int
	Y          int
	StepsCount int
}

func (e Empty) GetX() int {
	return e.X
}

func (e Empty) GetY() int {
	return e.Y
}

func (e Empty) IsObstacle() bool {
	return false
}

func (e Empty) GetSymbol() string {
	if e.StepsCount == 0 {
		return "."
	}

	colorIndex := e.StepsCount
	if colorIndex >= len(colors) {
		colorIndex = len(colors) - 1
	}
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[colorIndex]))
	return style.Render("X")
}

func (e Empty) GetStepsCount() int {
	return e.StepsCount
}

func (e *Empty) IncreaseSteps() {
	e.StepsCount++
}
