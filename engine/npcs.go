package engine

type NPC struct {
	Name     string
	Race     string
	Role     string
	NPCLevel int
	Health   int
	Mana     int
	Attack   int
	Armor    int
	Evade    int
	Speed    int
	Luck     int
	Skills   []int
}
