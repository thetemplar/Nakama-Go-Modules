package GameDB

type Effect struct {
	Id				int64
	Name 			string
	Description 	string
	Visible			bool
	EffectID		int64
	Duration 		float32
	Dispellable		bool
	School 			Spell_SchoolType
	Mechanic		Spell_Mechanic
	Type 			interface{}
	ValueMin 		int32
	ValueMax 		int32
}

type Effect_Autoattack struct {
}

type Effect_Damage struct {
}

type Effect_Heal struct {
}

type Effect_Apply_Aura_Periodic_Damage struct {
	Intervall   	float32
}

type Effect_Apply_Aura_Periodic_Heal struct {
	Intervall   	float32
}

type Effect_Apply_Aura_Mod struct {
	Stat			Stat
	Value 			float32
}

type Effect_Persistent_Area_Aura struct {
	Intervall   	float32
	Radius		   	float32
}

type Stat int
const (
	Stat_Speed = 0
	Stat_Stamina = 1
	Stat_Strength = 2
	Stat_Agility = 3
	Stat_Intellect = 4
	Stat_PhysicalAP = 5
	Stat_SpellAP = 6
	Stat_Armor = 7
)

type Spell_SchoolType int
const (
	Spell_SchoolType_Physical = 0
	Spell_SchoolType_Arcane = 1
	Spell_SchoolType_Fire = 2
	Spell_SchoolType_Frost = 3
	Spell_SchoolType_Nature = 4
	Spell_SchoolType_Shadow = 5
	Spell_SchoolType_Holy = 6
)

type Spell_Mechanic int
const (
	Spell_Mechanic_None = 0
	Spell_Mechanic_Rooted = 1
	Spell_Mechanic_Sapped = 2
	Spell_Mechanic_Invulnerable = 3
	Spell_Mechanic_Interrupted = 4
	Spell_Mechanic_Infected = 5
	Spell_Mechanic_Shielded = 6
	Spell_Mechanic_Slowed = 7
	Spell_Mechanic_Stunned = 8
	Spell_Mechanic_Healing = 9
)