package main

import (
	"fmt"
)

func Spiralize(size int) [][]int {
	// Make a snake
	spiral := make([][]int, size)

	// fill spiral with 1s
	for i := 0; i < size; i++ {
		spiral[i] = make([]int, size)
		for j := 0; j < size; j++ {
			spiral[i][j] = 1
		}
	}

	for margin := 1; margin <= size/2; margin += 2 {

		spiral[margin][margin-1] = 0

		for i := margin; i < size-margin; i++ {
			spiral[margin][i] = 0        // top row
			spiral[size-1-margin][i] = 0 // bot row
		}
		// right column
		for i := margin + 1; i < size-margin-1; i++ {
			spiral[i][size-1-margin] = 0
		}
		// left column
		for i := margin + 2; i < size-margin-1; i++ {
			spiral[i][margin] = 0
		}
	}
	return spiral
}

func printSpiral(saida [][]int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(saida[i][j], "")
		}
		fmt.Println("")
	}

}

func main() {
	size := 11
	saida := Spiralize(size)

	fmt.Println("Saida: ")
	printSpiral(saida, size)

}
