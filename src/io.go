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

	for scanner.Scan() {
		line := scanner.Text() // current line as string
		fmt.Println(line) // TODO: change this to put into board
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return //returns (newBoard(...), nil)
}

func saveBoard(b *Board, filename string) error

func printBoard(b *Board)