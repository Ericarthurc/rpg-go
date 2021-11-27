package engine

import "math/rand"

func RNGMinMax(min, max int) int {
	return rand.Intn((max+1)-min) + min
}

func GetRandomRole() string {
	roles := GetRoles()
	return roles[RNGMinMax(0, len(roles)-1)]
}

func GetRandomRace() string {
	races := GetRaces()
	return races[RNGMinMax(0, len(races)-1)]
}
