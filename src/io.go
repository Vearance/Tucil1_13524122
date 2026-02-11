package main

import (
	"fmt"
	"bufio"
	"os"
)

func constructBoard(filename string) (*Board, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//initialize colorMatrix
	var colorMatrix [][]rune

	for scanner.Scan() {
		line := scanner.Text() // current line as string
		row := []rune(line) // take this row line to a char array and put it into matrix
		colorMatrix = append(colorMatrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return newBoard(colorMatrix), nil
}

func saveBoard(b *Board, filename string) error {
	file, err := os.Create(filename) // if it exists, it will be replaced
	if err != nil {
		return err
	}
	defer file.Close()

	var char int

	for i := 0; i < b.size; i++ {
		for j:= 0; j < b.size; j++ {
			if b.queen[i][j] {
				char, err = file.WriteString("#")
			} else {
				char, err = file.WriteString(string(b.color[i][j]))
			}
		}

		char, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func printBoard(b *Board) {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.queen[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(string(b.color[i][j]))
			}
		}
		fmt.Println()
	}
}