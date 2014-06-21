package allergies

// Global makes this cleaner but... global.
// Bitshifting to make each entry a new power of 2.
var allergens = map[string]int{
	"eggs":         1 << 0,
	"peanuts":      1 << 1,
	"shellfish":    1 << 2,
	"strawberries": 1 << 3,
	"tomatoes":     1 << 4,
	"chocolate":    1 << 5,
	"pollen":       1 << 6,
	"cats":         1 << 7,
}

// Allergies returns a list of allergies one would have with a particular Allergen Score.
func Allergies(n int) []string {
	allergies := []string{}
	for allergen, _ := range allergens {
		if AllergicTo(n, allergen) {
			allergies = append(allergies, allergen)
		}
	}
	return allergies
}

// AllericTo checks to see if the user is allergic to a particular allergen based off their score.
func AllergicTo(n int, allergen string) bool {
	// "bitwise and" operation checks to see if the same bit is set to 1 in both numbers.
	return n&allergens[allergen] != 0
}
