package main

func Solve(b *Board, counter *int) bool {
	row := 0;
	rowPlaced := make([]int, b.size)

	for i:=0; i < len(rowPlaced); i++ {
		rowPlaced[i] = -1  // queen does not exist; later save column
	}

	for row >= 0 {  // while syntax in Go
		var placedOnThisRow bool
		startCol := 0
		
		// start after the row that was implemented
		if rowPlaced[row] != -1 { // implemented
			startCol = rowPlaced[row] + 1
			b.rmvQueen(row, rowPlaced[row])
			liveBoard(b, *counter)
		}
		
		for i:=startCol; i < b.size; i++ {
			(*counter)++

			if b.isPlaceable(row, i) {
				b.addQueen(row, i)
				rowPlaced[row] = i  // save Queen's column number on that row 

				liveBoard(b, *counter)

				row = row + 1
				placedOnThisRow = true

				break
			}

		}

		if !placedOnThisRow {
			rowPlaced[row] = -1  // say the queen does not exist again
			row = row - 1
		}

		if row == b.size {
			return true
		}

	}

	return false
}