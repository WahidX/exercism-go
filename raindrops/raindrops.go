package raindrops

import (
	"fmt"
	"strings"
)

var items = [3]struct{
	factor int
	word string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(input int) (output string) {
	var out strings.Builder

	for _, item := range items {
		if input%item.factor == 0 {
			out.WriteString(item.word)
		}
	}

	if out.Len() == 0 {
		output = fmt.Sprintf("%d", input)
	} else {
		output = out.String()
	}

	return 
}