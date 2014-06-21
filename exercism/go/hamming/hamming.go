package hamming

// min returns the smallest of the two given numbers.
/*func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}*/

// Distance computes the number of differences between two DNA strands.
func Distance(strandA, strandB string) int {
	// Only count differences for the length of the shortest strand.
	/*	length := min(len(strandA), len(strandB))*/
	distance := 0
	for i := 0; i < len(strandA) && i < len(strandB); i++ {
		if strandA[i] != strandB[i] {
			distance++
		}
	}
	return distance
}
