package main

import (
	"fmt"
	"gorpg/engine"
)

func main() {
	PlayerOne := engine.NewPlayer("eric", "char", "necromancer")
	fmt.Println(PlayerOne.Name)
}
