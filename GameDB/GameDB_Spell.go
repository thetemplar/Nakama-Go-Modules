package main

type GameDB_Spell struct {
	Id					int64
	Name 				string
	Description 		string
	Visible				bool
	ThreadModifier		float32
	Cooldown			float32
	GlobalCooldown		float32
	IgnoresGCD			bool
	IgnoresWeaponswing  bool

	MissileID			int32
	EffectID			int32
	IconID		 		int64
	Speed				float32
	Application_Type    GameDB_Spell_Application_Type

	BaseCost			int32
	CostPerSec			int32
	CostPercentage  	int32
	
	CastTime 			float32
	Range				float32
	FacingFront			bool

	TargetAuraRequired 	int64
	CasterAuraRequired 	int64
	
	Target_Type			GameDB_Spell_Target_Type
	
	InterruptedBy		GameDB_Interrupt_Type

	ApplyEffect			[]*GameDB_Effect
	ApplyProc			[]*GameDB_Proc
}

type GameDB_Spell_Application_Type int
const (
	GameDB_Spell_Application_Type_WeaponSwing = 0
	GameDB_Spell_Application_Type_Instant = 1
	GameDB_Spell_Application_Type_Missile = 2
	GameDB_Spell_Application_Type_Beam = 3
	GameDB_Spell_Application_Type_AoE = 4
	GameDB_Spell_Application_Type_Cone = 5
	GameDB_Spell_Application_Type_Summon = 6
)

type GameDB_Spell_Target_Type int8
const (
	GameDB_Spell_Target_Type_None = 0
	GameDB_Spell_Target_Type_Unit = 1
	GameDB_Spell_Target_Type_Enemy = 2
	GameDB_Spell_Target_Type_Ally = 3
	GameDB_Spell_Target_Type_Dead = 4
	GameDB_Spell_Target_Type_DeadEnemy = 5
	GameDB_Spell_Target_Type_DeadAlly = 6
	GameDB_Spell_Target_Type_AoE = 7
)

type GameDB_Interrupt_Type int8
const (
	GameDB_Interrupt_Type_None = 0
	GameDB_Interrupt_Type_OnMovement = 1
	GameDB_Interrupt_Type_OnKnockback = 2
	GameDB_Interrupt_Type_OnInterruptCast = 3
	GameDB_Interrupt_Type_OnDamageTaken = 4
	GameDB_Interrupt_Type_OnAttackingMeele = 5
	GameDB_Interrupt_Type_OnAttackingSpell = 6
)