package GameDB

type Spell struct {
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
	Application_Type    Spell_Application_Type

	BaseCost			int32
	CostPerSec			int32
	CostPercentage  	int32
	
	CastTime 			float32
	CastTimeChanneled	float32
	Range				float32
	NeedLoS				bool
	FacingFront			bool

	TargetAuraRequired 	int64
	CasterAuraRequired 	int64
	
	Target_Type			Spell_Target_Type
	
	InterruptedBy		Interrupt_Type

	ApplyEffect			[]*Effect
	ApplyProc			[]*Proc
}

type Spell_Application_Type int
const (
	Spell_Application_Type_WeaponSwing = 0
	Spell_Application_Type_Instant = 1
	Spell_Application_Type_Missile = 2
	Spell_Application_Type_Beam = 3
	Spell_Application_Type_AoE = 4
	Spell_Application_Type_Cone = 5
	Spell_Application_Type_Summon = 6
	Spell_Application_Type_Teleport = 7
)

type Spell_Target_Type int8
const (
	Spell_Target_Type_None = 0
	Spell_Target_Type_Unit = 1
	Spell_Target_Type_Self = 2
	Spell_Target_Type_Enemy = 3
	Spell_Target_Type_Ally = 4
	Spell_Target_Type_Dead = 5
	Spell_Target_Type_DeadEnemy = 6
	Spell_Target_Type_DeadAlly = 7
	Spell_Target_Type_AoE = 8
)

type Interrupt_Type int8
const (
	Interrupt_Type_None = 0
	Interrupt_Type_OnMovement = 1
	Interrupt_Type_OnKnockback = 2
	Interrupt_Type_OnInterruptCast = 3
	Interrupt_Type_OnDamageTaken = 4
	Interrupt_Type_OnAttackingMeele = 5
	Interrupt_Type_OnAttackingSpell = 6
)