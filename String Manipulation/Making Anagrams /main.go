package main

import (
	"fmt"
	"math"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	counter := map[rune]int{}
	for _, c := range a {
		counter[c] += 1
	}
	for _, c := range b {
		counter[c] -= 1
	}

	res := int32(0)
	for _, v := range counter {
		res += int32(math.Abs(float64(v)))
	}

	fmt.Println(counter)
	return res
}


func main() {
	r := makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke")
	e := int32(30)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}
