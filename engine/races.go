package engine

type Race struct {
	health int
	mana   int
	attack int
	armor  int
	evade  int
	speed  int
	luck   int
}

var RaceMap = map[string]Race{
	"human": {
		health: 15,
		mana:   10,
		attack: 6,
		armor:  0,
		evade:  2,
		speed:  2,
		luck:   2,
	},
	"elf": {
		health: 10,
		mana:   20,
		attack: 8,
		armor:  3,
		evade:  5,
		speed:  5,
		luck:   3,
	},
	"darkelf": {
		health: 15,
		mana:   15,
		attack: 8,
		armor:  4,
		evade:  2,
		speed:  4,
		luck:   1,
	},
	"undead": {
		health: -5,
		mana:   15,
		attack: 10,
		armor:  0,
		evade:  0,
		speed:  0,
		luck:   0,
	},
	"char": {
		health: 20,
		mana:   5,
		attack: 3,
		armor:  5,
		evade:  0,
		speed:  0,
		luck:   10,
	},
	"dwarf": {
		health: 7,
		mana:   7,
		attack: 4,
		armor:  3,
		evade:  2,
		speed:  2,
		luck:   5,
	},
	"orc": {
		health: 17,
		mana:   8,
		attack: 10,
		armor:  5,
		evade:  0,
		speed:  0,
		luck:   0,
	},
}

func GetRaces() []string {
	var slice []string
	for k := range RaceMap {
		slice = append(slice, k)
	}
	return slice
}
