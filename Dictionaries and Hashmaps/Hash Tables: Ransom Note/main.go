package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	counts := map[string]int{}
	for _, s := range magazine {
		c, ok := counts[s]
		if !ok {
			counts[s] = 1
		}

		c++
		counts[s] = c
	}

	for _, n := range note {
		c, ok := counts[n]
		if !ok || c < 1 {
			fmt.Println("No")
			return
		}
		c--
		counts[n] = c
	}
	fmt.Println("Yes")
}

func main() {
	checkMagazine(
		[]string{"give", "me", "one", "grand", "today", "night"},
		[]string{"give", "one", "grand", "today"},
	)
	e := int64(10)

	println()
	println(e)
}
