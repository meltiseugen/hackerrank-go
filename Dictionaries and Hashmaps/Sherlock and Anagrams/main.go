package main

import (
	"strconv"
)

func sliceToString(m []int) string {
	s := ""
	for _, v := range m {
		s += strconv.Itoa(v)
	}
	return s
}

func sherlockAndAnagrams(s string) int32 {
	signatures := map[string]int{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			signature := make([]int, 26)
			for _, letter := range s[i:j+1] {
				signature[letter - 'a'] += 1
			}
			signatures[sliceToString(signature)] += 1
		}
	}

	result := 0
	for _, v := range signatures {
		if v <= 1 {
			continue
		}
		result += v*(v-1)/2
	}

	return int32(result)
}


func main() {
	r := sherlockAndAnagrams("ifailuhkqq")
	e := int32(3)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}
