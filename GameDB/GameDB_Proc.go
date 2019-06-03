package main

type GameDB_Proc struct {	
	Id					int64
	Name 				string
	Description 		string
	Visible				bool

	OnEvent				[]*GameDB_Proc_Event
	Chance				float32
	Cooldown			float32

	SpellTriggered		*GameDB_Spell
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
