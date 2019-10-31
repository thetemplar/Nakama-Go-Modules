package main

import (
	"Nakama-Go-Modules/GameDB"
)

func init_db() *GameDB.Database {
	GameDatabase := &GameDB.Database{
		Spells:  make(map[int64]*GameDB.Spell),
		Effects: make(map[int64]*GameDB.Effect),
		Procs:   make(map[int64]*GameDB.Proc),
		Items:   make(map[int64]*GameDB.Item),
		Classes: make(map[string]*GameDB.Class),
	}

	//ITEMS
	GameDB_Item_Sword := &GameDB.Item{
		Id:          1,
		Name:        "Sword",
		Description: "Simple Sword",
		Type:        GameDB.Item_Type_Weapon_OneHand,
		Slot:        GameDB.Item_Slot_Weapon_MainHand,
		DamageMin:   15,
		DamageMax:   25,
		AttackSpeed: 2.8,
		Range:       2,
		BlockValue:  0,
	}
	GameDatabase.Items[GameDB_Item_Sword.Id] = GameDB_Item_Sword
	GameDB_Item_Twohander := &GameDB.Item{
		Id:          2,
		Name:        "Twohander",
		Description: "Simple Two Handed Sword",
		Type:        GameDB.Item_Type_Weapon_TwoHand,
		Slot:        GameDB.Item_Slot_Weapon_BothHands,
		DamageMin:   25,
		DamageMax:   35,
		AttackSpeed: 3.7,
		Range:       2,
		BlockValue:  0,
	}
	GameDatabase.Items[GameDB_Item_Twohander.Id] = GameDB_Item_Twohander
	GameDB_Item_Staff := &GameDB.Item{
		Id:          3,
		Name:        "Staff",
		Description: "Simple Caster Staff",
		Type:        GameDB.Item_Type_Weapon_TwoHand,
		Slot:        GameDB.Item_Slot_Weapon_BothHands,
		DamageMin:   15,
		DamageMax:   20,
		AttackSpeed: 3,
		Range:       2,
		BlockValue:  0,
	}
	GameDatabase.Items[GameDB_Item_Staff.Id] = GameDB_Item_Staff

	//EFFECTS
	GameDB_Effect_Fireball := &GameDB.Effect{
		Id:          1,
		Name:        "Fireball",
		Description: "Fireball",
		Visible:     true,
		EffectID:    0,
		Duration:    0,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Fire,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type:        &GameDB.Effect_Damage{},
		ValueMin:    20,
		ValueMax:    30,
	}
	GameDatabase.Effects[GameDB_Effect_Fireball.Id] = GameDB_Effect_Fireball
	GameDB_Effect_Frostbolt := &GameDB.Effect{
		Id:          2,
		Name:        "Frostbolt",
		Description: "Frostbolt",
		Visible:     true,
		EffectID:    0,
		Duration:    0,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Frost,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type:        &GameDB.Effect_Damage{},
		ValueMin:    15,
		ValueMax:    20,
	}
	GameDatabase.Effects[GameDB_Effect_Frostbolt.Id] = GameDB_Effect_Frostbolt
	GameDB_Effect_Chilled := &GameDB.Effect{
		Id:          3,
		Name:        "Chilled",
		Description: "Chilled",
		Visible:     true,
		EffectID:    0,
		Duration:    5,
		Dispellable: true,
		School:      GameDB.Spell_SchoolType_Frost,
		Mechanic:    GameDB.Spell_Mechanic_Slowed,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Speed,
			Value: 0.75,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_Chilled.Id] = GameDB_Effect_Chilled
	GameDB_Effect_Sunburn := &GameDB.Effect{
		Id:          4,
		Name:        "Sunburn",
		Description: "Sunburn",
		Visible:     true,
		EffectID:    0,
		Duration:    0,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Fire,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type:        &GameDB.Effect_Damage{},
		ValueMin:    5,
		ValueMax:    10,
	}
	GameDatabase.Effects[GameDB_Effect_Sunburn.Id] = GameDB_Effect_Sunburn
	GameDB_Effect_Sunburned := &GameDB.Effect{
		Id:          5,
		Name:        "Sunburned",
		Description: "Sunburn DoT",
		Visible:     true,
		EffectID:    0,
		Duration:    10,
		Dispellable: true,
		School:      GameDB.Spell_SchoolType_Fire,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Periodic_Damage{
			Intervall: 2,
		},
		ValueMin: 3,
		ValueMax: 5,
	}
	GameDatabase.Effects[GameDB_Effect_Sunburned.Id] = GameDB_Effect_Sunburned
	GameDB_Effect_Enraged := &GameDB.Effect{
		Id:          6,
		Name:        "Enraged",
		Description: "Enraged",
		Visible:     true,
		EffectID:    0,
		Duration:    1000,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Physical,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_PhysicalAP,
			Value: 100,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_Enraged.Id] = GameDB_Effect_Enraged
	GameDB_Effect_Haste := &GameDB.Effect{
		Id:          7,
		Name:        "Haste",
		Description: "Haste",
		Visible:     true,
		EffectID:    0,
		Duration:    10,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Physical,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Speed,
			Value: 1.5,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_Haste.Id] = GameDB_Effect_Haste

	//PROCS

	//SPELLS
	GameDB_Spell_Fireball := &GameDB.Spell{
		Id:                 1,
		Name:               "Fireball",
		Description:        "Fireball",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           3,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             0,
		Speed:              60,
		Application_Type:   GameDB.Spell_Application_Type_Missile,
		BaseCost:           15,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           2,
		CastTimeChanneled:  0,
		Range:              50,
		FacingFront:        true,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Enemy,
		InterruptedBy:      GameDB.Interrupt_Type_OnMovement,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[1]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Fireball.Id] = GameDB_Spell_Fireball
	GameDB_Spell_Frostbolt := &GameDB.Spell{
		Id:                 2,
		Name:               "Frostbolt",
		Description:        "Frostbolt",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           0,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             0,
		Speed:              50,
		Application_Type:   GameDB.Spell_Application_Type_Missile,
		BaseCost:           10,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              20,
		FacingFront:        true,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Enemy,
		InterruptedBy:      GameDB.Interrupt_Type_OnMovement,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[2], GameDatabase.Effects[3]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Frostbolt.Id] = GameDB_Spell_Frostbolt
	GameDB_Spell_Sunburn := &GameDB.Spell{
		Id:                 3,
		Name:               "Sunburn",
		Description:        "Sunburn",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           0,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             0,
		Speed:              1000,
		Application_Type:   GameDB.Spell_Application_Type_Beam,
		BaseCost:           20,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              300,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Enemy,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[4], GameDatabase.Effects[5]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Sunburn.Id] = GameDB_Spell_Sunburn
	GameDB_Spell_Enrage := &GameDB.Spell{
		Id:                 4,
		Name:               "Enrage",
		Description:        "Enrage",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           0,
		GlobalCooldown:     1.5,
		IgnoresGCD:         true,
		IgnoresWeaponswing: true,
		MissileID:          0,
		EffectID:           0,
		IconID:             0,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           0,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              0,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_None,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[6], GameDatabase.Effects[7]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Enrage.Id] = GameDB_Spell_Enrage

	//CLASSES
	GameDB_Class_Mage := &GameDB.Class{
		Name:                    "Mage",
		Description:             "Mage",
		Spells:                  []*GameDB.Spell{GameDatabase.Spells[1], GameDatabase.Spells[2], GameDatabase.Spells[3]},
		Items:                   []*GameDB.Item{GameDatabase.Items[3]},
		Procs:                   []*GameDB.Proc{},
		BaseStamina:             10,
		GainStamina:             2,
		FactorHPRegen:           1.2,
		FactorArmor:             1,
		FactorSpellResist:       1,
		FactorBlock:             1,
		BaseStrength:            3,
		GainStrength:            1.2,
		FactorStrengthAP:        1,
		FactorParry:             1,
		BaseAgility:             4,
		GainAgility:             1.5,
		FactorAgilityAP:         1,
		FactorMeeleAttackSpeed:  1,
		FactorMeeleCriticalHits: 1,
		FactorDodge:             1,
		BaseIntellect:           20,
		GainIntellect:           2,
		FactorManaRegen:         2,
		FactorSpellAP:           3,
		FactorSpellAttackSpeed:  1,
		FactorSpellCriticalHits: 1,
		MovementSpeed:           1,
	}
	GameDatabase.Classes[GameDB_Class_Mage.Name] = GameDB_Class_Mage
	GameDB_Class_Ogre := &GameDB.Class{
		Name:                    "Ogre",
		Description:             "Ogre",
		Spells:                  []*GameDB.Spell{GameDatabase.Spells[4]},
		Items:                   []*GameDB.Item{GameDatabase.Items[2]},
		Procs:                   []*GameDB.Proc{},
		BaseStamina:             20,
		GainStamina:             2,
		FactorHPRegen:           2,
		FactorArmor:             2,
		FactorSpellResist:       2,
		FactorBlock:             2,
		BaseStrength:            15,
		GainStrength:            2,
		FactorStrengthAP:        4,
		FactorParry:             2,
		BaseAgility:             10,
		GainAgility:             1.2,
		FactorAgilityAP:         1.2,
		FactorMeeleAttackSpeed:  1.2,
		FactorMeeleCriticalHits: 2,
		FactorDodge:             1,
		BaseIntellect:           2,
		GainIntellect:           0.2,
		FactorManaRegen:         1,
		FactorSpellAP:           1,
		FactorSpellAttackSpeed:  1,
		FactorSpellCriticalHits: 1,
		MovementSpeed:           0.8,
	}
	GameDatabase.Classes[GameDB_Class_Ogre.Name] = GameDB_Class_Ogre

	return GameDatabase
}

func map_init() *Map {
	m := &Map{}

	m.Borders = make([]Edge, 98)
	m.Borders[0] = Edge{A: PublicMatchState_Vector2Df{X: 4.833333, Y: 12}, B: PublicMatchState_Vector2Df{X: 4.333333, Y: 12.5}}
	m.Borders[1] = Edge{A: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, B: PublicMatchState_Vector2Df{X: 4.833333, Y: 12}}
	m.Borders[2] = Edge{A: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: 5.666667}}
	m.Borders[3] = Edge{A: PublicMatchState_Vector2Df{X: 11, Y: 1.166667}, B: PublicMatchState_Vector2Df{X: 10.66667, Y: 0.8333334}}
	m.Borders[4] = Edge{A: PublicMatchState_Vector2Df{X: 14.83333, Y: 5.666667}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: 1.166667}}
	m.Borders[5] = Edge{A: PublicMatchState_Vector2Df{X: 14.83333, Y: 1.166667}, B: PublicMatchState_Vector2Df{X: 11, Y: 1.166667}}
	m.Borders[6] = Edge{A: PublicMatchState_Vector2Df{X: 4.166667, Y: 8}, B: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}}
	m.Borders[7] = Edge{A: PublicMatchState_Vector2Df{X: 10.66667, Y: 0.8333334}, B: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}}
	m.Borders[8] = Edge{A: PublicMatchState_Vector2Df{X: 0, Y: 7.5}, B: PublicMatchState_Vector2Df{X: 4.166667, Y: 8}}
	m.Borders[9] = Edge{A: PublicMatchState_Vector2Df{X: 17, Y: 0}, B: PublicMatchState_Vector2Df{X: 17, Y: 6.5}}
	m.Borders[10] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: 0}}
	m.Borders[11] = Edge{A: PublicMatchState_Vector2Df{X: 17, Y: 6.5}, B: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}}
	m.Borders[12] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: 24.5}, B: PublicMatchState_Vector2Df{X: 24.5, Y: 16.16667}}
	m.Borders[13] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: 24.5}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}}
	m.Borders[14] = Edge{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 14.5}, B: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}}
	m.Borders[15] = Edge{A: PublicMatchState_Vector2Df{X: 4.333333, Y: 12.5}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: 12.5}}
	m.Borders[16] = Edge{A: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 7.833333, Y: 16.16667}}
	m.Borders[17] = Edge{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 15.33333}, B: PublicMatchState_Vector2Df{X: 6.833333, Y: 14.5}}
	m.Borders[18] = Edge{A: PublicMatchState_Vector2Df{X: 3.833333, Y: 12.5}, B: PublicMatchState_Vector2Df{X: 0, Y: 12}}
	m.Borders[19] = Edge{A: PublicMatchState_Vector2Df{X: 7.833333, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 6.833333, Y: 15.33333}}
	m.Borders[20] = Edge{A: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -5.333332, Y: 0}}
	m.Borders[21] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 0}}
	m.Borders[22] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 24.5}}
	m.Borders[23] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 24.5}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}}
	m.Borders[24] = Edge{A: PublicMatchState_Vector2Df{X: -2.666666, Y: 0}, B: PublicMatchState_Vector2Df{X: -1.833332, Y: 1.333333}}
	m.Borders[25] = Edge{A: PublicMatchState_Vector2Df{X: -7.166666, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}}
	m.Borders[26] = Edge{A: PublicMatchState_Vector2Df{X: -0.666666, Y: 7.833333}, B: PublicMatchState_Vector2Df{X: 0, Y: 7.5}}
	m.Borders[27] = Edge{A: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 2.833333}}
	m.Borders[28] = Edge{A: PublicMatchState_Vector2Df{X: -1, Y: 11.5}, B: PublicMatchState_Vector2Df{X: -1, Y: 10}}
	m.Borders[29] = Edge{A: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}, B: PublicMatchState_Vector2Df{X: -6.833332, Y: 8.833334}}
	m.Borders[30] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}}
	m.Borders[31] = Edge{A: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}, B: PublicMatchState_Vector2Df{X: -8.5, Y: 3}}
	m.Borders[32] = Edge{A: PublicMatchState_Vector2Df{X: -9, Y: 18.5}, B: PublicMatchState_Vector2Df{X: -9, Y: 8.833334}}
	m.Borders[33] = Edge{A: PublicMatchState_Vector2Df{X: -9, Y: 8.833334}, B: PublicMatchState_Vector2Df{X: -11.66667, Y: 6}}
	m.Borders[34] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 18.83333}}
	m.Borders[35] = Edge{A: PublicMatchState_Vector2Df{X: -11.66667, Y: 5.333333}, B: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}}
	m.Borders[36] = Edge{A: PublicMatchState_Vector2Df{X: -11.66667, Y: 6}, B: PublicMatchState_Vector2Df{X: -11.66667, Y: 5.333333}}
	m.Borders[37] = Edge{A: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: -9, Y: 18.5}}
	m.Borders[38] = Edge{A: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, B: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}}
	m.Borders[39] = Edge{A: PublicMatchState_Vector2Df{X: -0.666666, Y: 7.833333}, B: PublicMatchState_Vector2Df{X: -1, Y: 10}}
	m.Borders[40] = Edge{A: PublicMatchState_Vector2Df{X: -8.5, Y: 3}, B: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}}
	m.Borders[41] = Edge{A: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 2.833333}}
	m.Borders[42] = Edge{A: PublicMatchState_Vector2Df{X: -1.833332, Y: 1.333333}, B: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}}
	m.Borders[43] = Edge{A: PublicMatchState_Vector2Df{X: 0, Y: 12}, B: PublicMatchState_Vector2Df{X: -1, Y: 11.5}}
	m.Borders[44] = Edge{A: PublicMatchState_Vector2Df{X: -6.833332, Y: 8.833334}, B: PublicMatchState_Vector2Df{X: -6.833332, Y: 18.5}}
	m.Borders[45] = Edge{A: PublicMatchState_Vector2Df{X: -6.833332, Y: 18.5}, B: PublicMatchState_Vector2Df{X: -7.166666, Y: 18.83333}}
	m.Borders[46] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}}
	m.Borders[47] = Edge{A: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 19.83333, Y: -20.16667}}
	m.Borders[48] = Edge{A: PublicMatchState_Vector2Df{X: 17.33333, Y: -16}, B: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}}
	m.Borders[49] = Edge{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 1.5, Y: -13.83333}}
	m.Borders[50] = Edge{A: PublicMatchState_Vector2Df{X: 6.333333, Y: -2.666666}, B: PublicMatchState_Vector2Df{X: 5.666667, Y: -2.666666}}
	m.Borders[51] = Edge{A: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}, B: PublicMatchState_Vector2Df{X: 10.83333, Y: -2.666666}}
	m.Borders[52] = Edge{A: PublicMatchState_Vector2Df{X: 17.33333, Y: -14.83333}, B: PublicMatchState_Vector2Df{X: 17.33333, Y: -16}}
	m.Borders[53] = Edge{A: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 17.33333, Y: -14.83333}}
	m.Borders[54] = Edge{A: PublicMatchState_Vector2Df{X: 19.83333, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 22.66667, Y: -19.16667}}
	m.Borders[55] = Edge{A: PublicMatchState_Vector2Df{X: 22, Y: -15.83333}, B: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}}
	m.Borders[56] = Edge{A: PublicMatchState_Vector2Df{X: 22.66667, Y: -19.16667}, B: PublicMatchState_Vector2Df{X: 22, Y: -15.83333}}
	m.Borders[57] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -13.83333}}
	m.Borders[58] = Edge{A: PublicMatchState_Vector2Df{X: 1.5, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 2, Y: -12.83333}}
	m.Borders[59] = Edge{A: PublicMatchState_Vector2Df{X: 2, Y: -12.83333}, B: PublicMatchState_Vector2Df{X: 2.333333, Y: -12.83333}}
	m.Borders[60] = Edge{A: PublicMatchState_Vector2Df{X: 2.333333, Y: -12.83333}, B: PublicMatchState_Vector2Df{X: 3.166667, Y: -13.83333}}
	m.Borders[61] = Edge{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: 5.166667, Y: -7.666666}}
	m.Borders[62] = Edge{A: PublicMatchState_Vector2Df{X: 5.166667, Y: -7.666666}, B: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}}
	m.Borders[63] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -13.83333}}
	m.Borders[64] = Edge{A: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}}
	m.Borders[65] = Edge{A: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -10.16667}}
	m.Borders[66] = Edge{A: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}}
	m.Borders[67] = Edge{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -10.16667}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -9.5}}
	m.Borders[68] = Edge{A: PublicMatchState_Vector2Df{X: 0, Y: -15}, B: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}}
	m.Borders[69] = Edge{A: PublicMatchState_Vector2Df{X: 13.5, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: -3.166666}}
	m.Borders[70] = Edge{A: PublicMatchState_Vector2Df{X: 14.83333, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: -6.5}}
	m.Borders[71] = Edge{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -3.5}, B: PublicMatchState_Vector2Df{X: 7, Y: -3.166666}}
	m.Borders[72] = Edge{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -4.333332}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -3.5}}
	m.Borders[73] = Edge{A: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 13.5, Y: -3.166666}}
	m.Borders[74] = Edge{A: PublicMatchState_Vector2Df{X: 14.83333, Y: -6.5}, B: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}}
	m.Borders[75] = Edge{A: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -4.333332}}
	m.Borders[76] = Edge{A: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 17, Y: -6.5}}
	m.Borders[77] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: 0}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}}
	m.Borders[78] = Edge{A: PublicMatchState_Vector2Df{X: 17, Y: -6.5}, B: PublicMatchState_Vector2Df{X: 17, Y: 0}}
	m.Borders[79] = Edge{A: PublicMatchState_Vector2Df{X: 10.83333, Y: -2.666666}, B: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}}
	m.Borders[80] = Edge{A: PublicMatchState_Vector2Df{X: 7, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 6.333333, Y: -2.666666}}
	m.Borders[81] = Edge{A: PublicMatchState_Vector2Df{X: 3.166667, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}}
	m.Borders[82] = Edge{A: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}, B: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}}
	m.Borders[83] = Edge{A: PublicMatchState_Vector2Df{X: 5.666667, Y: -2.666666}, B: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}}
	m.Borders[84] = Edge{A: PublicMatchState_Vector2Df{X: -5.333332, Y: 0}, B: PublicMatchState_Vector2Df{X: -10.33333, Y: -8}}
	m.Borders[85] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 0}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}}
	m.Borders[86] = Edge{A: PublicMatchState_Vector2Df{X: -13.16667, Y: -14.5}, B: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}}
	m.Borders[87] = Edge{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -24.33333}}
	m.Borders[88] = Edge{A: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}, B: PublicMatchState_Vector2Df{X: 0, Y: -15}}
	m.Borders[89] = Edge{A: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -24.33333}}
	m.Borders[90] = Edge{A: PublicMatchState_Vector2Df{X: -9.166666, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}}
	m.Borders[91] = Edge{A: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}, B: PublicMatchState_Vector2Df{X: -13.16667, Y: -13}}
	m.Borders[92] = Edge{A: PublicMatchState_Vector2Df{X: -13.16667, Y: -13}, B: PublicMatchState_Vector2Df{X: -13.16667, Y: -14.5}}
	m.Borders[93] = Edge{A: PublicMatchState_Vector2Df{X: -10.33333, Y: -8.666666}, B: PublicMatchState_Vector2Df{X: -9.166666, Y: -9.5}}
	m.Borders[94] = Edge{A: PublicMatchState_Vector2Df{X: -10.33333, Y: -8}, B: PublicMatchState_Vector2Df{X: -10.33333, Y: -8.666666}}
	m.Borders[95] = Edge{A: PublicMatchState_Vector2Df{X: -0.5, Y: -12.66667}, B: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}}
	m.Borders[96] = Edge{A: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}, B: PublicMatchState_Vector2Df{X: -0.5, Y: -12.66667}}
	m.Borders[97] = Edge{A: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 0}}

	m.Triangles = make([]Triangle, 110)
	m.Triangles[0] = Triangle{A: PublicMatchState_Vector2Df{X: 4.833333, Y: 12}, B: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}, C: PublicMatchState_Vector2Df{X: 4.333333, Y: 12.5}, W: PublicMatchState_Vector2Df{X: 5.944444, Y: 12.33333}}
	m.Triangles[1] = Triangle{A: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, B: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}, C: PublicMatchState_Vector2Df{X: 4.833333, Y: 12}, W: PublicMatchState_Vector2Df{X: 6.222222, Y: 11}}
	m.Triangles[2] = Triangle{A: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: 5.666667}, C: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}, W: PublicMatchState_Vector2Df{X: 9.555554, Y: 8.888889}}
	m.Triangles[3] = Triangle{A: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, B: PublicMatchState_Vector2Df{X: 11, Y: 1.166667}, C: PublicMatchState_Vector2Df{X: 14.83333, Y: 5.666667}, W: PublicMatchState_Vector2Df{X: 10.33333, Y: 5.111111}}
	m.Triangles[4] = Triangle{A: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, B: PublicMatchState_Vector2Df{X: 10.66667, Y: 0.8333334}, C: PublicMatchState_Vector2Df{X: 11, Y: 1.166667}, W: PublicMatchState_Vector2Df{X: 8.944446, Y: 3.5}}
	m.Triangles[5] = Triangle{A: PublicMatchState_Vector2Df{X: 14.83333, Y: 5.666667}, B: PublicMatchState_Vector2Df{X: 11, Y: 1.166667}, C: PublicMatchState_Vector2Df{X: 14.83333, Y: 1.166667}, W: PublicMatchState_Vector2Df{X: 13.55555, Y: 2.666667}}
	m.Triangles[6] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, C: PublicMatchState_Vector2Df{X: 4.166667, Y: 8}, W: PublicMatchState_Vector2Df{X: 3.111111, Y: 5.5}}
	m.Triangles[7] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 10.66667, Y: 0.8333334}, C: PublicMatchState_Vector2Df{X: 5.166667, Y: 8.5}, W: PublicMatchState_Vector2Df{X: 5.277779, Y: 3.111111}}
	m.Triangles[8] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}, C: PublicMatchState_Vector2Df{X: 10.66667, Y: 0.8333334}, W: PublicMatchState_Vector2Df{X: 7.111113, Y: 0.2777778}}
	m.Triangles[9] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 4.166667, Y: 8}, C: PublicMatchState_Vector2Df{X: 0, Y: 7.5}, W: PublicMatchState_Vector2Df{X: 1.388889, Y: 5.166667}}
	m.Triangles[10] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 0}, B: PublicMatchState_Vector2Df{X: 17, Y: 6.5}, C: PublicMatchState_Vector2Df{X: 17, Y: 0}, W: PublicMatchState_Vector2Df{X: 19.5, Y: 2.166667}}
	m.Triangles[11] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 17, Y: 6.5}, C: PublicMatchState_Vector2Df{X: 24.5, Y: 0}, W: PublicMatchState_Vector2Df{X: 22, Y: 7.555557}}
	m.Triangles[12] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}, C: PublicMatchState_Vector2Df{X: 17, Y: 6.5}, W: PublicMatchState_Vector2Df{X: 16.61111, Y: 12.94445}}
	m.Triangles[13] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 24.5}, B: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}, C: PublicMatchState_Vector2Df{X: 24.5, Y: 16.16667}, W: PublicMatchState_Vector2Df{X: 19.11111, Y: 18.94445}}
	m.Triangles[14] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 24.5}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}, W: PublicMatchState_Vector2Df{X: 10.94444, Y: 21.72222}}
	m.Triangles[15] = Triangle{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 14.5}, B: PublicMatchState_Vector2Df{X: 4.333333, Y: 12.5}, C: PublicMatchState_Vector2Df{X: 8.666667, Y: 12.5}, W: PublicMatchState_Vector2Df{X: 6.611111, Y: 13.16667}}
	m.Triangles[16] = Triangle{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 14.5}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: 12.5}, C: PublicMatchState_Vector2Df{X: 4.333333, Y: 12.5}, W: PublicMatchState_Vector2Df{X: 5, Y: 13.16667}}
	m.Triangles[17] = Triangle{A: PublicMatchState_Vector2Df{X: 8.333334, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: 7.833333, Y: 16.16667}, W: PublicMatchState_Vector2Df{X: 5.388889, Y: 18.94445}}
	m.Triangles[18] = Triangle{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 15.33333}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: 12.5}, C: PublicMatchState_Vector2Df{X: 6.833333, Y: 14.5}, W: PublicMatchState_Vector2Df{X: 5.833333, Y: 14.11111}}
	m.Triangles[19] = Triangle{A: PublicMatchState_Vector2Df{X: 6.833333, Y: 15.33333}, B: PublicMatchState_Vector2Df{X: 0, Y: 12}, C: PublicMatchState_Vector2Df{X: 3.833333, Y: 12.5}, W: PublicMatchState_Vector2Df{X: 3.555555, Y: 13.27778}}
	m.Triangles[20] = Triangle{A: PublicMatchState_Vector2Df{X: 7.833333, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 0, Y: 12}, C: PublicMatchState_Vector2Df{X: 6.833333, Y: 15.33333}, W: PublicMatchState_Vector2Df{X: 4.888889, Y: 14.5}}
	m.Triangles[21] = Triangle{A: PublicMatchState_Vector2Df{X: 7.833333, Y: 16.16667}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: 0, Y: 12}, W: PublicMatchState_Vector2Df{X: 2.611111, Y: 17.55556}}
	m.Triangles[22] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -5.333332, Y: 0}, C: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, W: PublicMatchState_Vector2Df{X: -11.05555, Y: 1.888889}}
	m.Triangles[23] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 0}, C: PublicMatchState_Vector2Df{X: -5.333332, Y: 0}, W: PublicMatchState_Vector2Df{X: -18, Y: 0.9444444}}
	m.Triangles[24] = Triangle{A: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 24.5}, C: PublicMatchState_Vector2Df{X: -24.33333, Y: 18.83333}, W: PublicMatchState_Vector2Df{X: -19.11111, Y: 20.72222}}
	m.Triangles[25] = Triangle{A: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: -24.33333, Y: 24.5}, W: PublicMatchState_Vector2Df{X: -11, Y: 22.61111}}
	m.Triangles[26] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: -1.833332, Y: 1.333333}, C: PublicMatchState_Vector2Df{X: -2.666666, Y: 0}, W: PublicMatchState_Vector2Df{X: -1.499999, Y: 0.4444443}}
	m.Triangles[27] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, B: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}, C: PublicMatchState_Vector2Df{X: -7.166666, Y: 18.83333}, W: PublicMatchState_Vector2Df{X: -5.277777, Y: 20.72222}}
	m.Triangles[28] = Triangle{A: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, B: PublicMatchState_Vector2Df{X: 0, Y: 7.5}, C: PublicMatchState_Vector2Df{X: -0.666666, Y: 7.833333}, W: PublicMatchState_Vector2Df{X: -1.888889, Y: 7.333333}}
	m.Triangles[29] = Triangle{A: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, B: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}, C: PublicMatchState_Vector2Df{X: 0, Y: 7.5}, W: PublicMatchState_Vector2Df{X: -2.277777, Y: 5.444445}}
	m.Triangles[30] = Triangle{A: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 2.833333}, C: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}, W: PublicMatchState_Vector2Df{X: -3.166666, Y: 3.888889}}
	m.Triangles[31] = Triangle{A: PublicMatchState_Vector2Df{X: -1, Y: 11.5}, B: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}, C: PublicMatchState_Vector2Df{X: -1, Y: 10}, W: PublicMatchState_Vector2Df{X: -2.333333, Y: 9.611111}}
	m.Triangles[32] = Triangle{A: PublicMatchState_Vector2Df{X: -1, Y: 11.5}, B: PublicMatchState_Vector2Df{X: -6.833332, Y: 8.833334}, C: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}, W: PublicMatchState_Vector2Df{X: -4.277777, Y: 9.222222}}
	m.Triangles[33] = Triangle{A: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, C: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, W: PublicMatchState_Vector2Df{X: -12.33333, Y: 2.888889}}
	m.Triangles[34] = Triangle{A: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, C: PublicMatchState_Vector2Df{X: -24.33333, Y: 2.833333}, W: PublicMatchState_Vector2Df{X: -19.27777, Y: 2.944444}}
	m.Triangles[35] = Triangle{A: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}, B: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, C: PublicMatchState_Vector2Df{X: -8.5, Y: 3}, W: PublicMatchState_Vector2Df{X: -7.055555, Y: 2.944444}}
	m.Triangles[36] = Triangle{A: PublicMatchState_Vector2Df{X: -9, Y: 18.5}, B: PublicMatchState_Vector2Df{X: -11.66667, Y: 6}, C: PublicMatchState_Vector2Df{X: -9, Y: 8.833334}, W: PublicMatchState_Vector2Df{X: -9.88889, Y: 11.11111}}
	m.Triangles[37] = Triangle{A: PublicMatchState_Vector2Df{X: -9, Y: 18.5}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, C: PublicMatchState_Vector2Df{X: -11.66667, Y: 6}, W: PublicMatchState_Vector2Df{X: -15, Y: 9.166667}}
	m.Triangles[38] = Triangle{A: PublicMatchState_Vector2Df{X: -9, Y: 18.5}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 18.83333}, C: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, W: PublicMatchState_Vector2Df{X: -19.22222, Y: 13.44444}}
	m.Triangles[39] = Triangle{A: PublicMatchState_Vector2Df{X: -11.66667, Y: 5.333333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, C: PublicMatchState_Vector2Df{X: -9.166666, Y: 3}, W: PublicMatchState_Vector2Df{X: -15.05556, Y: 3.777778}}
	m.Triangles[40] = Triangle{A: PublicMatchState_Vector2Df{X: -11.66667, Y: 6}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: 3}, C: PublicMatchState_Vector2Df{X: -11.66667, Y: 5.333333}, W: PublicMatchState_Vector2Df{X: -15.88889, Y: 4.777778}}
	m.Triangles[41] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 18.83333}, B: PublicMatchState_Vector2Df{X: -9, Y: 18.5}, C: PublicMatchState_Vector2Df{X: -8.666666, Y: 18.83333}, W: PublicMatchState_Vector2Df{X: -14, Y: 18.72222}}
	m.Triangles[42] = Triangle{A: PublicMatchState_Vector2Df{X: -0.666666, Y: 7.833333}, B: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}, C: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, W: PublicMatchState_Vector2Df{X: -3.555555, Y: 7.277778}}
	m.Triangles[43] = Triangle{A: PublicMatchState_Vector2Df{X: -0.666666, Y: 7.833333}, B: PublicMatchState_Vector2Df{X: -1, Y: 10}, C: PublicMatchState_Vector2Df{X: -5, Y: 7.333333}, W: PublicMatchState_Vector2Df{X: -2.222222, Y: 8.388888}}
	m.Triangles[44] = Triangle{A: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, C: PublicMatchState_Vector2Df{X: -8.5, Y: 3}, W: PublicMatchState_Vector2Df{X: -5.666667, Y: 4.166667}}
	m.Triangles[45] = Triangle{A: PublicMatchState_Vector2Df{X: -3.5, Y: 2.833333}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 2.833333}, C: PublicMatchState_Vector2Df{X: -5, Y: 6.666667}, W: PublicMatchState_Vector2Df{X: -3.722222, Y: 4.111111}}
	m.Triangles[46] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}, C: PublicMatchState_Vector2Df{X: -1.833332, Y: 1.333333}, W: PublicMatchState_Vector2Df{X: -1.222221, Y: 1.166667}}
	m.Triangles[47] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 0, Y: 7.5}, C: PublicMatchState_Vector2Df{X: -1.833332, Y: 2.166667}, W: PublicMatchState_Vector2Df{X: -0.6111106, Y: 3.222222}}
	m.Triangles[48] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 12}, B: PublicMatchState_Vector2Df{X: -6.833332, Y: 8.833334}, C: PublicMatchState_Vector2Df{X: -1, Y: 11.5}, W: PublicMatchState_Vector2Df{X: -2.611111, Y: 10.77778}}
	m.Triangles[49] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 12}, B: PublicMatchState_Vector2Df{X: -6.833332, Y: 18.5}, C: PublicMatchState_Vector2Df{X: -6.833332, Y: 8.833334}, W: PublicMatchState_Vector2Df{X: -4.555555, Y: 13.11111}}
	m.Triangles[50] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 12}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: -6.833332, Y: 18.5}, W: PublicMatchState_Vector2Df{X: -2.277777, Y: 18.33333}}
	m.Triangles[51] = Triangle{A: PublicMatchState_Vector2Df{X: -6.833332, Y: 18.5}, B: PublicMatchState_Vector2Df{X: 0, Y: 24.5}, C: PublicMatchState_Vector2Df{X: -7.166666, Y: 18.83333}, W: PublicMatchState_Vector2Df{X: -4.666666, Y: 20.61111}}
	m.Triangles[52] = Triangle{A: PublicMatchState_Vector2Df{X: 19.83333, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, W: PublicMatchState_Vector2Df{X: 22.94444, Y: -21.55556}}
	m.Triangles[53] = Triangle{A: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 19.83333, Y: -20.16667}, W: PublicMatchState_Vector2Df{X: 21, Y: -21.55556}}
	m.Triangles[54] = Triangle{A: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}, C: PublicMatchState_Vector2Df{X: 17.33333, Y: -16}, W: PublicMatchState_Vector2Df{X: 13.27778, Y: -16.66667}}
	m.Triangles[55] = Triangle{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 3.166667, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 1.5, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 1.777778, Y: -14.22222}}
	m.Triangles[56] = Triangle{A: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}, B: PublicMatchState_Vector2Df{X: 5.666667, Y: -2.666666}, C: PublicMatchState_Vector2Df{X: 6.333333, Y: -2.666666}, W: PublicMatchState_Vector2Df{X: 7.555557, Y: -1.777777}}
	m.Triangles[57] = Triangle{A: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}, B: PublicMatchState_Vector2Df{X: 6.333333, Y: -2.666666}, C: PublicMatchState_Vector2Df{X: 10.83333, Y: -2.666666}, W: PublicMatchState_Vector2Df{X: 9.277778, Y: -1.777777}}
	m.Triangles[58] = Triangle{A: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 17.33333, Y: -16}, C: PublicMatchState_Vector2Df{X: 17.33333, Y: -14.83333}, W: PublicMatchState_Vector2Df{X: 12.83333, Y: -14.88889}}
	m.Triangles[59] = Triangle{A: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 17.33333, Y: -14.83333}, C: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 13.72222, Y: -14.16666}}
	m.Triangles[60] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 22.66667, Y: -19.16667}, C: PublicMatchState_Vector2Df{X: 19.83333, Y: -20.16667}, W: PublicMatchState_Vector2Df{X: 22.33333, Y: -19.83334}}
	m.Triangles[61] = Triangle{A: PublicMatchState_Vector2Df{X: 22, Y: -15.83333}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 22.55556, Y: -14.5}}
	m.Triangles[62] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 22, Y: -15.83333}, C: PublicMatchState_Vector2Df{X: 22.66667, Y: -19.16667}, W: PublicMatchState_Vector2Df{X: 23.05556, Y: -18.38889}}
	m.Triangles[63] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: -20.16667}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 22, Y: -15.83333}, W: PublicMatchState_Vector2Df{X: 23.66667, Y: -16.61111}}
	m.Triangles[64] = Triangle{A: PublicMatchState_Vector2Df{X: 1.5, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 2.333333, Y: -12.83333}, C: PublicMatchState_Vector2Df{X: 2, Y: -12.83333}, W: PublicMatchState_Vector2Df{X: 1.944444, Y: -13.16666}}
	m.Triangles[65] = Triangle{A: PublicMatchState_Vector2Df{X: 1.5, Y: -13.83333}, B: PublicMatchState_Vector2Df{X: 3.166667, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 2.333333, Y: -12.83333}, W: PublicMatchState_Vector2Df{X: 2.333333, Y: -13.5}}
	m.Triangles[66] = Triangle{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}, C: PublicMatchState_Vector2Df{X: 5.166667, Y: -7.666666}, W: PublicMatchState_Vector2Df{X: 6.111111, Y: -8}}
	m.Triangles[67] = Triangle{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, C: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 9.444446, Y: -7.722221}}
	m.Triangles[68] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 24.5, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 23.38889, Y: -11.5}}
	m.Triangles[69] = Triangle{A: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 17.27778, Y: -9.166665}}
	m.Triangles[70] = Triangle{A: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 13, Y: -11.5}}
	m.Triangles[71] = Triangle{A: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -10.16667}, C: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 8.777779, Y: -10.27778}}
	m.Triangles[72] = Triangle{A: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 20.77778, Y: -9.166665}}
	m.Triangles[73] = Triangle{A: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 20, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 21.16667, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 19.27778, Y: -11.5}}
	m.Triangles[74] = Triangle{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -10.16667}, B: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, C: PublicMatchState_Vector2Df{X: 7.333333, Y: -9.5}, W: PublicMatchState_Vector2Df{X: 9.944446, Y: -8.833334}}
	m.Triangles[75] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}, B: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, C: PublicMatchState_Vector2Df{X: 0, Y: -15}, W: PublicMatchState_Vector2Df{X: 0.2222222, Y: -18.11111}}
	m.Triangles[76] = Triangle{A: PublicMatchState_Vector2Df{X: 13.5, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: -6.5}, C: PublicMatchState_Vector2Df{X: 14.83333, Y: -3.166666}, W: PublicMatchState_Vector2Df{X: 14.38889, Y: -4.277777}}
	m.Triangles[77] = Triangle{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -3.5}, B: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, C: PublicMatchState_Vector2Df{X: 7, Y: -3.166666}, W: PublicMatchState_Vector2Df{X: 8.500001, Y: -3.277777}}
	m.Triangles[78] = Triangle{A: PublicMatchState_Vector2Df{X: 7.333333, Y: -4.333332}, B: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, C: PublicMatchState_Vector2Df{X: 7.333333, Y: -3.5}, W: PublicMatchState_Vector2Df{X: 8.611112, Y: -3.666666}}
	m.Triangles[79] = Triangle{A: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 14.83333, Y: -6.5}, C: PublicMatchState_Vector2Df{X: 13.5, Y: -3.166666}, W: PublicMatchState_Vector2Df{X: 13.16667, Y: -4.277777}}
	m.Triangles[80] = Triangle{A: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, C: PublicMatchState_Vector2Df{X: 14.83333, Y: -6.5}, W: PublicMatchState_Vector2Df{X: 13.72222, Y: -5.5}}
	m.Triangles[81] = Triangle{A: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}, C: PublicMatchState_Vector2Df{X: 15.16667, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 10.72222, Y: -5.61111}}
	m.Triangles[82] = Triangle{A: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, B: PublicMatchState_Vector2Df{X: 7.333333, Y: -4.333332}, C: PublicMatchState_Vector2Df{X: 5.833333, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 8.111112, Y: -4.777777}}
	m.Triangles[83] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}, B: PublicMatchState_Vector2Df{X: 17, Y: -6.5}, C: PublicMatchState_Vector2Df{X: 16.66667, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 19.38889, Y: -6.722221}}
	m.Triangles[84] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 0}, B: PublicMatchState_Vector2Df{X: 17, Y: -6.5}, C: PublicMatchState_Vector2Df{X: 24.5, Y: -6.833332}, W: PublicMatchState_Vector2Df{X: 22, Y: -4.444444}}
	m.Triangles[85] = Triangle{A: PublicMatchState_Vector2Df{X: 24.5, Y: 0}, B: PublicMatchState_Vector2Df{X: 17, Y: 0}, C: PublicMatchState_Vector2Df{X: 17, Y: -6.5}, W: PublicMatchState_Vector2Df{X: 19.5, Y: -2.166667}}
	m.Triangles[86] = Triangle{A: PublicMatchState_Vector2Df{X: 10.83333, Y: -2.666666}, B: PublicMatchState_Vector2Df{X: 7, Y: -3.166666}, C: PublicMatchState_Vector2Df{X: 11.16667, Y: -3.166666}, W: PublicMatchState_Vector2Df{X: 9.666667, Y: -2.999999}}
	m.Triangles[87] = Triangle{A: PublicMatchState_Vector2Df{X: 10.83333, Y: -2.666666}, B: PublicMatchState_Vector2Df{X: 6.333333, Y: -2.666666}, C: PublicMatchState_Vector2Df{X: 7, Y: -3.166666}, W: PublicMatchState_Vector2Df{X: 8.055554, Y: -2.833333}}
	m.Triangles[88] = Triangle{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, C: PublicMatchState_Vector2Df{X: 3.166667, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 2.555556, Y: -14.22222}}
	m.Triangles[89] = Triangle{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}, C: PublicMatchState_Vector2Df{X: 3.833333, Y: -13.83333}, W: PublicMatchState_Vector2Df{X: 7.722223, Y: -16.33333}}
	m.Triangles[90] = Triangle{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 18.66667, Y: -20.16667}, W: PublicMatchState_Vector2Df{X: 14.61111, Y: -19.83333}}
	m.Triangles[91] = Triangle{A: PublicMatchState_Vector2Df{X: 0.6666667, Y: -15}, B: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 24.5, Y: -24.33333}, W: PublicMatchState_Vector2Df{X: 8.388889, Y: -21.22222}}
	m.Triangles[92] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 5.666667, Y: -2.666666}, C: PublicMatchState_Vector2Df{X: 10.66667, Y: 0}, W: PublicMatchState_Vector2Df{X: 5.444446, Y: -0.8888887}}
	m.Triangles[93] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: 0}, B: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}, C: PublicMatchState_Vector2Df{X: 5.666667, Y: -2.666666}, W: PublicMatchState_Vector2Df{X: 1.888889, Y: -4.777779}}
	m.Triangles[94] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 0}, B: PublicMatchState_Vector2Df{X: -10.33333, Y: -8}, C: PublicMatchState_Vector2Df{X: -5.333332, Y: 0}, W: PublicMatchState_Vector2Df{X: -13.33333, Y: -2.666667}}
	m.Triangles[95] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: 0}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, C: PublicMatchState_Vector2Df{X: -10.33333, Y: -8}, W: PublicMatchState_Vector2Df{X: -19.66666, Y: -5.833333}}
	m.Triangles[96] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}, C: PublicMatchState_Vector2Df{X: -13.16667, Y: -14.5}, W: PublicMatchState_Vector2Df{X: -16.77778, Y: -12.94444}}
	m.Triangles[97] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}, W: PublicMatchState_Vector2Df{X: -20.5, Y: -16.22222}}
	m.Triangles[98] = Triangle{A: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}, B: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 0, Y: -15}, W: PublicMatchState_Vector2Df{X: -4.277777, Y: -18.05555}}
	m.Triangles[99] = Triangle{A: PublicMatchState_Vector2Df{X: -12.83333, Y: -14.83333}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -24.33333}, C: PublicMatchState_Vector2Df{X: 0, Y: -24.33333}, W: PublicMatchState_Vector2Df{X: -12.38889, Y: -21.16666}}
	m.Triangles[100] = Triangle{A: PublicMatchState_Vector2Df{X: -9.166666, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}, C: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}, W: PublicMatchState_Vector2Df{X: -10.16667, Y: -10.55556}}
	m.Triangles[101] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}, C: PublicMatchState_Vector2Df{X: -9.166666, Y: -9.5}, W: PublicMatchState_Vector2Df{X: -15.44444, Y: -10.55556}}
	m.Triangles[102] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -13.16667, Y: -13}, C: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}, W: PublicMatchState_Vector2Df{X: -16.77778, Y: -11.72222}}
	m.Triangles[103] = Triangle{A: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, B: PublicMatchState_Vector2Df{X: -13.16667, Y: -14.5}, C: PublicMatchState_Vector2Df{X: -13.16667, Y: -13}, W: PublicMatchState_Vector2Df{X: -16.88889, Y: -12.33333}}
	m.Triangles[104] = Triangle{A: PublicMatchState_Vector2Df{X: -10.33333, Y: -8.666666}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, C: PublicMatchState_Vector2Df{X: -9.166666, Y: -9.5}, W: PublicMatchState_Vector2Df{X: -14.61111, Y: -9.222222}}
	m.Triangles[105] = Triangle{A: PublicMatchState_Vector2Df{X: -10.33333, Y: -8}, B: PublicMatchState_Vector2Df{X: -24.33333, Y: -9.5}, C: PublicMatchState_Vector2Df{X: -10.33333, Y: -8.666666}, W: PublicMatchState_Vector2Df{X: -15, Y: -8.722222}}
	m.Triangles[106] = Triangle{A: PublicMatchState_Vector2Df{X: -0.5, Y: -12.66667}, B: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}, C: PublicMatchState_Vector2Df{X: -12.83333, Y: -12.66667}, W: PublicMatchState_Vector2Df{X: -7.277777, Y: -11.61111}}
	m.Triangles[107] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}, B: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}, C: PublicMatchState_Vector2Df{X: -0.5, Y: -12.66667}, W: PublicMatchState_Vector2Df{X: -3, Y: -11.27778}}
	m.Triangles[108] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}, B: PublicMatchState_Vector2Df{X: -2.666666, Y: 0}, C: PublicMatchState_Vector2Df{X: -8.5, Y: -9.5}, W: PublicMatchState_Vector2Df{X: -3.722222, Y: -7.055557}}
	m.Triangles[109] = Triangle{A: PublicMatchState_Vector2Df{X: 0, Y: -11.66667}, B: PublicMatchState_Vector2Df{X: 0, Y: 0}, C: PublicMatchState_Vector2Df{X: -2.666666, Y: 0}, W: PublicMatchState_Vector2Df{X: -0.8888887, Y: -3.88889}}

	return m
}
