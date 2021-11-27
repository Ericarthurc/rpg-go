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

func NewNPC(name string) *NPC {
	role := GetRandomRole()
	race := GetRandomRace()
	npc := new(NPC)
	npc.Name = name
	npc.Race = race
	npc.Role = role
	npc.Health = RaceMap[race].health + RoleMap[role].health
	npc.Mana = RaceMap[race].mana + RoleMap[role].mana
	npc.Attack = RaceMap[race].attack + RoleMap[role].attack
	npc.Armor = RaceMap[race].armor + RoleMap[role].armor
	npc.Evade = RaceMap[race].evade + RoleMap[role].evade
	npc.Speed = RaceMap[race].speed + RoleMap[role].speed
	npc.Luck = RaceMap[race].luck + RoleMap[role].luck

	for _, v := range SkillMap[role] {
		if v.levelRequirement == 0 {
			npc.Skills = append(npc.Skills, v.id)
		}
	}

	return npc
}
