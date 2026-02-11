package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Queens Solver Brute Force Approach")

	// input file name
	var filename string
	fmt.Print("Input nama file (no whitespace): ")
	fmt.Scan(&filename)

	// construct board
	board, err := // constructBoard function (param filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// initialized counter var
	var counter int

	// start var to count duration
	startTime := time.Now()

	// do solve
	// bruteforceSolve function (param board from input and counrer)

	duration := time.Since(startTime).Milliseconds()

}