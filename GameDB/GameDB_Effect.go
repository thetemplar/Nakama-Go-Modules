package main

type GameDB_Effect struct {
	Id				int64
	Name 			string
	Description 	string
	Visible			bool
	EffectID		int64
	Duration 		float32
	Dispellable		bool
	School 			GameDB_Spell_SchoolType
	Mechanic		GameDB_Spell_Mechanic
	Type 			interface{}
	ValueMin 		int32
	ValueMax 		int32
}

type GameDB_Effect_Autoattack struct {
}

type GameDB_Effect_Damage struct {
}

type GameDB_Effect_Heal struct {
}

type GameDB_Effect_Apply_Aura_Periodic_Damage struct {
	Intervall   	float32
}

type GameDB_Effect_Apply_Aura_Periodic_Heal struct {
	Intervall   	float32
}

type GameDB_Effect_Apply_Aura_Mod struct {
	Stat			GameDB_Stat
	Value 			float32
}

type GameDB_Effect_Persistent_Area_Aura struct {
	Intervall   	float32
	Radius		   	float32
}

type GameDB_Stat int
const (
	GameDB_Stat_Speed = 0
)

type GameDB_Spell_SchoolType int
const (
	GameDB_Spell_SchoolType_Physical = 0
	GameDB_Spell_SchoolType_Arcane = 1
	GameDB_Spell_SchoolType_Fire = 2
	GameDB_Spell_SchoolType_Frost = 3
	GameDB_Spell_SchoolType_Nature = 4
	GameDB_Spell_SchoolType_Shadow = 5
	GameDB_Spell_SchoolType_Holy = 6
)

type GameDB_Spell_Mechanic int
const (
	GameDB_Spell_Mechanic_None = 0
	GameDB_Spell_Mechanic_Rooted = 1
	GameDB_Spell_Mechanic_Sapped = 2
	GameDB_Spell_Mechanic_Invulnerable = 3
	GameDB_Spell_Mechanic_Interrupted = 4
	GameDB_Spell_Mechanic_Infected = 5
	GameDB_Spell_Mechanic_Shielded = 6
	GameDB_Spell_Mechanic_Slowed = 7
	GameDB_Spell_Mechanic_Stunned = 8
	GameDB_Spell_Mechanic_Healing = 9
)