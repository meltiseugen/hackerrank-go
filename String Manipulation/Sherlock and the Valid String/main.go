package main

// Complete the isValid function below.
func isValid(s string) string {
	counter := map[rune]int{}
	min := 100000
	max := 0

	for _, c := range s {
		counter[c] += 1
		if counter[c] > max {
			max = counter[c]
		}
		if counter[c] < min {
			min = counter[c]
		}
	}

	moreThan2G := 0
	for _, v := range counter{
		if v > 2 {
			return "NO"
		}

		if v >= 2 {
			moreThan2G++
		}

		if moreThan2G >= 2 {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	r := isValid("abcdefghhgfedecba")
	e := "NO"

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}

}
