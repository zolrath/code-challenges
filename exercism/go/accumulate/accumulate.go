package accumulate

// Accumulate performs given function op on all elements of a slice of strings col.
func Accumulate(col []string, op func(n string) string) []string {
	result := make([]string, len(col))
	for i, v := range col {
		result[i] = op(v)
	}
	return result
}
