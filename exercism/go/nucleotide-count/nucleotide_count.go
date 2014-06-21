package dna

import "fmt"

// DNA contains a string representing the DNA sequence.
type DNA struct {
	sequence string
}

// Histogram holds the nucleotide counts in a DNA sequence.
type Histogram map[string]int

// Count returns the number of times a given nucleotide appears in a DNA sequence.
func (d DNA) Count(nuc string) (int, error) {
	nucleotides := map[string]bool{"A": true, "C": true, "G": true, "T": true}
	if !nucleotides[nuc] {
		return 0, fmt.Errorf("%s is an invalid nucleotide.", nuc)
	}

	sum := 0
	for _, v := range d.sequence {
		if nuc == string(v) {
			sum++
		}
	}

	return sum, nil
}

// Counts returns a Histogram object containing the counts of each nucleotide in a DNA sequnce.
func (d DNA) Counts() *Histogram {
	aCount, err := d.Count("A")
	cCount, err := d.Count("C")
	gCount, err := d.Count("G")
	tCount, err := d.Count("T")

	if err != nil {
		// Not really able to handle errors properly with required return values.
		fmt.Println(err)
	}

	return &Histogram{"A": aCount, "C": cCount, "G": gCount, "T": tCount}
}
