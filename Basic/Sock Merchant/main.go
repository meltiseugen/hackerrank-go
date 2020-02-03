package main

import "sort"

func sockMerchant(n int32, ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
	currentSockType := ar[0]
	counter := 1
	pairs := 0
	for i, v := range ar[1:] {
		if v == currentSockType {
			counter++
		}
		if v != currentSockType || i == len(ar) - 2 {
			if counter % 2 == 0 {
				pairs += counter / 2
			} else if counter != 1 {
				pairs += (counter - 1) / 2
			}
			currentSockType = v
			counter = 1
		}
	}

	return int32(pairs)
}

func main() {
	r := sockMerchant(10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3})
	e := int32(4)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}