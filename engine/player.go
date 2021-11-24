package engine

type Player struct {
	Name             string
	Race             string
	Role             string
	PlayerExperience int
	PlayerLevel      int
	Health           int
	Mana             int
	Attack           int
	Armor            int
	Evade            int
	Speed            int
	Luck             int
	Skills           []int
}

func NewPlayer(name, race, role string) *Player {
	p := new(Player)
	p.Name = name
	p.Race = race
	p.Role = role
	p.PlayerExperience = 0
	p.PlayerLevel = 0
	p.Health = RaceMap[race].health + RoleMap[role].health
	p.Mana = RaceMap[race].mana + RoleMap[role].mana
	p.Attack = RaceMap[race].attack + RoleMap[role].attack
	p.Armor = RaceMap[race].armor + RoleMap[role].armor
	p.Evade = RaceMap[race].evade + RoleMap[role].evade
	p.Speed = RaceMap[race].speed + RoleMap[role].speed
	p.Luck = RaceMap[race].luck + RoleMap[role].luck

	for _, v := range SkillMap[role] {
		if v.levelRequirement == 0 {
			p.Skills = append(p.Skills, v.id)
		}
	}

	return p
}

func (p Player) ListSkills() []Skill {
	var equipped []Skill
	for _, id := range p.Skills {
		for _, v := range SkillMap[p.Role] {
			if v.id == id {
				equipped = append(equipped, v)
			}
		}
	}
	return equipped
}

func Attack(bot Player) {

}
