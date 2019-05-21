package main

//hp/mana
func (c Character) getHpBonus() float32 {
	return float32(c.BaseStats.Stamina) * 10
}
func (c Character) getManaBonus() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Ranger: res += 15 * float32(c.BaseStats.Intellect)
	case Character_Warlock: res += 15 * float32(c.BaseStats.Intellect)
	case Character_Wizard: res += 15 * float32(c.BaseStats.Intellect)
	case Character_Cleric: res += 15 * float32(c.BaseStats.Intellect)
	case Character_Priest: res += 15 * float32(c.BaseStats.Intellect)
	}
	return res
}
func (c Character) getHpRegen() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior: res += 0.48 + 0.8 * float32(c.BaseStats.Wisdom)
	case Character_Ranger: res += 0.6 + 1 * float32(c.BaseStats.Wisdom)
	case Character_Warlock: res += 0.42 + 0.7 * float32(c.BaseStats.Wisdom)
	case Character_Wizard: res += 1 + 0.7 * float32(c.BaseStats.Wisdom)
	case Character_Cleric: res += 1 + 0.7 * float32(c.BaseStats.Wisdom)
	case Character_Rouge: res += 0.36 + 0.6 * float32(c.BaseStats.Wisdom)
	case Character_Priest: res += 0.65 + 1 * float32(c.BaseStats.Wisdom)
	}
	return res
}
func (c Character) getManaRegen() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Ranger: res += 15 + 0.2 * float32(c.BaseStats.Wisdom)
	case Character_Warlock: res += 8 + 0.25 * float32(c.BaseStats.Wisdom)
	case Character_Wizard: res += 13 + 0.25 * float32(c.BaseStats.Wisdom)
	case Character_Cleric: res += 13 + 0.25 * float32(c.BaseStats.Wisdom)
	case Character_Priest:  res += 13 + 0.25 * float32(c.BaseStats.Wisdom)
	}
	return res
}

//meele
func (c Character) getMeeleAttackSpeed() float32 {
	res := float32(1)
	switch c.Class {
	case Character_Warrior: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Ranger: res += 0.02 * float32(c.BaseStats.Agility)
	case Character_Warlock:  res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Wizard: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Cleric: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Rouge: res += 0.03 * float32(c.BaseStats.Agility)
	case Character_Priest: res += 0.05 * float32(c.BaseStats.Agility)
	}
	return res
}
func (c Character) getMeeleAttackPower() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior: res += 2 * float32(c.BaseStats.Strength)
	case Character_Ranger: res += 1 * float32(c.BaseStats.Strength) + 1 * float32(c.BaseStats.Agility) 
	case Character_Warlock: res += 1 * float32(c.BaseStats.Strength)
	case Character_Wizard: res += 1 * float32(c.BaseStats.Strength)
	case Character_Cleric: res += 2 * float32(c.BaseStats.Strength)
	case Character_Rouge: res += 1 * float32(c.BaseStats.Strength) + 1 * float32(c.BaseStats.Agility) 
	case Character_Priest: res += 1 * float32(c.BaseStats.Strength)
	}
	return res
}
func (c Character) getMeeleCritChance() float32 {
	res := float32(0.05)
	switch c.Class {
	case Character_Warrior: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Ranger: res += 0.02 * float32(c.BaseStats.Agility)
	case Character_Warlock:  res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Wizard: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Cleric: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Rouge: res += 0.03 * float32(c.BaseStats.Agility)
	case Character_Priest: res += 0.05 * float32(c.BaseStats.Agility)
	}
	return res
}
func (c Character) getMeeleHitChance() float32 {
	res := float32(0.80)
	switch c.Class {
	case Character_Warrior: 
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}

//spell
func (c Character) getSpellAttackSpeed() float32 {
	res := float32(1)
	switch c.Class {
	case Character_Warlock: res += 0.02 * float32(c.BaseStats.Intellect)
	case Character_Wizard: res += 0.02 * float32(c.BaseStats.Intellect)
	case Character_Cleric: res += 0.03 * float32(c.BaseStats.Intellect)
	case Character_Priest: res += 0.02 * float32(c.BaseStats.Intellect)
	}
	return res
}
func (c Character) getSpellAttackPower() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getSpellCritChance() float32 {
	res := float32(0.05)
	switch c.Class {
	case Character_Warlock: res += 0.02 * float32(c.BaseStats.Intellect)
	case Character_Wizard: res += 0.02 * float32(c.BaseStats.Intellect)
	case Character_Cleric: res += 0.03 * float32(c.BaseStats.Intellect)
	case Character_Priest: res += 0.02 * float32(c.BaseStats.Intellect)
	}
	return res
}
func (c Character) getSpellHitChance() float32 {
	res := float32(0.90)
	switch c.Class {
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}

//armor
func (c Character) getArmor() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior: res += 2 * float32(c.BaseStats.Agility)
	case Character_Ranger: res += 2 * float32(c.BaseStats.Agility)
	case Character_Warlock: res += 2 * float32(c.BaseStats.Agility)
	case Character_Wizard: res += 2 * float32(c.BaseStats.Agility)
	case Character_Cleric: res += 2 * float32(c.BaseStats.Agility)
	case Character_Rouge: res += 2 * float32(c.BaseStats.Agility)
	case Character_Priest: res += 2 * float32(c.BaseStats.Agility)
	}
	return res
}
func (c Character) getBlockPercentage() float32 {
	res := float32(0.10)
	switch c.Class {
	case Character_Warrior: res += 0.05 * float32(c.BaseStats.Strength)
	case Character_Cleric: res += 0.05 * float32(c.BaseStats.Strength)
	}
	return res
}
func (c Character) getDodgeChance() float32 {
	res := float32(0.10)
	switch c.Class {
	case Character_Warrior: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Ranger: res += 0.04 * float32(c.BaseStats.Agility)
	case Character_Warlock: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Wizard: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Cleric: res += 0.05 * float32(c.BaseStats.Agility)
	case Character_Rouge: res += 0.07 * float32(c.BaseStats.Agility)
	case Character_Priest: res += 0.05 * float32(c.BaseStats.Agility)
	}
	return res
}
func (c Character) getParryChance() float32 {
	res := float32(0.10)
	switch c.Class {
	case Character_Warrior: res += 0.05 * float32(c.BaseStats.Strength)
	case Character_Ranger: res += 0.04 * float32(c.BaseStats.Strength)
	case Character_Warlock: res += 0.05 * float32(c.BaseStats.Strength)
	case Character_Wizard: res += 0.05 * float32(c.BaseStats.Strength)
	case Character_Cleric: res += 0.05 * float32(c.BaseStats.Strength)
	case Character_Rouge: res += 0.07 * float32(c.BaseStats.Strength)
	case Character_Priest: res += 0.05 * float32(c.BaseStats.Strength)
	}
	return res
}

//resistance
func (c Character) getResistanceArcane() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistanceFire() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistanceFrost() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistanceNature() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistanceShadow() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistanceHoly() float32 {
	res := float32(0)
	switch c.Class {
	case Character_Warrior:
	case Character_Ranger:
	case Character_Warlock:
	case Character_Wizard:
	case Character_Cleric:
	case Character_Rouge:
	case Character_Priest:
	}
	return res
}
func (c Character) getResistance(school GameDB_Spell_SchoolType) float32 {
	switch school {
	case GameDB_Spell_SchoolType_Arcane:
		return c.getResistanceArcane()
	case GameDB_Spell_SchoolType_Fire:
		return c.getResistanceArcane()
	case GameDB_Spell_SchoolType_Frost:
		return c.getResistanceArcane()
	case GameDB_Spell_SchoolType_Nature:
		return c.getResistanceArcane()
	case GameDB_Spell_SchoolType_Shadow:
		return c.getResistanceArcane()
	case GameDB_Spell_SchoolType_Holy:
		return c.getResistanceArcane()
	}
	return float32(0)
}
