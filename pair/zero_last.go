package pair

func ZeroLast(num []int) []int {
	var zeros []int
	var nonZeros []int

	for i := 0; i < len(num); i++ {
		if num[i] == 0 {
			zeros = append(zeros, num[i])
		} else {
			nonZeros = append(nonZeros, num[i])
		}
	}

	combined := append(nonZeros, zeros...)

	return combined
}
