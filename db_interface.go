package main

type GameDB struct {
	Spells				map[int64]*GameDB_Spell
	Effects				map[int64]*GameDB_Effect
	Procs				map[int64]*GameDB_Proc
	
	Spellbook			[]*GameDB_Spell
}

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
	GameDB_Spell_ApplicationType_Instant = 0
	GameDB_Spell_ApplicationType_Missile = 1
	GameDB_Spell_ApplicationType_Beam = 2
	GameDB_Spell_ApplicationType_AoE = 3
	GameDB_Spell_ApplicationType_Cone = 4
	GameDB_Spell_ApplicationType_Summon = 5
)

type GameDB_Effect struct {
	Id				int64
	Name 			string
	Description 	string
	Visible			bool
	EffectID		int64
	Duration 		float32
	Dispellable		bool
	Type 			interface{}
}

type GameDB_Effect_Damage struct {
	Type 			GameDB_Spell_DamageType
	ValueMin 		int32
	ValueMax 		int32
}

type GameDB_Effect_Heal struct {
	ValueMin 		int32
	ValueMax 		int32
}

type GameDB_Effect_Apply_Aura_Periodic_Damage struct {
	Type 			GameDB_Spell_DamageType
	ValueMin 		int32
	ValueMax 		int32
	Intervall   	float32
}

type GameDB_Effect_Apply_Aura_Periodic_Heal struct {
	ValueMin 		int32
	ValueMax 		int32
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

type GameDB_Spell_DamageType int
const (
	GameDB_Spell_DamageType_Physical = 0
	GameDB_Spell_DamageType_Arcane = 1
	GameDB_Spell_DamageType_Fire = 2
	GameDB_Spell_DamageType_Frost = 3
	GameDB_Spell_DamageType_Nature = 4
	GameDB_Spell_DamageType_Shadow = 5
	GameDB_Spell_DamageType_Holy = 6
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