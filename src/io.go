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
		fmt.Println(line) // TODO: change this to put into board
		row := []rune(line) // take this row line to a char array and put it into matrix
		colorMatrix = append(colorMatrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return newBoard(colorMatrix), nil
}

func saveBoard(b *Board, filename string) error

func printBoard(b *Board)