package main

import "strings"

func twoStrings(s1 string, s2 string) string {
	workString := s1
	compareString := s2

	if len(s2) < len(s1) {
		workString = s2
		compareString = s1
	}

	for _, c := range workString {
		if strings.Contains(compareString, string(c)) {
			return "YES"
		}
	}
	return "NO"
}


func main() {
	r := twoStrings("a", "today")
	e := "YES"

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}


