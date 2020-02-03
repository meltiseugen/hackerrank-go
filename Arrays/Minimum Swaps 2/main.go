package main

import (
	"fmt"
	"sort"
)

type pair struct {
	V int32
	P int
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	visited := make([]bool, len(arr))
	var pairs []pair
	finalCounter := 0

	for i, v := range arr {
		pairs = append(pairs, pair{V: v, P: i})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].V < pairs[j].V
	})

	for i := 0; i < len(arr); i++ {
		j := i
		counter := 0

		if visited[i] || pairs[i].P == i {
			continue
		}

		for !visited[j] {
			visited[j] = true
			j = pairs[j].P
			counter++
		}
		fmt.Println(counter)
		finalCounter += counter-1
	}

	return int32(finalCounter)
}

func main() {
	r := minimumSwaps([]int32{2, 3, 4, 1, 5})
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
