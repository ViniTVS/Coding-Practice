package main

import "fmt"

var min, max, total int

// pow for integer values
func pow(i, j int) int {
	// recursion base cases
	if j == 0 {
		return 1
	}
	if j == 1 {
		return i
	}

	var floor, ceiling int

	floor = j / 2
	ceiling = j - floor

	// divide and conquer
	return pow(i, floor) * pow(i, ceiling)
}

// convert an array of integers into it's integer representation
func array_to_num(a []int) int {
	n := len(a)
	res := 0
	for i := 0; i < n; i++ {
		res += a[i] * pow(10, n-i-1)
	}

	return res
}

// check if the proposed solution is feasible by checking if it's digits sum
// is what we're looking for
func is_feasible(array []int, sum int) bool {
	sum_check := 0
	// check sum
	for i := 0; i < len(array); i++ {
		sum_check += array[i]
	}

	if sum_check != sum {
		return false
	}
	return true
}

// use backtracking to check and generate number sequences
// every iteration we choose
func FindAllBacktracking(array []int, i, size, sum int) {
	// if we didn't add any number to our array yet, consider inserting the number one
	first_num := 1
	if i > 0 { // or we insert the past inserted value
		first_num = array[i-1]
	}
	// if we inserted the desired number of digits to our array
	if i == size {
		// check if it is feasible and update our values
		if is_feasible(array, sum) {
			total++
			aux := array_to_num(array)
			if aux < min {
				min = aux
			}
			if aux > max {
				max = aux
			}
		}
		// and we stop creating more nodes
		return
	}
	// if we still have digits to insert to our array
	for j := first_num; j < 10; j++ {
		array_copy := array
		array_copy[i] = j
		//  i+1 to keep track how many digits we inserted
		FindAllBacktracking(array_copy, i+1, size, sum)
	}

}

func FindAll(sum int, count int) []int {
	var teste []int = make([]int, count)
	max = 0
	min = 2147483647 // max int val
	total = 0
	FindAllBacktracking(teste, 0, count, sum)

	if total == 0 {
		return nil
	}
	res := []int{total, min, max}

	return res
}

func main() {

	fmt.Println(FindAll(35, 6))

}
