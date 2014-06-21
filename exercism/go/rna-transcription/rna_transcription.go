package strand

func ToRna(dna string) string {
	// Using a slice of runes is more efficient than creating many immutable strings via concatination.
	rna := make([]rune, len(dna))
	complements := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for i, n := range dna {
		rna[i] = complements[n]
	}

	return string(rna)
}
