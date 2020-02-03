package main

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	a = append(a[d:], a[:d]...)
	return a
}

func main() {
	r := rotLeft([]int32{1, 2, 3, 4, 5}, 2)
	e := int32(3)

	println()
	println(r[0])
	println(e)
	if r[0] == e {
		println("pass")
	} else {
		println("fail")
	}
}
