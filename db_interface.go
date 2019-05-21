package main

type GameDB struct {
	Spells				map[int64]*GameDB_Spell
	Effects				map[int64]*GameDB_Effect
	Procs				map[int64]*GameDB_Proc

	Items				map[int64]*GameDB_Item
	
	Spellbook			[]*GameDB_Spell
}

type GameDB_Item struct {
	Id					int64
	Name 				string
	Description 		string

	Type				GameDB_Item_Type
	Slot				GameDB_Item_Slot

	DamageMin			int32
	DamageMax			int32
	AttackSpeed			float32
	Range 				float32

	BlockValue			int32
}

type GameDB_Item_Type int
const (
	GameDB_Item_Type_Weapon_OneHand = 0;
	GameDB_Item_Type_Weapon_MainHand = 1;
	GameDB_Item_Type_Weapon_OffHand = 2;
	GameDB_Item_Type_Weapon_TwoHand = 3;
	GameDB_Item_Type_Weapon_Shield = 4;
	GameDB_Item_Type_Weapon_Bow = 5;
	GameDB_Item_Type_Weapon_Wand = 6;
)

type GameDB_Item_Slot int
const (
	GameDB_Item_Slot_Weapon_MainHand = 0;
	GameDB_Item_Slot_Weapon_OffHand = 1;
)

type GameDB_Proc struct {
	Spell 				*GameDB_Spell
	Proc				[]*GameDB_Proc_Event
	Chance				float32
}

type GameDB_Proc_Event int8
const (
	GameDB_Proc_Event_OnDeath = 0
	GameDB_Proc_Event_AutoAttack = 1
	GameDB_Proc_Event_OnDamageDone = 2
	GameDB_Proc_Event_OnDamageReceived = 3
	GameDB_Proc_Event_OnPhysicalAttackDone = 4
	GameDB_Proc_Event_OnPhysicalAttackReceived = 5
	GameDB_Proc_Event_OnMagicalAttackDone = 6
	GameDB_Proc_Event_OnMagicalAttackReceived = 7
	GameDB_Proc_Event_OnHealingDone = 8
	GameDB_Proc_Event_OnHealingReceived = 9
)

type GameDB_Spell struct {
	Id					int64
	Name 				string
	Description 		string
	Visible				bool
	ThreadModifier		int32
	Cooldown			float32
	GlobalCooldown		float32
	IgnoresGCD			bool
	IgnoresWeaponswing  bool

	MissileID			int32
	EffectID			int32
	IconID		 		int64
	Speed				float32
	ApplicationType     GameDB_Spell_ApplicationType

	BaseCost			int32
	CostPerSec			int32
	CostPercentage  	int32
	
	CastTime 			float32
	Range				float32
	FacingFront			bool

	TargetAuraRequired 	int64
	CasterAuraRequired 	int64
	
	Target				GameDB_Spell_Target

	Effect			    []*GameDB_Effect
	
	InterruptedBy		GameDB_Interrupt
}

type GameDB_Spell_ApplicationType int
const (
	GameDB_Spell_ApplicationType_WeaponSwing = 0
	GameDB_Spell_ApplicationType_Instant = 1
	GameDB_Spell_ApplicationType_Missile = 2
	GameDB_Spell_ApplicationType_Beam = 3
	GameDB_Spell_ApplicationType_AoE = 4
	GameDB_Spell_ApplicationType_Cone = 5
	GameDB_Spell_ApplicationType_Summon = 6
)

type GameDB_Effect struct {
	Id				int64
	Name 			string
	Description 	string
	Visible			bool
	EffectID		int64
	Duration 		float32
	Dispellable		bool
	School 			GameDB_Spell_SchoolType
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

type GameDB_Spell_Target int8
const (
	GameDB_Spell_Target_None = 0
	GameDB_Spell_Target_Unit = 1
	GameDB_Spell_Target_Enemy = 2
	GameDB_Spell_Target_Ally = 3
	GameDB_Spell_Target_Dead = 4
	GameDB_Spell_Target_DeadEnemy = 5
	GameDB_Spell_Target_DeadAlly = 6
	GameDB_Spell_Target_AoE = 7
)

type GameDB_Interrupt int8
const (
	GameDB_Interrupt_None = 0
	GameDB_Interrupt_OnMovement = 1
	GameDB_Interrupt_OnKnockback = 2
	GameDB_Interrupt_OnInterruptCast = 3
	GameDB_Interrupt_OnDamageTaken = 4
	GameDB_Interrupt_OnAttackingMeele = 5
	GameDB_Interrupt_OnAttackingSpell = 6
)


func (g *GameDB) searchSpellsByName(name string) (*GameDB_Spell) {
	for _, spell := range g.Spells { 
		if spell.Name == name {
			return spell;
		}
	}
	return nil
}
func (g *GameDB) searchEffectByName(name string) (*GameDB_Effect) {
	for _, effect := range g.Effects { 
		if effect.Name == name {
			return effect;
		}
	}
	return nil
}