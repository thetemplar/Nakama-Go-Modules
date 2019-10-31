package main

import (
	"Nakama-Go-Modules/GameDB"
)

//base stats
func (c *Character)  getCurrentStamina(dbc *GameDB.Class) float32 {
	return dbc.BaseStamina + dbc.GainStamina * float32(c.Level)
}

func (c *Character)  getCurrentStrength(dbc *GameDB.Class) float32 {
	return dbc.BaseStrength + dbc.GainStrength * float32(c.Level)
}

func (c *Character)  getCurrentAgility(dbc *GameDB.Class) float32 {
	return dbc.BaseAgility + dbc.GainAgility * float32(c.Level)
}

func (c *Character)  getCurrentIntellect(dbc *GameDB.Class) float32 {
	return dbc.BaseIntellect + dbc.GainIntellect * float32(c.Level)
}

func (c *Character)  getMaxHp(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * 10
}
func (c *Character)  getMaxMana(dbc *GameDB.Class) float32 {
	return c.getCurrentIntellect(dbc) * 10
}

func (c *Character)  getHpRegen(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorHPRegen
}
func (c *Character)  getManaRegen(dbc *GameDB.Class) float32 {
	return c.getCurrentIntellect(dbc) * dbc.FactorManaRegen
}

//meele
func (c *Character)  getMeeleAttackSpeed(dbc *GameDB.Class) float32 {	
	return 1-(1/(200/(c.getCurrentAgility(dbc) * dbc.FactorMeeleAttackSpeed))) * 0.25 
}

func (c *Character)  getMeeleAttackPower(dbc *GameDB.Class) float32 {
	return c.getCurrentAgility(dbc) * dbc.FactorAgilityAP + c.getCurrentStrength(dbc) * dbc.FactorStrengthAP
}

func (c *Character)  getMeeleCritChance(dbc *GameDB.Class) float32 {
	return (1/(200/(c.getCurrentAgility(dbc) * dbc.FactorMeeleCriticalHits))) * 0.25 
}
func (c *Character)  getMeeleHitChance(dbc *GameDB.Class) float32 {
	return float32(0.8) + float32(c.Level) * float32(0.01)
}

//spell
func (c *Character)  getSpellAttackSpeed(dbc *GameDB.Class) float32 {
	return 1-(1/(200/(c.getCurrentIntellect(dbc) * dbc.FactorSpellAttackSpeed))) * 0.25 
}
func (c *Character)  getSpellAttackPower(dbc *GameDB.Class) float32 {
	return c.getCurrentIntellect(dbc) * dbc.FactorSpellAP
}
func (c *Character)  getSpellCritChance(dbc *GameDB.Class) float32 {
	return (1/(200/(c.getCurrentIntellect(dbc) * dbc.FactorSpellCriticalHits))) * 0.25 
}
func (c *Character)  getSpellHitChance(dbc *GameDB.Class) float32 {
	return float32(0.8) + float32(c.Level) * float32(0.01)
}

//armor
func (c *Character)  getArmor(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorArmor
}
func (c *Character)  getBlockPercentage(dbc *GameDB.Class) float32 {
	return (1/(200/(c.getCurrentStamina(dbc) * dbc.FactorBlock))) * 0.75 
}
func (c *Character)  getDodgeChance(dbc *GameDB.Class) float32 {
	return (c.getCurrentAgility(dbc) * dbc.FactorDodge / 100) / 4 /*4=magic number*/
}
func (c *Character)  getParryChance(dbc *GameDB.Class) float32 {
	return (c.getCurrentStrength(dbc) * dbc.FactorParry / 100) / 4 /*4=magic number*/
}

//resistance
func (c *Character)  getResistanceArcane(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (c *Character)  getResistanceFire(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (c *Character)  getResistanceFrost(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (c *Character)  getResistanceNature(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (c *Character)  getResistanceShadow(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (c *Character)  getResistanceHoly(dbc *GameDB.Class) float32 {
	return c.getCurrentStamina(dbc) * dbc.FactorSpellResist
}

func (c *Character)  getResistance(dbc *GameDB.Class, school GameDB.Spell_SchoolType) float32 {
	switch school {
	case GameDB.Spell_SchoolType_Arcane:
		return c.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Fire:
		return c.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Frost:
		return c.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Nature:
		return c.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Shadow:
		return c.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Holy:
		return c.getResistanceArcane(dbc)
	}
	return float32(0)
}
