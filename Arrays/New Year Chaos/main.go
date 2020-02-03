package main

import "fmt"

func minimumBribes(q []int32) {
	minBribes := 0
	for i := len(q)-1; i>=0; i-- {
		if q[i] != int32(i+1) {
			if i-1 >= 0 && q[i-1] == int32(i+1) {
				q[i-1] = q[i]
				q[i] = int32(i)
				minBribes++
			} else if i-2 >= 0 && q[i-2] == int32(i+1) {
				q[i-2] = q[i-1]
				q[i-1] = q[i]
				q[i] = int32(i + 1)
				minBribes += 2
			} else {
				fmt.Println("Too chaotic")
				return
			}
		}
	}
	fmt.Println(minBribes)
}


func main() {
	minimumBribes([]int32{2, 1, 5, 3, 4})
}
