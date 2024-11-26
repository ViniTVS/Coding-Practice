package main

import (
	"fmt"
	"math"
)

var cherryNum int = 0

func branchBottomRight(grid [][]int, lin, col, cherries, last_index int) {
	// thorn
	if grid[lin][col] == -1 {
		return
	}
	new_cherries := cherries + grid[lin][col]
	new_grid := gridCopy(grid, last_index+1)
	new_grid[lin][col] = 0
	// reached [n-1][n-1]
	if lin == last_index && last_index == col {
		branchTopLeft(new_grid, lin, col, new_cherries, last_index)
		return
	}

	// branching
	if lin < last_index && canPassBottom(grid, lin+1, col, last_index) {
		branchBottomRight(new_grid, lin+1, col, new_cherries, last_index)
	}
	if col < last_index && canPassBottom(grid, lin, col+1, last_index) {
		branchBottomRight(new_grid, lin, col+1, new_cherries, last_index)
	}
}

func gridCopy(grid [][]int, size int) [][]int {
	var ret [][]int = make([][]int, size)

	for i := range grid {
		ret[i] = make([]int, size)
		copy(ret[i], grid[i])
	}

	return ret
}

func branchTopLeft(grid [][]int, lin, col, cherries, last_index int) {
	// thorn
	if grid[lin][col] == -1 {
		return
	}
	new_cherries := cherries + grid[lin][col]

	new_grid := gridCopy(grid, last_index+1)
	new_grid[lin][col] = 0
	// reached [0][0]
	if 0 == lin && 0 == col {
		if new_cherries > cherryNum {
			cherryNum = new_cherries
		}
		return
	}

	// branching
	if lin > 0 && boundTopLeft(new_grid, lin, col, last_index)+new_cherries > cherryNum {
		branchTopLeft(new_grid, lin-1, col, new_cherries, last_index)
	}
	if col > 0 && boundTopLeft(new_grid, lin, col, last_index)+new_cherries > cherryNum {
		branchTopLeft(new_grid, lin, col-1, new_cherries, last_index)
	}
}

func boundTopLeft(grid [][]int, lin, col, last_index int) int {

	total_cherries := 0
	no_exit := true
	last_thorn := -1

	for i := 0; i <= lin; i++ {
		cherry_line := 0
		curr_thorn := -1
		for j := 0; j <= col; j++ {
			if grid[i][j] < 0 {
				cherry_line = 0
				curr_thorn = j
			} else {
				cherry_line += grid[i][j]
			}
		}

		if no_exit && curr_thorn > -1 && i > 0 {
			no_exit = curr_thorn-1 >= last_thorn
		} else {
			no_exit = false
		}
		last_thorn = curr_thorn
		total_cherries += cherry_line
	}

	if no_exit {
		return math.MinInt64
	}

	return total_cherries
}

func findThorn(line []int, col, last_index int) int {
	for i := col; i <= last_index; i++ {
		if line[col] == -1 {
			return i
		}
	}
	return -1
}

func canPassBottom(grid [][]int, lin, col, last_index int) bool {
	// return true
	thorn := -1
	past_thorn := findThorn(grid[lin], col, last_index)

	if lin == last_index {
		return past_thorn < 0
	}

	for i := lin + 1; i < last_index; i++ {
		// no thorn on the way
		if past_thorn == -1 {
			return true
		}
		thorn = findThorn(grid[i], col, last_index)
		// you can skip thorn
		if thorn < past_thorn-1 {
			return true
		}

		past_thorn = thorn
	}

	return findThorn(grid[last_index], col, last_index) >= past_thorn-1
}

func cherryPickup(grid [][]int) int {
	// printMap(grid, len(grid))
	branchBottomRight(grid, 0, 0, 0, len(grid)-1)
	val := cherryNum
	cherryNum = 0
	return val
}

func main() {

	// saida := cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}})
	// fmt.Println("Saida: ", saida)
	// saida = cherryPickup([][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}})
	// fmt.Println("Saida: ", saida)

	// saida := cherryPickup([][]int{{1, 1, 1, 1, -1, -1, -1, 1, 0, 0}, {1, 0, 0, 0, 1, 0, 0, 0, 1, 0}, {0, 0, 1, 1, 1, 1, 0, 1, 1, 1}, {1, 1, 0, 1, 1, 1, 0, -1, 1, 1}, {0, 0, 0, 0, 1, -1, 0, 0, 1, -1}, {1, 0, 1, 1, 1, 0, 0, -1, 1, 0}, {1, 1, 0, 1, 0, 0, 1, 0, 1, -1}, {1, -1, 0, 1, 0, 0, 0, 1, -1, 1}, {1, 0, -1, 0, -1, 0, 0, 1, 0, 0}, {0, 0, -1, 0, 1, 0, 1, 0, 0, 1}})
	// fmt.Println("Saida: ", saida)

	saida := cherryPickup([][]int{{1, -1, 1, 1, -1, 1, 1, 1, -1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, -1, -1, 1, 1, 1, 1, 1, 1}, {1, 1, -1, 1, 1, -1, 1, 1, 1, 1}, {1, 1, 1, -1, 1, 1, 1, -1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {-1, 1, 1, 1, 1, 1, -1, 1, 1, 1}, {1, 1, 0, 1, 1, 1, -1, 1, 1, 1}, {1, -1, 1, -1, -1, 1, 1, 1, 1, 1}, {1, 1, -1, 1, -1, 0, 1, 1, 1, 1}})
	fmt.Println("Saida: ", saida)

}
