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

// ColumnToSlice takes a column number and returns a slice containing the values in that column
func ColumnToSlice(column int, s [][]byte) []byte {
	newSlice := make([]byte, len(s))
	for i, r := range s {
		newSlice[i] = r[column]
	}
	return newSlice
}
