package main

import (
	"Nakama-Go-Modules/GameDB"
)

//base stats
func (p *InternalInteractable)  getCurrentStamina(dbc *GameDB.Class) float32 {
	return dbc.BaseStamina + dbc.GainStamina * float32(p.Level)
}

func (p *InternalInteractable)  getCurrentStrength(dbc *GameDB.Class) float32 {
	return dbc.BaseStrength + dbc.GainStrength * float32(p.Level)
}

func (p *InternalInteractable)  getCurrentAgility(dbc *GameDB.Class) float32 {
	return dbc.BaseAgility + dbc.GainAgility * float32(p.Level)
}

func (p *InternalInteractable)  getCurrentIntellect(dbc *GameDB.Class) float32 {
	return dbc.BaseIntellect + dbc.GainIntellect * float32(p.Level)
}

func (p *InternalInteractable)  getMaxHp(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * 10
}
func (p *InternalInteractable)  getMaxMana(dbc *GameDB.Class) float32 {
	return p.getCurrentIntellect(dbc) * 10
}

func (p *InternalInteractable)  getHpRegen(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorHPRegen
}
func (p *InternalInteractable)  getManaRegen(dbc *GameDB.Class) float32 {
	return p.getCurrentIntellect(dbc) * dbc.FactorManaRegen
}

//meele
func (p *InternalInteractable)  getMeeleAttackSpeed(dbc *GameDB.Class) float32 {	
	return 1-(1/(200/(p.getCurrentAgility(dbc) * dbc.FactorMeeleAttackSpeed))) * 0.25 
}

func (p *InternalInteractable)  getMeeleAttackPower(dbc *GameDB.Class) float32 {
	return p.getCurrentAgility(dbc) * dbc.FactorAgilityAP + p.getCurrentStrength(dbc) * dbc.FactorStrengthAP
}

func (p *InternalInteractable)  getMeeleCritChance(dbc *GameDB.Class) float32 {
	return (1/(200/(p.getCurrentAgility(dbc) * dbc.FactorMeeleCriticalHits))) * 0.25 
}
func (p *InternalInteractable)  getMeeleHitChance(dbc *GameDB.Class) float32 {
	return float32(0.8) + float32(p.Level) * float32(0.01)
}

//spell
func (p *InternalInteractable)  getSpellAttackSpeed(dbc *GameDB.Class) float32 {
	return 1-(1/(200/(p.getCurrentIntellect(dbc) * dbc.FactorSpellAttackSpeed))) * 0.25 
}
func (p *InternalInteractable)  getSpellAttackPower(dbc *GameDB.Class) float32 {
	return p.getCurrentIntellect(dbc) * dbc.FactorSpellAP
}
func (p *InternalInteractable)  getSpellCritChance(dbc *GameDB.Class) float32 {
	return (1/(200/(p.getCurrentIntellect(dbc) * dbc.FactorSpellCriticalHits))) * 0.25 
}
func (p *InternalInteractable)  getSpellHitChance(dbc *GameDB.Class) float32 {
	return float32(0.8) + float32(p.Level) * float32(0.01)
}

//armor
func (p *InternalInteractable)  getArmor(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorArmor
}
func (p *InternalInteractable)  getBlockPercentage(dbc *GameDB.Class) float32 {
	return (1/(200/(p.getCurrentStamina(dbc) * dbc.FactorBlock))) * 0.75 
}
func (p *InternalInteractable)  getDodgeChance(dbc *GameDB.Class) float32 {
	return (p.getCurrentAgility(dbc) * dbc.FactorDodge / 100) / 4 /*4=magic number*/
}
func (p *InternalInteractable)  getParryChance(dbc *GameDB.Class) float32 {
	return (p.getCurrentStrength(dbc) * dbc.FactorParry / 100) / 4 /*4=magic number*/
}

//resistance
func (p *InternalInteractable)  getResistanceArcane(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (p *InternalInteractable)  getResistanceFire(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (p *InternalInteractable)  getResistanceFrost(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (p *InternalInteractable)  getResistanceNature(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (p *InternalInteractable)  getResistanceShadow(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}
func (p *InternalInteractable)  getResistanceHoly(dbc *GameDB.Class) float32 {
	return p.getCurrentStamina(dbc) * dbc.FactorSpellResist
}

func (p *InternalInteractable)  getResistance(dbc *GameDB.Class, school GameDB.Spell_SchoolType) float32 {
	switch school {
	case GameDB.Spell_SchoolType_Arcane:
		return p.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Fire:
		return p.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Frost:
		return p.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Nature:
		return p.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Shadow:
		return p.getResistanceArcane(dbc)
	case GameDB.Spell_SchoolType_Holy:
		return p.getResistanceArcane(dbc)
	}
	return float32(0)
}
