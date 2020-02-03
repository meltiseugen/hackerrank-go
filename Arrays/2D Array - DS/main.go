package main

type pair struct {
	I int
	J int
}

var pattern = []pair{
	{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1},
}

func computeHSum(arr [][]int32, i, j int) int32 {
	sum := arr[i][j]
	for _, p := range pattern {
		sum += arr[i + p.I][j + p.J]
	}
	return sum
}

func hourglassSum(arr [][]int32) int32 {
	max := int32(-64)
	for i := 1; i < len(arr) - 1; i++ {
		for j := 1; j < len(arr) - 1; j++ {
			sum := computeHSum(arr, i, j)
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func main() {
	arr := [][]int32{
		{-9, -9, -9,  1, 1, 1},
		{0, -9,  0,  4, 3, 2},
		{-9, -9, -9,  1, 2, 3},
		{0,  0,  8, 6, 6, 0},
		{0,  0,  0, -2, 0, 0},
		{0,  0,  1,  2, 4, 0},
	}
	r := hourglassSum(arr)
	e := int32(28)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}

}
