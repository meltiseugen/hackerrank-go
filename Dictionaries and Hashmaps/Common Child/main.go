package main

func max(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func commonChild(s1 string, s2 string) int32 {
	matcher := make([][]int32, len(s1))
	for i := 0; i < len(s1); i++ {
		matcher[i] = append(make([]int32, len(s2)))
	}

	for i, c1 := range s1 {
		for j, c2 := range s2 {
			xm := 0
			ym := 0
			if i - 1 > 0 {
				xm = i - 1
			}
			if j - 1 > 0 {
				ym = j - 1
			}
			if c1 == c2 {

				matcher[i][j] = matcher[xm][ym] + 1
			} else {
				matcher[i][j] = max(matcher[i][ym], matcher[xm][j])
			}
		}
	}
	return matcher[len(s1) - 1][len(s2) - 1]
}

func main() {
	r := commonChild("HARRY", "SALLY")
	e := int32(2)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}
