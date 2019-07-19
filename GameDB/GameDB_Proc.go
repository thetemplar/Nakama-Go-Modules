package GameDB

type Proc struct {	
	Id					int64
	Name 				string
	Description 		string
	Visible				bool

	OnEvent				*Proc_Event
	Chance				float32
	Cooldown			float32

	//SpellTriggered		*Spell
}

type Proc_Event int8
const (
	Proc_Event_OnDeath = 0
	Proc_Event_AutoAttack = 1
	Proc_Event_OnDamageDone = 2
	Proc_Event_OnDamageReceived = 3
	Proc_Event_OnPhysicalAttackDone = 4
	Proc_Event_OnPhysicalAttackReceived = 5
	Proc_Event_OnMagicalAttackDone = 6
	Proc_Event_OnMagicalAttackReceived = 7
	Proc_Event_OnHealingDone = 8
	Proc_Event_OnHealingReceived = 9
)
