package engine

import "math/rand"

func RNGMinMax(min, max int) int {
	return rand.Intn((max+1)-min) + min
}
