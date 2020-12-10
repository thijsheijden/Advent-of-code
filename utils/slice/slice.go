package slice

// Max gets the largest value from a slice of ints
func Max(input []int) int {
	max := 0
	for _, element := range input {
		if element > max {
			max = element
		}
	}

	return max
}
