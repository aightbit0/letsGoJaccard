package jaccard

import (
	"strings"
)

func jaccard(searchval string, stest string) float64 {
	intersection := make(map[string]bool)
	union := make(map[string]bool)
	val1 := strings.Split(searchval, "")
	val2 := strings.Split(stest, "")
	for i := 0; i < len(val1); i++ {
		if union[val1[i]] != true {
			union[val1[i]] = true
		}
		for u := 0; u < len(val2); u++ {
			if val1[i] == val2[u] {
				intersection[val1[i]] = true
			}
			if union[val2[u]] != true {
				union[val2[u]] = true
			}
		}
	}
	coefficient := (float64(len(intersection)) / float64(len(union)))

	return coefficient
}
