package engine

type Role struct {
	health int
	mana   int
	attack int
	armor  int
	evade  int
	speed  int
	luck   int
}

var RoleMap = map[string]Role{
	"warrior": {
		health: 100,
		mana:   50,
		attack: 10,
		armor:  15,
		evade:  5,
		speed:  5,
		luck:   5,
	},
	"wizard": {
		health: 80,
		mana:   100,
		attack: 15,
		armor:  5,
		evade:  10,
		speed:  8,
		luck:   10,
	},
	"ranger": {
		health: 70,
		mana:   70,
		attack: 8,
		armor:  5,
		evade:  15,
		speed:  15,
		luck:   13,
	},
	"thief": {
		health: 60,
		mana:   60,
		attack: 12,
		armor:  10,
		evade:  25,
		speed:  20,
		luck:   17,
	},
	"guardian": {
		health: 150,
		mana:   50,
		attack: 5,
		armor:  20,
		evade:  1,
		speed:  1,
		luck:   1,
	},
	"necromancer": {
		health: 60,
		mana:   90,
		attack: 15,
		armor:  10,
		evade:  5,
		speed:  5,
		luck:   -5,
	},
}

func GetRoles() []string {
	var slice []string
	for k := range RoleMap {
		slice = append(slice, k)
	}
	return slice
}
