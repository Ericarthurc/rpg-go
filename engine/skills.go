package engine

type Skill struct {
	id               int
	skillName        string
	description      string
	damage           int
	levelRequirement int
	manaRequirement  int
}

var SkillMap = map[string][]Skill{
	"warrior": {
		{
			id:               1,
			skillName:        "Head Butt",
			description:      "Swing your head forward with mighty force!",
			damage:           5,
			levelRequirement: 0,
			manaRequirement:  5,
		},
		{
			id:               2,
			skillName:        "Slash",
			description:      "Slash your weapon forward striking whats in front of you!",
			damage:           12,
			levelRequirement: 0,
			manaRequirement:  8,
		},
	},
	"wizard": {
		{
			id:               1,
			skillName:        "Lightning Strike",
			description:      "Lightning strikes from above!",
			damage:           12,
			levelRequirement: 0,
			manaRequirement:  10,
		},
	},
	"ranger": {
		{
			id:               1,
			skillName:        "Ranger1",
			description:      "",
			damage:           0,
			levelRequirement: 0,
			manaRequirement:  0,
		},
	},
	"thief": {
		{
			id:               1,
			skillName:        "Thief1",
			description:      "",
			damage:           0,
			levelRequirement: 0,
			manaRequirement:  0,
		},
	},
	"guardian": {
		{
			id:               1,
			skillName:        "Guardian1",
			description:      "",
			damage:           0,
			levelRequirement: 0,
			manaRequirement:  0,
		},
	},
	"necromancer": {
		{
			id:               1,
			skillName:        "Necromancer1",
			description:      "",
			damage:           0,
			levelRequirement: 0,
			manaRequirement:  0,
		},
	},
}
