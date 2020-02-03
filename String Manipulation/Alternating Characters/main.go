package main

// Complete the alternatingCharacters function below.
func alternatingCharacters(s string) int32 {
	count := int32(0)
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == s[i+1] {
			count++
		}
	}

	return count
}

func main() {
	r := alternatingCharacters("AAABBB")
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
