package main

import "strings"

func countLetters(letter, s string) int64 {
	return int64(strings.Count(s, letter))
}


func repeatedString(s string, n int64) int64 {
	d := n / int64(len(s))
	r := n % int64(len(s))

	return countLetters("a", s) * d + countLetters("a", s[:r])
}

func main() {
	r := repeatedString("aba", 10)
	e := int64(7)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}

}
