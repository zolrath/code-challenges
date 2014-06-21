package allergies

// This avoids using globals aside from the const declaration at the top but results in much more duplication.
const (
	// iota starts at 0 and increases by 1 each line, shifting the bit to the next power of 2.
	eggs         = 1 << iota
	peanuts      = 1 << iota
	shellfish    = 1 << iota
	strawberries = 1 << iota
	tomatoes     = 1 << iota
	chocolate    = 1 << iota
	pollen       = 1 << iota
	cats         = 1 << iota
)

// Allergies returns a list of allergies one would have with a particular Allergen Score.
func Allergies2(n int) []string {
	allergies := []string{}
	possibleAllergies := []string{"eggs", "peanuts", "shellfish",
		"strawberries", "tomatoes", "chocolate", "pollen", "cats"}

	for _, allergen := range possibleAllergies {
		if AllergicTo2(n, allergen) {
			allergies = append(allergies, allergen)
		}
	}
	return allergies
}

// AllericTo checks to see if the user is allergic to a particular allergen based off their score.
func AllergicTo2(n int, allergen string) bool {
	allergens := map[string]int{"eggs": eggs, "peanuts": peanuts,
		"shellfish": shellfish, "strawberries": strawberries, "tomatoes": tomatoes,
		"chocolate": chocolate, "pollen": pollen, "cats": cats}
	// "bitwise and" operation checks to see if the same bit is set to 1 in both numbers.
	return n&allergens[allergen] != 0
}
