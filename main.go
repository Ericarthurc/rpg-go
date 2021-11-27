package main

import (
	"fmt"
	"gorpg/engine"
	"math/rand"
	"time"
)

func main() {
	// seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	// var x int
	// for x <= 100 && x >= 0 {
	// 	x = engine.RNGMinMax(-1, 101)
	// 	fmt.Println(x)
	// }

	// PlayerOne := engine.NewPlayer("eric", "char", "warrior")
	// fmt.Println(PlayerOne.Skills)

	// fmt.Println(PlayerOne.ListSkills())

	NpcOne := engine.NewNPC("Botty")
	fmt.Println(NpcOne)
}
