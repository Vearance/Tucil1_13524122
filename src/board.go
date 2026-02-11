package main

type Board struct {
    size int
    color [][]rune
	queen [][]bool
}

func newBoard(colors [][]rune) *Board
func (b *Board) isValid() bool
func (b *Board) isPlaceable(row, col int)
func (b *Board) addQueen(row, col int)
func (b *Board) rmvQueen(row, col int)