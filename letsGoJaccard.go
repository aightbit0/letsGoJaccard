package jaccard

import (
	"strings"
)

//Jaccard ...
func Jaccard(firstString string, secondString string) float64 {
	intersection := make(map[string]bool)
	union := make(map[string]bool)
	firstSplitted := strings.Split(firstString, "")
	secondSplitted := strings.Split(secondString, "")
	for i := 0; i < len(firstSplitted); i++ {
		if union[firstSplitted[i]] != true {
			union[firstSplitted[i]] = true
		}
		for u := 0; u < len(secondSplitted); u++ {
			if firstSplitted[i] == secondSplitted[u] {
				intersection[firstSplitted[i]] = true
			}
			if union[secondSplitted[u]] != true {
				union[secondSplitted[u]] = true
			}
		}
	}
	coefficient := (float64(len(intersection)) / float64(len(union)))

	return coefficient
}
