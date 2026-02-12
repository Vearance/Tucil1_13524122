package main

import "fmt"

func Solve(b *Board, counter *int) bool {
	row := 0;
	rowPlaced := make([]int, b.size)

	for i:=0; i < len(rowPlaced); i++ {
		rowPlaced[i] = -1  // queen does not exist; later save column
	}

	for row >= 0 {  // while syntax in Go
		var placedOnThisRow bool
		
		for i:=0; i < b.size; i++ {
			(*counter)++
			fmt.Println(*counter)

			if b.isPlaceable(row, i) {
				b.addQueen(row, i)
				rowPlaced[row] = i  // save Queen's column number on that row 

				row = row + 1
				placedOnThisRow = true

				break
			}

		}

		if !placedOnThisRow {
			if row > 0 {
				b.rmvQueen(row, rowPlaced[row-1])
				rowPlaced[row] = -1  // say the queen does not exist again
			}

			row = row - 1
		}

		if row == b.size {
			return true
		}

	}

	return false
}