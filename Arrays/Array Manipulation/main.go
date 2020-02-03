package main

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int32, n)

	for ki := 0; ki < len(queries); ki++ {
		a := queries[ki][0]-1
		b := queries[ki][1]-1
		c := queries[ki][2]

		arr[a] += c
		if b+1 >= int32(len(arr)) {
			continue
		}
		arr[b+1] -= c
	}

	max := int64(arr[0])
	crr := int64(arr[0])
	for i := 1; i < len(arr); i++ {
		crr += int64(arr[i])

		if crr > max {
			max = crr
		}
	}

	return int64(max)
}

func main() {
	r := arrayManipulation(10, [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 10, 1},
	})
	e := int64(10)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}

}
