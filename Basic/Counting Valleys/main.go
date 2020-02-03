package main

var (
	mp = map[int32]int{
		'U': 1,
		'D': -1,
	}
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	position := 0
	valleys := 0
	inValley := false

	for _, action := range s {
		position += mp[action]
		if position < 0 && inValley == false {
			inValley = true
		}
		if position >= 0 && inValley == true {
			inValley = false
			valleys++
		}
	}

	return int32(valleys)
}

func main() {
	r := countingValleys(8, "UDDDUDUU")
	e := int32(1)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}
}
