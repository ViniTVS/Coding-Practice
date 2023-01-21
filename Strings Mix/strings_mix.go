package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Entries struct {
	count1 int  // s1 char freq
	count2 int  // s2 char freq
	val    rune // char
}

// Returns the chr index if exists or -1 othewise
func findEntry(chr rune, array []*Entries) int {
	for i, j := range array {
		if j.val == chr {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Mix(s1, s2 string) string {

	var list []*Entries

	// count number of char appearances in string 1
	for _, chr := range s1 {
		// can only be lowercase letters
		if !unicode.IsLower(chr) || !unicode.IsLetter(chr) {
			continue
		}

		index := findEntry(chr, list)
		if index >= 0 {
			list[index].count1++
		} else { // create new entry
			new_insert := new(Entries)
			new_insert.val = chr
			new_insert.count1 = 1
			list = append(list, new_insert)
		}
	}
	// do the same for string 2
	for _, chr := range s2 {
		if !unicode.IsLower(chr) || !unicode.IsLetter(chr) {
			continue
		}
		index := findEntry(chr, list)
		if index >= 0 {
			list[index].count2++
		} else {
			new_insert := new(Entries)
			new_insert.val = chr
			new_insert.count2 = 1
			list = append(list, new_insert)
		}
	}

	var all_str []string
	// create the strings that will be the answer
	for _, v := range list {
		// maximum is strictly greater than 1
		if v.count1 < 2 && v.count2 < 2 {
			continue
		}

		aux_str := ""
		// check if new string starts with 1, 2 or =
		if v.count1 > v.count2 {
			aux_str += "1:"
		} else if v.count2 > v.count1 {
			aux_str += "2:"
		} else {
			aux_str += "=:"
		}
		// add the number of necessary chars
		for i := 0; i < max(v.count1, v.count2); i++ {
			aux_str += string(v.val)
		}
		all_str = append(all_str, aux_str)
	}
	// and sort'em
	sort.SliceStable(all_str, func(i, j int) bool {
		str_len_i := len(all_str[i])
		str_len_j := len(all_str[j])
		// if they have the same length, sort in ascending lexicographic order
		if str_len_j == str_len_i {
			return strings.Compare(all_str[i], all_str[j]) < 0
		}
		// otherwise, by length
		return str_len_i > str_len_j
	})

	return strings.Join(all_str, "/")
}

func main() {
	fmt.Println(
		Mix("my&friend&Paul has heavy hats! &", "my friend John has many many friends &")
	)
}
