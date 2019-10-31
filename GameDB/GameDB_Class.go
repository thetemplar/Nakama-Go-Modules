package GameDB

type Class struct {
	Name 					string
	Description 			string

	Spells					[]*Spell
	Mainhand				*Item
	Offhand					*Item
	Items					[]*Item
	Procs					[]*Proc

	BaseStamina				float32
	GainStamina				float32
	FactorHPRegen			float32
	FactorArmor				float32
	FactorSpellResist		float32
	FactorBlock				float32

	BaseStrength			float32
	GainStrength			float32
	FactorStrengthAP		float32
	FactorParry				float32

	BaseAgility				float32
	GainAgility				float32
	FactorAgilityAP			float32
	FactorMeeleAttackSpeed	float32
	FactorMeeleCriticalHits	float32
	FactorDodge				float32

	BaseIntellect			float32
	GainIntellect			float32
	FactorManaRegen			float32
	FactorSpellAP			float32
	FactorSpellAttackSpeed	float32
	FactorSpellCriticalHits	float32

	MovementSpeed			float32
}