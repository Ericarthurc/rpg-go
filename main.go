package main

import (
	"fmt"
	"gorpg/engine"
)

func main() {
	PlayerOne := engine.NewPlayer("eric", "char", "warrior")
	fmt.Println(PlayerOne.Skills)

	fmt.Println(PlayerOne.ListSkills())
}
