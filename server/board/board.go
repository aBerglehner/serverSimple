package board

import (
	"fmt"
	"strings"
)

type Cell byte

const (
	Empty Cell = ' '
	X     Cell = 'X'
	O     Cell = 'O'
)

type Board [3][3]Cell

func NewBoard() Board {
	return Board{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
}

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
