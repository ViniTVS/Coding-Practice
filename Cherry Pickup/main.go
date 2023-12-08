package main

import (
	"fmt"
)

var cherries int = -1
var cherryMap [][]int = nil

func printMap(saida [][]int, size int) {
	fmt.Println("MAPA:")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(saida[i][j], ", ")
		}
		fmt.Println("")
	}

}

func branchBottomRight(grid [][]int, lin, col, cherries, last_index int) {
	// fmt.Println(lin, col, last_index)
	// thorn
	if grid[lin][col] == -1 {
		return
	}
	new_grid := grid
	new_cherries := cherries + grid[lin][col]
	new_grid[lin][col] = 0
	// reached [n-1][n-1]
	if lin == col && lin == last_index {
		printMap(grid, last_index+1)
		return
	}

	// branching
	if lin < last_index {
		branchBottomRight(new_grid, lin+1, col, new_cherries, last_index)
	}
	if col < last_index {
		branchBottomRight(new_grid, lin, col+1, new_cherries, last_index)
	}

}

func cherryPickup(grid [][]int) int {
	// a := 0

	new_grid := grid
	branchBottomRight(new_grid, 0, 0, 0, len(grid)-1)

	return cherries
}

func main() {

	saida := cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}})
	fmt.Println("Saida: ", saida)
	saida = cherryPickup([][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}})
	fmt.Println("Saida: ", saida)
}
