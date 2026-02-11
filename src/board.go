package main

import "unique"

type Board struct {
    size int
    color [][]rune
	queen [][]bool
}

func newBoard(colorMatrix [][]rune) *Board {
	size := len(colorMatrix)

	queenMatrix := make([][]bool, size)
	for i := 0; i < size; i++ {
		queenMatrix[i] = make([]bool, size)
	}

	return &Board {
		size: size,
		color: colorMatrix,
		queen: queenMatrix,
	}
}

func (b *Board) isValid() bool {  // every color must appear once, number of color == board size/length
	// check unique colors
	unique := make(map[rune]bool)
	for i:=0; i < b.size; i++ {
		for j:=0; j < b.size; j++ {
			unique[b.color[i][j]] = true
		}
	}
	count := len(unique)

	if count == b.size {
		return true
	} else {
		return false
	}

}

func (b *Board) isPlaceable(row, col int) bool {
	if b.queen[row][col] == true {
		return false
	}

	return true
}

func (b *Board) addQueen(row, col int) {
	b.queen[row][col] = true
}

func (b *Board) rmvQueen(row, col int) {
	b.queen[row][col] = false
}