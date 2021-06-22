package acronym

import (
	"unicode"
)

func Abbreviate(s string) (acronym string) {
	
	for idx, char := range s {
		if char == '_' || char == '-' || char == ' ' {
			continue
		} else if idx == 0 {
			acronym += string(char)
		} else {
			prev := s[idx-1]
			isPrevLetter := unicode.IsLetter(rune(prev)) || prev == 39
			if !isPrevLetter {
				acronym += string(unicode.ToUpper(char))
			}
		}
	}

	return
}
