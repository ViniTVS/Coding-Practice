package main

import (
	"fmt"
	"sort"
	"unicode"
)

type Entries struct {
	count_a int
	count_b int
	val     string
}

func printItens(array []*Entries) {
	for _, j := range array {
		fmt.Println(j.val)
		fmt.Println("\t", j.count_a)
		fmt.Println("\t", j.count_b)
	}
}

func findEntry(str string, array []*Entries) int {
	for i, j := range array {
		if j.val == str {
			return i
		}
	}

	return -1
}

func Mix(s1, s2 string) string {

	var list []*Entries

	res := ""

	for _, chr := range s1 {
		if !unicode.IsLower(chr) || !unicode.IsLetter(chr) {
			continue
		}
		index := findEntry(string(chr), list)
		if index >= 0 {
			list[index].count_a++
		} else {
			new_insert := new(Entries)
			new_insert.val = string(chr)
			new_insert.count_a = 1
			list = append(list, new_insert)
		}
	}

	for _, chr := range s2 {
		if !unicode.IsLower(chr) || !unicode.IsLetter(chr) {
			continue
		}
		index := findEntry(string(chr), list)
		if index >= 0 {
			list[index].count_b++
		} else {
			new_insert := new(Entries)
			new_insert.val = string(chr)
			new_insert.count_b = 1
			list = append(list, new_insert)
		}
	}

	sort.SliceStable(list, func(i, j int) bool {
		biggest_i := list[i].count_a
		biggest_j := list[j].count_a

		if biggest_i < list[i].count_b {
			biggest_i = list[i].count_b
		}
		if biggest_j < list[j].count_b {
			biggest_j = list[j].count_b
		}

		return biggest_i > biggest_j
	})

	for index, item := range list {
		if index > 0 {
			res += "/"
		}
		if item.count_a > item.count_b {
			res += "1:"
		} else if item.count_b > item.count_a {
			res += "2:"
		} else {
			res += "=:"
		}

		i := 0

		for ; i < item.count_b; i++ {
			res += item.val
		}

		for ; i < item.count_a; i++ {
			res += item.val
		}

	}

	return res
}
func main() {
	fmt.Println(Mix("Are they here", "yes, they are here"))
}
