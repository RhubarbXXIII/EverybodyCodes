package utils

import (
	"fmt"
	"iter"
	"strings"
)

type Direction struct {
	vertical, horizontal int
}

var UP = Direction{-1, 0}
var RIGHT = Direction{0, 1}
var DOWN = Direction{1, 0}
var LEFT = Direction{0, -1}
var DIRECTIONS = [4]Direction{UP, RIGHT, DOWN, LEFT}

func NewDirection(vertical, horizontal int) Direction {
	return Direction{vertical, horizontal}
}

func (direction Direction) Horizontal() int {
	return direction.horizontal
}

func (direction Direction) Vertical() int {
	return direction.vertical
}

func (direction Direction) Negate() Direction {
	return Direction{-direction.vertical, -direction.horizontal}
}

func (direction Direction) RotateRight() Direction {
	return Direction{direction.horizontal, -direction.vertical}
}

func (direction Direction) RotateLeft() Direction {
	return Direction{-direction.horizontal, direction.vertical}
}

func (direction Direction) Multiply(scale int) Direction {
	return Direction{scale * direction.vertical, scale * direction.horizontal}
}

type Position struct {
	row, column int
}

func NewPosition(row, column int) Position {
	return Position{row, column}
}

func (position Position) Row() int {
	return position.row
}

func (position Position) Column() int {
	return position.column
}

func (position Position) Add(direction Direction) Position {
	return Position{position.row + direction.vertical, position.column + direction.horizontal}
}

func (position Position) Subtract(direction Direction) Position {
	return position.Add(direction.Negate())
}

type Grid struct {
	cells [][]rune
}

func NewGrid(values [][]rune) Grid {
	cells := make([][]rune, len(values))
	for i, row := range values {
		cells[i] = make([]rune, len(row))
		copy(cells[i], values[i])

		if i > 0 && len(cells[i]) != len(cells[i-1]) {
			panic(fmt.Sprintf("Row %d has %d cells and row %d has %d cells", i-1, len(cells[i-1]), i, len(cells[i])))
		}
	}

	return Grid{cells}
}

func (grid *Grid) RowCount() int {
	return len(grid.cells)
}

func (grid *Grid) ColumnCount() int {
	if len(grid.cells) == 0 {
		return 0
	}

	return len(grid.cells[0])
}

func (grid *Grid) CellCount() int {
	return grid.RowCount() * grid.ColumnCount()
}

func (grid *Grid) Check(row, column int) bool {
	return row >= 0 && row < grid.RowCount() && column >= 0 && column < grid.ColumnCount()
}

func (grid *Grid) CheckPosition(position Position) bool {
	return grid.Check(position.row, position.column)
}

func (grid *Grid) At(row, column int) rune {
	grid.validateBounds(row, column)
	return grid.cells[row][column]
}

func (grid *Grid) AtPosition(position Position) rune {
	return grid.At(position.row, position.column)
}

func (grid *Grid) Set(row, column int, value rune) {
	grid.validateBounds(row, column)
	grid.cells[row][column] = value
}

func (grid *Grid) SetPosition(position Position, value rune) {
	grid.Set(position.row, position.column, value)
}

func (grid *Grid) Rows() iter.Seq[[]rune] {
	return func(yield func([]rune) bool) {
		for _, row := range grid.cells {
			if !yield(row) {
				return
			}
		}
	}
}

func (grid *Grid) Cells() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, row := range grid.cells {
			for _, cell := range row {
				if !yield(cell) {
					return
				}
			}
		}
	}
}

func (grid *Grid) Positions() iter.Seq[Position] {
	return func(yield func(Position) bool) {
		for r, row := range grid.cells {
			for c := range row {
				if !yield(Position{r, c}) {
					return
				}
			}
		}
	}
}

func (grid *Grid) String() string {
	builder := strings.Builder{}
	for _, row := range grid.cells {
		for _, cell := range row {
			builder.WriteRune(cell)
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func (grid *Grid) validateBounds(row, column int) {
	if row < 0 || row >= grid.RowCount() {
		panic(fmt.Sprintf("Row '%d' is out of bounds (0-%d)", row, grid.RowCount()-1))
	}

	if column < 0 || column >= grid.ColumnCount() {
		panic(fmt.Sprintf("Column '%d' is out of bounds (0-%d)", column, grid.ColumnCount()-1))
	}
}
