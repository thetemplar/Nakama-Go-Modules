package main

type GameDB_Class struct {
	Name 					string
	Description 			string

	Spells					[]*GameDB_Spell
	Items					[]*GameDB_Item
	Procs					[]*GameDB_Proc

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


//base stats
func (dbc *GameDB_Class)  getCurrentStamina(c *Character) float32 {
	return dbc.BaseStamina + dbc.GainStamina * float32(c.Level)
}

func (dbc *GameDB_Class)  getCurrentStrength(c *Character) float32 {
	return dbc.BaseStrength + dbc.GainStrength * float32(c.Level)
}

func (dbc *GameDB_Class)  getCurrentAgility(c *Character) float32 {
	return dbc.BaseAgility + dbc.GainAgility * float32(c.Level)
}

func (dbc *GameDB_Class)  getCurrentIntellect(c *Character) float32 {
	return dbc.BaseIntellect + dbc.GainIntellect * float32(c.Level)
}

func (dbc *GameDB_Class)  getHpRegen(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorHPRegen
}
func (dbc *GameDB_Class)  getManaRegen(c *Character) float32 {
	return dbc.getCurrentIntellect(c) * dbc.FactorManaRegen
}

//meele
func (dbc *GameDB_Class)  getMeeleAttackSpeed(c *Character) float32 {	
	return dbc.getCurrentAgility(c) * dbc.FactorMeeleAttackSpeed
}

func (dbc *GameDB_Class)  getMeeleAttackPower(c *Character) float32 {
	return dbc.getCurrentAgility(c) * dbc.FactorAgilityAP + dbc.getCurrentStrength(c) * dbc.FactorStrengthAP
}

func (dbc *GameDB_Class)  getMeeleCritChance(c *Character) float32 {
	return dbc.getCurrentAgility(c) * dbc.FactorMeeleCriticalHits
}
func (dbc *GameDB_Class)  getMeeleHitChance(c *Character) float32 {
	return float32(0.8) + float32(c.Level) * float32(0.01)
}

//spell
func (dbc *GameDB_Class)  getSpellAttackSpeed(c *Character) float32 {
	return dbc.getCurrentIntellect(c) * dbc.FactorSpellAttackSpeed
}
func (dbc *GameDB_Class)  getSpellAttackPower(c *Character) float32 {
	return dbc.getCurrentIntellect(c) * dbc.FactorSpellAP
}
func (dbc *GameDB_Class)  getSpellCritChance(c *Character) float32 {
	return dbc.getCurrentIntellect(c) * dbc.FactorSpellCriticalHits
}
func (dbc *GameDB_Class)  getSpellHitChance(c *Character) float32 {
	return float32(0.8) + float32(c.Level) * float32(0.01)
}

//armor
func (dbc *GameDB_Class)  getArmor(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorArmor
}
func (dbc *GameDB_Class)  getBlockPercentage(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorBlock
}
func (dbc *GameDB_Class)  getDodgeChance(c *Character) float32 {
	return dbc.getCurrentAgility(c) * dbc.FactorDodge
}
func (dbc *GameDB_Class)  getParryChance(c *Character) float32 {
	return dbc.getCurrentStrength(c) * dbc.FactorParry
}

//resistance
func (dbc *GameDB_Class)  getResistanceArcane(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}
func (dbc *GameDB_Class)  getResistanceFire(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}
func (dbc *GameDB_Class)  getResistanceFrost(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}
func (dbc *GameDB_Class)  getResistanceNature(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}
func (dbc *GameDB_Class)  getResistanceShadow(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}
func (dbc *GameDB_Class)  getResistanceHoly(c *Character) float32 {
	return dbc.getCurrentStamina(c) * dbc.FactorSpellResist
}

func (dbc *GameDB_Class)  getResistance(c *Character, school GameDB_Spell_SchoolType) float32 {
	switch school {
	case GameDB_Spell_SchoolType_Arcane:
		return dbc.getResistanceArcane(c)
	case GameDB_Spell_SchoolType_Fire:
		return dbc.getResistanceArcane(c)
	case GameDB_Spell_SchoolType_Frost:
		return dbc.getResistanceArcane(c)
	case GameDB_Spell_SchoolType_Nature:
		return dbc.getResistanceArcane(c)
	case GameDB_Spell_SchoolType_Shadow:
		return dbc.getResistanceArcane(c)
	case GameDB_Spell_SchoolType_Holy:
		return dbc.getResistanceArcane(c)
	}
	return float32(0)
}
