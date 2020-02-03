package main


func balanced(gN int, d map[rune]int) bool {
	if d['A'] > gN || d['C'] > gN || d['G'] > gN || d['T'] > gN {
		return false
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Complete the steadyGene function below.
func steadyGene(gene string) int32 {
	N := len(gene)
	gN := N/4
	dict := map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	minLength := N
	left, right := 0, 0

	for _, c := range gene {
		dict[c]++
	}

	if dict['A'] == gN && dict['C'] == gN && dict['G'] == gN && dict['T'] == gN {
		return 0
	}

	for left < N && right < N {
		if !balanced(gN, dict) {
			dict[[]rune(gene)[right]] -= 1
			right += 1
		} else {
			minLength = min(minLength, right - left)
			dict[[]rune(gene)[left]] += 1
			left += 1
		}
	}

	return int32(minLength)
}

func main() {
	r := steadyGene("GAAATAAA")
	e := int32(5)

	println()
	println(r)
	println(e)
	if r == e {
		println("pass")
	} else {
		println("fail")
	}

}
