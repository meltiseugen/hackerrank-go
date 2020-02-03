package main

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	passes := 0
	currentCloud := 0
	for currentCloud != len(c) - 1 {
		if currentCloud + 2 <= len(c) - 1 && c[currentCloud + 2] != 1 {
			currentCloud += 2
		} else {
			currentCloud++
		}
		passes++
	}

	return int32(passes)
}

func main() {
	r := jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0})
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
