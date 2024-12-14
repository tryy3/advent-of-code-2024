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

func randomizeColors() {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	// Shuffle the colors slice
	for i := len(colors) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		colors[i], colors[j] = colors[j], colors[i]
	}
}

func init() {
	randomizeColors()
}

type EntityType int

const (
	SpaceEntity EntityType = iota
	GuardEntity
	ObstacleEntity
	VirtualObstacleEntity
)

type StepEvent struct {
	Position  Position
	Direction Direction
}

func NewEntityWithStep(position *Position, entityType EntityType, direction Direction) *Entity {
	return NewEntity(position, entityType, []StepEvent{{Position: *position, Direction: direction}})
}

func NewEntity(position *Position, entityType EntityType, steps []StepEvent) *Entity {
	if steps == nil {
		steps = make([]StepEvent, 0)
	}
	return &Entity{position: position, entityType: entityType, steps: steps}
}

type Entity struct {
	position   *Position
	entityType EntityType
	steps      []StepEvent
}

func (e Entity) GetX() int {
	return e.position.GetX()
}

func (e Entity) GetY() int {
	return e.position.GetY()
}

func (e Entity) GetPosition() *Position {
	return e.position
}

func (e Entity) IsObstacle() bool {
	return e.GetType() == ObstacleEntity || e.GetType() == VirtualObstacleEntity
}

func (e Entity) GetSteps() []StepEvent {
	return e.steps
}

func (e Entity) GetType() EntityType {
	return e.entityType
}

func (e Entity) GetSymbol() string {
	if e.GetType() == ObstacleEntity {
		return "#"
	}
	if e.GetType() == VirtualObstacleEntity {
		return "@"
	}

	if e.GetStepsCount() == 0 {
		return "."
	}

	colorIndex := e.GetStepsCount()
	if colorIndex >= len(colors) {
		colorIndex = len(colors) - 1
	}
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[colorIndex]))
	return style.Render("X")
}

func (e Entity) GetStepsCount() int {
	return len(e.steps)
}

func (e *Entity) TryAddStep(step StepEvent) {
	e.steps = append(e.steps, step)
}

func (e *Entity) ResetSteps() {
	e.steps = make([]StepEvent, 0)
}

func (e *Entity) ChangeEntityType(newType EntityType) {
	e.entityType = newType
}
