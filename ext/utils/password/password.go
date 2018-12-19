package password

import "unicode"

// ValidPassword check if password is valid
func ValidPassword(s string) bool {
	var mustHave = []func(rune) bool{
		unicode.IsUpper,
		unicode.IsLower,
	}
	for _, have := range mustHave {
		found := false

		for _, r := range s {
			if have(r) {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
}
