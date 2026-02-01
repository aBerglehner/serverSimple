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

func (b *Board) Won() bool {
	// check rows
	for _, row := range b {
		cur := row[0]
		rowWin := true
		for _, col := range row {
			if cur != col {
				rowWin = false
			}
		}
		if rowWin && cur != Empty {
			return true
		}
	}

	// check vertical cells
	// first vertical col
	if b[0][0] == b[1][0] && b[0][0] == b[2][0] {
		if b[0][0] != Empty {
			return true
		}
	}
	// second vertical col
	if b[0][1] == b[1][1] && b[0][1] == b[2][1] {
		if b[0][1] != Empty {
			return true
		}
	}
	// third vertical col
	if b[0][2] == b[1][2] && b[0][2] == b[2][2] {
		if b[0][2] != Empty {
			return true
		}
	}

	// cross check
	// 0,0; 1,1; 2,2
	if b[0][0] == b[1][1] && b[0][0] == b[2][2] {
		if b[0][0] != Empty {
			return true
		}
	}
	// 0,2; 1,1; 2,0
	if b[0][2] == b[1][1] && b[0][0] == b[2][0] {
		if b[0][2] != Empty {
			return true
		}
	}

	return false
}
