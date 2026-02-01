package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Cell byte

const (
	Empty Cell = ' '
	X     Cell = 'X'
	O     Cell = 'O'
)

type Board [3][3]Cell

func NewBoard() *Board {
	return &Board{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
}

func (b *Board) Update(cellType Cell, move string) {
	// A=0,B=1,C=3
	c := strings.ToUpper(move[0:1])
	var col int
	switch c {
	case "A":
		col = 0
	case "B":
		col = 1
	case "C":
		col = 2
	}

	row, _ := strconv.Atoi(strings.TrimSpace(move[1:]))
	b[row-1][col] = cellType
}

// TODO: add win func
func (b Board) String() string {
	var sb strings.Builder
	sb.WriteString("  A   B   C\n")
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%d %c | %c | %c\n",
			i+1, b[i][0], b[i][1], b[i][2]))
		if i < 2 {
			sb.WriteString(" ---+---+---\n")
		}
	}
	return sb.String()
}
