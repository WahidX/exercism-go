package etl

import "strings"


func Transform(input map[int][]string) (map[string]int) {
	ans := make(map[string]int)
	for key, strs := range input {
		for _, str := range strs {
			ans[strings.ToLower(str)] = key
		}
	}
	return ans
}