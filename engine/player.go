package engine

type Player struct {
	Name             string
	Race             string
	Role             string
	playerExperience int
	playerLevel      int
	health           int
	mana             int
	attack           int
	armor            int
	evade            int
	speed            int
	luck             int
	skills           []int
}

func NewPlayer(name, race, role string) *Player {
	p := new(Player)
	p.Name = name
	p.Race = race
	p.Role = role
	p.playerExperience = 0
	p.playerLevel = 0
	p.health = RaceMap[race].health + RoleMap[role].health
	p.mana = RaceMap[race].mana + RoleMap[role].mana
	p.attack = RaceMap[race].attack + RoleMap[role].attack
	p.armor = RaceMap[race].armor + RoleMap[role].armor
	p.evade = RaceMap[race].evade + RoleMap[role].evade
	p.speed = RaceMap[race].speed + RoleMap[role].speed
	p.luck = RaceMap[race].luck + RoleMap[role].luck
	return p
}
