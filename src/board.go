package main

type Board struct {
    size int
    color [][]rune
	queen [][]bool
}

func newBoard(colors [][]rune) *Board
func (b *Board) isValid() bool
func (b *Board) addQueen(row, col int) bool
func (b *Board) rmvQueen(row, col int) bool