package main

import (
	"fmt"
)

func numLastDigit(n int) int {
	return (n % 10)
}

func expLastDigit(i int, j int) int {
	if j == 0 {
		return 1
	}
	if j == 1 {
		return numLastDigit(i)
	}

	var floor, ceiling int

	floor = j / 2
	ceiling = j - floor

	return numLastDigit(numLastDigit(expLastDigit(i, floor)) * numLastDigit(expLastDigit(i, ceiling)))
}

func lastDigit(array []int) int {
	fmt.Println(array)
	n := len(array)
	var aux int

	if n == 0 {
		return 1
	}
	if n == 1 {
		return array[0]
	}

	aux = expLastDigit(array[n-2], array[n-1])
	array = array[:n-2]
	array = append(array, aux)

	return lastDigit(array)
}

func main() {
	list := []int{12, 30, 21}

	fmt.Println(lastDigit(list))

}
