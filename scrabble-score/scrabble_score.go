package scrabble

import "strings"

var charValues = make(map[string]int)

func setChars(keys []string, score int){
	for _, key := range keys{
		charValues[key] = score
	}
}

func init() {
	setChars([]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}, 1)
	setChars([]string{"D", "G"}, 2)
	setChars([]string{"B", "C", "M", "P"}, 3)
	setChars([]string{"F", "H", "V", "W", "Y"}, 4)
	setChars([]string{"K"}, 5)
	setChars([]string{"J", "X"}, 8)
	setChars([]string{"Q", "Z"}, 10)
}

func Score(word string) (score int) {
	for _, c := range word {
		if val, ok := charValues[strings.ToUpper(string(c))]; ok {
			score += val
		}
	}
	return 
}