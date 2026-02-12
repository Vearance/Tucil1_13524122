package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Queens Solver Brute Force Approach")

	// input file name
	var filename string
	fmt.Print("Input nama file (no whitespace, put inside test/ folder): ")
	fmt.Scan(&filename)

	// construct board and read file
	filepath := "../test/" + filename  //TODO: ask if the test folder is used in 
	board, err := constructBoard(filepath) // constructBoard function (param filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// initialized counter var
	var counter int

	// validation
	if !board.isValid() {
		fmt.Println("Board tidak valid.")
		return
	}

	// start var to count duration
	startTime := time.Now()
	

	// do solve
	// bruteforceSolve function (param board from input and counter)
	Solve(board, &counter)

	duration := time.Since(startTime)
	
	fmt.Printf("Waktu pencarian: %.3f ms\n", float64(duration.Microseconds())/1000.0)
	fmt.Println("Banyak kasus yang ditinjau", counter)


	// save function optional
	var save string
	fmt.Println("Apakah Anda ingin menyimpan solusi? (Y/N)")
	fmt.Scan(&save)

	filepath = "../test/solution/" + filename
	if save == "Y" {
		saveBoard(board, filepath)
	}
	

	//TODO: add UI / board changing

}