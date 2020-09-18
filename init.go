package main

import (
	"Nakama-Go-Modules/GameDB"
	"Nakama-Go-Modules/graphmap"
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
		Range:       5,
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
		Range:       5,
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
		Range:       5,
		BlockValue:  0,
	}
	GameDatabase.Items[GameDB_Item_Staff.Id] = GameDB_Item_Staff
	GameDB_Item_Shield := &GameDB.Item{
		Id:          4,
		Name:        "Shield",
		Description: "Shield",
		Type:        GameDB.Item_Type_Weapon_Shield,
		Slot:        GameDB.Item_Slot_Weapon_OffHand,
		DamageMin:   0,
		DamageMax:   0,
		AttackSpeed: 0,
		Range:       0,
		BlockValue:  50,
	}
	GameDatabase.Items[GameDB_Item_Shield.Id] = GameDB_Item_Shield

	//EFFECTS
	GameDB_Effect_Fireball := &GameDB.Effect{
		Id:          1,
		Name:        "Fireball",
		Description: "Fireball",
		Visible:     true,
		EffectID:    0,
		IconID:      1,
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
		IconID:      2,
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
		IconID:      2,
		Duration:    5,
		Dispellable: true,
		School:      GameDB.Spell_SchoolType_Frost,
		Mechanic:    GameDB.Spell_Mechanic_Slowed,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Speed,
			Value: 0.5,
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
		IconID:      3,
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
		IconID:      3,
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
		IconID:      4,
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
		IconID:      5,
		Duration:    5,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Physical,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Speed,
			Value: 1.75,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_Haste.Id] = GameDB_Effect_Haste
	GameDB_Effect_HealingSpirits := &GameDB.Effect{
		Id:          8,
		Name:        "Healing Spirits",
		Description: "Healing Spirits",
		Visible:     true,
		EffectID:    0,
		IconID:      6,
		Duration:    6,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Holy,
		Mechanic:    GameDB.Spell_Mechanic_Healing,
		Type: &GameDB.Effect_Apply_Aura_Periodic_Heal{
			Intervall: 1,
		},
		ValueMin: 20,
		ValueMax: 30,
	}
	GameDatabase.Effects[GameDB_Effect_HealingSpirits.Id] = GameDB_Effect_HealingSpirits
	GameDB_Effect_HealingSurge := &GameDB.Effect{
		Id:          9,
		Name:        "Healing Surge",
		Description: "Healing Surge",
		Visible:     true,
		EffectID:    0,
		IconID:      7,
		Duration:    0,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Holy,
		Mechanic:    GameDB.Spell_Mechanic_Healing,
		Type:        &GameDB.Effect_Heal{},
		ValueMin:    50,
		ValueMax:    60,
	}
	GameDatabase.Effects[GameDB_Effect_HealingSurge.Id] = GameDB_Effect_HealingSurge
	GameDB_Effect_BlessedArmor := &GameDB.Effect{
		Id:          10,
		Name:        "Blessed Armor",
		Description: "Blessed Armor",
		Visible:     true,
		EffectID:    0,
		IconID:      8,
		Duration:    5,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Holy,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Armor,
			Value: 50,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_BlessedArmor.Id] = GameDB_Effect_BlessedArmor
	GameDB_Effect_ShieldWalled := &GameDB.Effect{
		Id:          11,
		Name:        "Shield Walled",
		Description: "Shield Walled",
		Visible:     true,
		EffectID:    0,
		IconID:      9,
		Duration:    5,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Physical,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Apply_Aura_Mod{
			Stat:  GameDB.Stat_Armor,
			Value: 75,
		},
		ValueMin: 0,
		ValueMax: 0,
	}
	GameDatabase.Effects[GameDB_Effect_ShieldWalled.Id] = GameDB_Effect_ShieldWalled
	GameDB_Effect_MightyStrike := &GameDB.Effect{
		Id:          12,
		Name:        "Mighty Strike",
		Description: "Mighty Strike",
		Visible:     false,
		EffectID:    0,
		IconID:      10,
		Duration:    5,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Physical,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type:        &GameDB.Effect_Autoattack{},
		ValueMin:    20,
		ValueMax:    30,
	}
	GameDatabase.Effects[GameDB_Effect_MightyStrike.Id] = GameDB_Effect_MightyStrike
	GameDB_Effect_FireZone := &GameDB.Effect{
		Id:          13,
		Name:        "Fire Zone",
		Description: "Fire Zone",
		Visible:     false,
		EffectID:    0,
		IconID:      11,
		Duration:    10,
		Dispellable: false,
		School:      GameDB.Spell_SchoolType_Fire,
		Mechanic:    GameDB.Spell_Mechanic_None,
		Type: &GameDB.Effect_Persistent_Area_Aura{
			Intervall: 1,
			Radius:    3,
		},
		ValueMin: 25,
		ValueMax: 30,
	}
	GameDatabase.Effects[GameDB_Effect_FireZone.Id] = GameDB_Effect_FireZone

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
		IconID:             1,
		Speed:              60,
		Application_Type:   GameDB.Spell_Application_Type_Missile,
		BaseCost:           15,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           2,
		CastTimeChanneled:  0,
		Range:              50,
		NeedLoS:            true,
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
		IconID:             2,
		Speed:              50,
		Application_Type:   GameDB.Spell_Application_Type_Missile,
		BaseCost:           10,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              20,
		NeedLoS:            true,
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
		IconID:             3,
		Speed:              1000,
		Application_Type:   GameDB.Spell_Application_Type_Beam,
		BaseCost:           20,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              300,
		NeedLoS:            true,
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
		IconID:             4,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           0,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              0,
		NeedLoS:            false,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Self,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[6], GameDatabase.Effects[7]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Enrage.Id] = GameDB_Spell_Enrage
	GameDB_Spell_HealingSpirits := &GameDB.Spell{
		Id:                 5,
		Name:               "Healing Spirits",
		Description:        "Healing Spirits",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           10,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             5,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           50,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              50,
		NeedLoS:            true,
		FacingFront:        true,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Ally,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[8]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_HealingSpirits.Id] = GameDB_Spell_HealingSpirits
	GameDB_Spell_HealingSurge := &GameDB.Spell{
		Id:                 6,
		Name:               "Healing Surge",
		Description:        "Healing Surge",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           0,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             6,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           20,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           2,
		CastTimeChanneled:  0,
		Range:              30,
		NeedLoS:            true,
		FacingFront:        true,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Ally,
		InterruptedBy:      GameDB.Interrupt_Type_OnMovement,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[9]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_HealingSurge.Id] = GameDB_Spell_HealingSurge
	GameDB_Spell_BlessingArmor := &GameDB.Spell{
		Id:                 7,
		Name:               "Blessing Armor",
		Description:        "Blessing Armor",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           30,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             7,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           15,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              50,
		NeedLoS:            true,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Ally,
		InterruptedBy:      GameDB.Interrupt_Type_OnMovement,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[10]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_BlessingArmor.Id] = GameDB_Spell_BlessingArmor
	GameDB_Spell_ShieldWall := &GameDB.Spell{
		Id:                 8,
		Name:               "Shield Wall",
		Description:        "Shield Wall",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           20,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             8,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Instant,
		BaseCost:           0,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              0,
		NeedLoS:            false,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Self,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[11]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_ShieldWall.Id] = GameDB_Spell_ShieldWall
	GameDB_Spell_MightyStrike := &GameDB.Spell{
		Id:                 9,
		Name:               "Mighty Strike",
		Description:        "Mighty Strike",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           5,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             9,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_WeaponSwing,
		BaseCost:           5,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              0,
		NeedLoS:            false,
		FacingFront:        true,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Enemy,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[12]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_MightyStrike.Id] = GameDB_Spell_MightyStrike
	GameDB_Spell_FireZone := &GameDB.Spell{
		Id:                 10,
		Name:               "Fire Zone",
		Description:        "Fire Zone",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           10,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             10,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_AoE,
		BaseCost:           0,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              60,
		NeedLoS:            false,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_AoE,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{GameDatabase.Effects[13]},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_FireZone.Id] = GameDB_Spell_FireZone
	GameDB_Spell_Teleport := &GameDB.Spell{
		Id:                 11,
		Name:               "Teleport",
		Description:        "Teleport",
		Visible:            true,
		ThreadModifier:     1,
		Cooldown:           0,
		GlobalCooldown:     1.5,
		IgnoresGCD:         false,
		IgnoresWeaponswing: false,
		MissileID:          0,
		EffectID:           0,
		IconID:             12,
		Speed:              0,
		Application_Type:   GameDB.Spell_Application_Type_Teleport,
		BaseCost:           0,
		CostPerSec:         0,
		CostPercentage:     0,
		CastTime:           0,
		CastTimeChanneled:  0,
		Range:              80,
		NeedLoS:            false,
		FacingFront:        false,
		TargetAuraRequired: 0,
		CasterAuraRequired: 0,
		Target_Type:        GameDB.Spell_Target_Type_Self,
		InterruptedBy:      GameDB.Interrupt_Type_None,
		ApplyEffect:        []*GameDB.Effect{},
		ApplyProc:          []*GameDB.Proc{},
	}
	GameDatabase.Spells[GameDB_Spell_Teleport.Id] = GameDB_Spell_Teleport

	//CLASSES
	GameDB_Class_Mage := &GameDB.Class{
		Name:                    "Mage",
		Description:             "Mage",
		Spells:                  []*GameDB.Spell{GameDatabase.Spells[1], GameDatabase.Spells[2], GameDatabase.Spells[3], GameDatabase.Spells[5], GameDatabase.Spells[6], GameDatabase.Spells[7], GameDatabase.Spells[10]},
		Mainhand:                GameDatabase.Items[3],
		Offhand:                 nil,
		Items:                   []*GameDB.Item{},
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
		MovementSpeed:           65,
	}
	GameDatabase.Classes[GameDB_Class_Mage.Name] = GameDB_Class_Mage
	GameDB_Class_Ogre := &GameDB.Class{
		Name:                    "Ogre",
		Description:             "Ogre",
		Spells:                  []*GameDB.Spell{GameDatabase.Spells[4], GameDatabase.Spells[10]},
		Mainhand:                GameDatabase.Items[2],
		Offhand:                 nil,
		Items:                   []*GameDB.Item{},
		Procs:                   []*GameDB.Proc{},
		BaseStamina:             20,
		GainStamina:             2,
		FactorHPRegen:           2,
		FactorArmor:             1.2,
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
		MovementSpeed:           45,
	}
	GameDatabase.Classes[GameDB_Class_Ogre.Name] = GameDB_Class_Ogre
	GameDB_Class_Warrior := &GameDB.Class{
		Name:                    "Warrior",
		Description:             "Warrior",
		Spells:                  []*GameDB.Spell{GameDatabase.Spells[8], GameDatabase.Spells[9]},
		Mainhand:                GameDatabase.Items[1],
		Offhand:                 GameDatabase.Items[4],
		Items:                   []*GameDB.Item{},
		Procs:                   []*GameDB.Proc{},
		BaseStamina:             20,
		GainStamina:             2,
		FactorHPRegen:           2,
		FactorArmor:             1.2,
		FactorSpellResist:       2,
		FactorBlock:             2,
		BaseStrength:            20,
		GainStrength:            2,
		FactorStrengthAP:        2,
		FactorParry:             2,
		BaseAgility:             15,
		GainAgility:             1,
		FactorAgilityAP:         1,
		FactorMeeleAttackSpeed:  1,
		FactorMeeleCriticalHits: 1.5,
		FactorDodge:             1.5,
		BaseIntellect:           1,
		GainIntellect:           0.5,
		FactorManaRegen:         0.3,
		FactorSpellAP:           0.1,
		FactorSpellAttackSpeed:  0.1,
		FactorSpellCriticalHits: 0.1,
		MovementSpeed:           70,
	}
	GameDatabase.Classes[GameDB_Class_Warrior.Name] = GameDB_Class_Warrior

	return GameDatabase
}

func map_init() *graphmap.Map {
	m := &graphmap.Map{}

	m.Borders = make([]graphmap.Edge, 98)
	m.Borders[0] = graphmap.Edge{A: graphmap.Vector{X: -5.333332, Y: 0}, B: graphmap.Vector{X: -10.33333, Y: -8}}
	m.Borders[1] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 0}, B: graphmap.Vector{X: -24.33333, Y: -9.5}}
	m.Borders[2] = graphmap.Edge{A: graphmap.Vector{X: -13.16667, Y: -14.5}, B: graphmap.Vector{X: -12.83333, Y: -14.83333}}
	m.Borders[3] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -24.33333, Y: -24.33333}}
	m.Borders[4] = graphmap.Edge{A: graphmap.Vector{X: -12.83333, Y: -14.83333}, B: graphmap.Vector{X: 0, Y: -15}}
	m.Borders[5] = graphmap.Edge{A: graphmap.Vector{X: 0, Y: -24.33333}, B: graphmap.Vector{X: -24.33333, Y: -24.33333}}
	m.Borders[6] = graphmap.Edge{A: graphmap.Vector{X: -9.166666, Y: -9.5}, B: graphmap.Vector{X: -8.5, Y: -9.5}}
	m.Borders[7] = graphmap.Edge{A: graphmap.Vector{X: -12.83333, Y: -12.66667}, B: graphmap.Vector{X: -13.16667, Y: -13}}
	m.Borders[8] = graphmap.Edge{A: graphmap.Vector{X: -13.16667, Y: -13}, B: graphmap.Vector{X: -13.16667, Y: -14.5}}
	m.Borders[9] = graphmap.Edge{A: graphmap.Vector{X: -10.33333, Y: -8.666666}, B: graphmap.Vector{X: -9.166666, Y: -9.5}}
	m.Borders[10] = graphmap.Edge{A: graphmap.Vector{X: -10.33333, Y: -8}, B: graphmap.Vector{X: -10.33333, Y: -8.666666}}
	m.Borders[11] = graphmap.Edge{A: graphmap.Vector{X: -0.5, Y: -12.66667}, B: graphmap.Vector{X: -12.83333, Y: -12.66667}}
	m.Borders[12] = graphmap.Edge{A: graphmap.Vector{X: 0, Y: -11.66667}, B: graphmap.Vector{X: -0.5, Y: -12.66667}}
	m.Borders[13] = graphmap.Edge{A: graphmap.Vector{X: -8.5, Y: -9.5}, B: graphmap.Vector{X: -2.666666, Y: 0}}
	m.Borders[14] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: -20.16667}, B: graphmap.Vector{X: 24.5, Y: -24.33333}}
	m.Borders[15] = graphmap.Edge{A: graphmap.Vector{X: 18.66667, Y: -20.16667}, B: graphmap.Vector{X: 19.83333, Y: -20.16667}}
	m.Borders[16] = graphmap.Edge{A: graphmap.Vector{X: 17.33333, Y: -16}, B: graphmap.Vector{X: 18.66667, Y: -20.16667}}
	m.Borders[17] = graphmap.Edge{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 1.5, Y: -13.83333}}
	m.Borders[18] = graphmap.Edge{A: graphmap.Vector{X: 6.333333, Y: -2.666666}, B: graphmap.Vector{X: 5.666667, Y: -2.666666}}
	m.Borders[19] = graphmap.Edge{A: graphmap.Vector{X: 10.66667, Y: 0}, B: graphmap.Vector{X: 10.83333, Y: -2.666666}}
	m.Borders[20] = graphmap.Edge{A: graphmap.Vector{X: 17.33333, Y: -14.83333}, B: graphmap.Vector{X: 17.33333, Y: -16}}
	m.Borders[21] = graphmap.Edge{A: graphmap.Vector{X: 20, Y: -13.83333}, B: graphmap.Vector{X: 17.33333, Y: -14.83333}}
	m.Borders[22] = graphmap.Edge{A: graphmap.Vector{X: 19.83333, Y: -20.16667}, B: graphmap.Vector{X: 22.66667, Y: -19.16667}}
	m.Borders[23] = graphmap.Edge{A: graphmap.Vector{X: 22, Y: -15.83333}, B: graphmap.Vector{X: 21.16667, Y: -13.83333}}
	m.Borders[24] = graphmap.Edge{A: graphmap.Vector{X: 22.66667, Y: -19.16667}, B: graphmap.Vector{X: 22, Y: -15.83333}}
	m.Borders[25] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: -20.16667}, B: graphmap.Vector{X: 24.5, Y: -13.83333}}
	m.Borders[26] = graphmap.Edge{A: graphmap.Vector{X: 1.5, Y: -13.83333}, B: graphmap.Vector{X: 2, Y: -12.83333}}
	m.Borders[27] = graphmap.Edge{A: graphmap.Vector{X: 2, Y: -12.83333}, B: graphmap.Vector{X: 2.333333, Y: -12.83333}}
	m.Borders[28] = graphmap.Edge{A: graphmap.Vector{X: 2.333333, Y: -12.83333}, B: graphmap.Vector{X: 3.166667, Y: -13.83333}}
	m.Borders[29] = graphmap.Edge{A: graphmap.Vector{X: 7.333333, Y: -9.5}, B: graphmap.Vector{X: 5.166667, Y: -7.666666}}
	m.Borders[30] = graphmap.Edge{A: graphmap.Vector{X: 5.166667, Y: -7.666666}, B: graphmap.Vector{X: 5.833333, Y: -6.833332}}
	m.Borders[31] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: -6.833332}, B: graphmap.Vector{X: 24.5, Y: -13.83333}}
	m.Borders[32] = graphmap.Edge{A: graphmap.Vector{X: 15.16667, Y: -6.833332}, B: graphmap.Vector{X: 16.66667, Y: -6.833332}}
	m.Borders[33] = graphmap.Edge{A: graphmap.Vector{X: 3.833333, Y: -13.83333}, B: graphmap.Vector{X: 7.333333, Y: -10.16667}}
	m.Borders[34] = graphmap.Edge{A: graphmap.Vector{X: 21.16667, Y: -13.83333}, B: graphmap.Vector{X: 20, Y: -13.83333}}
	m.Borders[35] = graphmap.Edge{A: graphmap.Vector{X: 7.333333, Y: -10.16667}, B: graphmap.Vector{X: 7.333333, Y: -9.5}}
	m.Borders[36] = graphmap.Edge{A: graphmap.Vector{X: 0, Y: -15}, B: graphmap.Vector{X: 0.6666667, Y: -15}}
	m.Borders[37] = graphmap.Edge{A: graphmap.Vector{X: 13.5, Y: -3.166666}, B: graphmap.Vector{X: 14.83333, Y: -3.166666}}
	m.Borders[38] = graphmap.Edge{A: graphmap.Vector{X: 14.83333, Y: -3.166666}, B: graphmap.Vector{X: 14.83333, Y: -6.5}}
	m.Borders[39] = graphmap.Edge{A: graphmap.Vector{X: 7.333333, Y: -3.5}, B: graphmap.Vector{X: 7, Y: -3.166666}}
	m.Borders[40] = graphmap.Edge{A: graphmap.Vector{X: 7.333333, Y: -4.333332}, B: graphmap.Vector{X: 7.333333, Y: -3.5}}
	m.Borders[41] = graphmap.Edge{A: graphmap.Vector{X: 11.16667, Y: -3.166666}, B: graphmap.Vector{X: 13.5, Y: -3.166666}}
	m.Borders[42] = graphmap.Edge{A: graphmap.Vector{X: 14.83333, Y: -6.5}, B: graphmap.Vector{X: 15.16667, Y: -6.833332}}
	m.Borders[43] = graphmap.Edge{A: graphmap.Vector{X: 5.833333, Y: -6.833332}, B: graphmap.Vector{X: 7.333333, Y: -4.333332}}
	m.Borders[44] = graphmap.Edge{A: graphmap.Vector{X: 16.66667, Y: -6.833332}, B: graphmap.Vector{X: 17, Y: -6.5}}
	m.Borders[45] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: 0}, B: graphmap.Vector{X: 24.5, Y: -6.833332}}
	m.Borders[46] = graphmap.Edge{A: graphmap.Vector{X: 17, Y: -6.5}, B: graphmap.Vector{X: 17, Y: 0}}
	m.Borders[47] = graphmap.Edge{A: graphmap.Vector{X: 10.83333, Y: -2.666666}, B: graphmap.Vector{X: 11.16667, Y: -3.166666}}
	m.Borders[48] = graphmap.Edge{A: graphmap.Vector{X: 7, Y: -3.166666}, B: graphmap.Vector{X: 6.333333, Y: -2.666666}}
	m.Borders[49] = graphmap.Edge{A: graphmap.Vector{X: 3.166667, Y: -13.83333}, B: graphmap.Vector{X: 3.833333, Y: -13.83333}}
	m.Borders[50] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: -24.33333}, B: graphmap.Vector{X: 0, Y: -24.33333}}
	m.Borders[51] = graphmap.Edge{A: graphmap.Vector{X: 5.666667, Y: -2.666666}, B: graphmap.Vector{X: 0, Y: -11.66667}}
	m.Borders[52] = graphmap.Edge{A: graphmap.Vector{X: -3.5, Y: 2.833333}, B: graphmap.Vector{X: -5.333332, Y: 0}}
	m.Borders[53] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 2.833333}, B: graphmap.Vector{X: -24.33333, Y: 0}}
	m.Borders[54] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 18.83333}, B: graphmap.Vector{X: -24.33333, Y: 24.5}}
	m.Borders[55] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 24.5}, B: graphmap.Vector{X: 0, Y: 24.5}}
	m.Borders[56] = graphmap.Edge{A: graphmap.Vector{X: -2.666666, Y: 0}, B: graphmap.Vector{X: -1.833332, Y: 1.333333}}
	m.Borders[57] = graphmap.Edge{A: graphmap.Vector{X: -7.166666, Y: 18.83333}, B: graphmap.Vector{X: -8.666666, Y: 18.83333}}
	m.Borders[58] = graphmap.Edge{A: graphmap.Vector{X: -0.666666, Y: 7.833333}, B: graphmap.Vector{X: 0, Y: 7.5}}
	m.Borders[59] = graphmap.Edge{A: graphmap.Vector{X: -1.833332, Y: 2.166667}, B: graphmap.Vector{X: -2.666666, Y: 2.833333}}
	m.Borders[60] = graphmap.Edge{A: graphmap.Vector{X: -1, Y: 11.5}, B: graphmap.Vector{X: -1, Y: 10}}
	m.Borders[61] = graphmap.Edge{A: graphmap.Vector{X: -5, Y: 7.333333}, B: graphmap.Vector{X: -6.833332, Y: 8.833334}}
	m.Borders[62] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 2.833333}, B: graphmap.Vector{X: -24.33333, Y: 3}}
	m.Borders[63] = graphmap.Edge{A: graphmap.Vector{X: -9.166666, Y: 3}, B: graphmap.Vector{X: -8.5, Y: 3}}
	m.Borders[64] = graphmap.Edge{A: graphmap.Vector{X: -9, Y: 18.5}, B: graphmap.Vector{X: -9, Y: 8.833334}}
	m.Borders[65] = graphmap.Edge{A: graphmap.Vector{X: -9, Y: 8.833334}, B: graphmap.Vector{X: -11.66667, Y: 6}}
	m.Borders[66] = graphmap.Edge{A: graphmap.Vector{X: -24.33333, Y: 3}, B: graphmap.Vector{X: -24.33333, Y: 18.83333}}
	m.Borders[67] = graphmap.Edge{A: graphmap.Vector{X: -11.66667, Y: 5.333333}, B: graphmap.Vector{X: -9.166666, Y: 3}}
	m.Borders[68] = graphmap.Edge{A: graphmap.Vector{X: -11.66667, Y: 6}, B: graphmap.Vector{X: -11.66667, Y: 5.333333}}
	m.Borders[69] = graphmap.Edge{A: graphmap.Vector{X: -8.666666, Y: 18.83333}, B: graphmap.Vector{X: -9, Y: 18.5}}
	m.Borders[70] = graphmap.Edge{A: graphmap.Vector{X: -5, Y: 6.666667}, B: graphmap.Vector{X: -5, Y: 7.333333}}
	m.Borders[71] = graphmap.Edge{A: graphmap.Vector{X: -0.666666, Y: 7.833333}, B: graphmap.Vector{X: -1, Y: 10}}
	m.Borders[72] = graphmap.Edge{A: graphmap.Vector{X: -8.5, Y: 3}, B: graphmap.Vector{X: -5, Y: 6.666667}}
	m.Borders[73] = graphmap.Edge{A: graphmap.Vector{X: -3.5, Y: 2.833333}, B: graphmap.Vector{X: -2.666666, Y: 2.833333}}
	m.Borders[74] = graphmap.Edge{A: graphmap.Vector{X: -1.833332, Y: 1.333333}, B: graphmap.Vector{X: -1.833332, Y: 2.166667}}
	m.Borders[75] = graphmap.Edge{A: graphmap.Vector{X: 0, Y: 12}, B: graphmap.Vector{X: -1, Y: 11.5}}
	m.Borders[76] = graphmap.Edge{A: graphmap.Vector{X: -6.833332, Y: 8.833334}, B: graphmap.Vector{X: -6.833332, Y: 18.5}}
	m.Borders[77] = graphmap.Edge{A: graphmap.Vector{X: -6.833332, Y: 18.5}, B: graphmap.Vector{X: -7.166666, Y: 18.83333}}
	m.Borders[78] = graphmap.Edge{A: graphmap.Vector{X: 4.833333, Y: 12}, B: graphmap.Vector{X: 4.333333, Y: 12.5}}
	m.Borders[79] = graphmap.Edge{A: graphmap.Vector{X: 5.166667, Y: 8.5}, B: graphmap.Vector{X: 4.833333, Y: 12}}
	m.Borders[80] = graphmap.Edge{A: graphmap.Vector{X: 8.666667, Y: 12.5}, B: graphmap.Vector{X: 14.83333, Y: 5.666667}}
	m.Borders[81] = graphmap.Edge{A: graphmap.Vector{X: 11, Y: 1.166667}, B: graphmap.Vector{X: 10.66667, Y: 0.8333334}}
	m.Borders[82] = graphmap.Edge{A: graphmap.Vector{X: 14.83333, Y: 5.666667}, B: graphmap.Vector{X: 14.83333, Y: 1.166667}}
	m.Borders[83] = graphmap.Edge{A: graphmap.Vector{X: 14.83333, Y: 1.166667}, B: graphmap.Vector{X: 11, Y: 1.166667}}
	m.Borders[84] = graphmap.Edge{A: graphmap.Vector{X: 4.166667, Y: 8}, B: graphmap.Vector{X: 5.166667, Y: 8.5}}
	m.Borders[85] = graphmap.Edge{A: graphmap.Vector{X: 10.66667, Y: 0.8333334}, B: graphmap.Vector{X: 10.66667, Y: 0}}
	m.Borders[86] = graphmap.Edge{A: graphmap.Vector{X: 0, Y: 7.5}, B: graphmap.Vector{X: 4.166667, Y: 8}}
	m.Borders[87] = graphmap.Edge{A: graphmap.Vector{X: 17, Y: 0}, B: graphmap.Vector{X: 17, Y: 6.5}}
	m.Borders[88] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: 16.16667}, B: graphmap.Vector{X: 24.5, Y: 0}}
	m.Borders[89] = graphmap.Edge{A: graphmap.Vector{X: 17, Y: 6.5}, B: graphmap.Vector{X: 8.333334, Y: 16.16667}}
	m.Borders[90] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: 24.5}, B: graphmap.Vector{X: 24.5, Y: 16.16667}}
	m.Borders[91] = graphmap.Edge{A: graphmap.Vector{X: 24.5, Y: 24.5}, B: graphmap.Vector{X: 0, Y: 24.5}}
	m.Borders[92] = graphmap.Edge{A: graphmap.Vector{X: 6.833333, Y: 14.5}, B: graphmap.Vector{X: 8.666667, Y: 12.5}}
	m.Borders[93] = graphmap.Edge{A: graphmap.Vector{X: 4.333333, Y: 12.5}, B: graphmap.Vector{X: 3.833333, Y: 12.5}}
	m.Borders[94] = graphmap.Edge{A: graphmap.Vector{X: 8.333334, Y: 16.16667}, B: graphmap.Vector{X: 7.833333, Y: 16.16667}}
	m.Borders[95] = graphmap.Edge{A: graphmap.Vector{X: 6.833333, Y: 15.33333}, B: graphmap.Vector{X: 6.833333, Y: 14.5}}
	m.Borders[96] = graphmap.Edge{A: graphmap.Vector{X: 3.833333, Y: 12.5}, B: graphmap.Vector{X: 0, Y: 12}}
	m.Borders[97] = graphmap.Edge{A: graphmap.Vector{X: 7.833333, Y: 16.16667}, B: graphmap.Vector{X: 6.833333, Y: 15.33333}}

	m.Triangles = make([]graphmap.Triangle, 110)
	m.Triangles[0] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: 0}, B: graphmap.Vector{X: -10.33333, Y: -8}, C: graphmap.Vector{X: -5.333332, Y: 0}, W: graphmap.Vector{X: -13.33333, Y: -2.666667}, Neighbors: map[int32]float32{1: 8.14112, 59: 14.72736}}
	m.Triangles[1] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: 0}, B: graphmap.Vector{X: -24.33333, Y: -9.5}, C: graphmap.Vector{X: -10.33333, Y: -8}, W: graphmap.Vector{X: -19.66667, Y: -5.833333}, Neighbors: map[int32]float32{0: 17.29242, 11: 11.31930}}
	m.Triangles[2] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -12.83333, Y: -14.83333}, C: graphmap.Vector{X: -13.16667, Y: -14.5}, W: graphmap.Vector{X: -16.77778, Y: -12.94444}, Neighbors: map[int32]float32{3: 3.32453, 9: 4.48626}}
	m.Triangles[3] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -24.33333, Y: -24.33333}, C: graphmap.Vector{X: -12.83333, Y: -14.83333}, W: graphmap.Vector{X: -20.5, Y: -16.22222}, Neighbors: map[int32]float32{2: 8.23591, 5: 4.98919}}
	m.Triangles[4] = graphmap.Triangle{A: graphmap.Vector{X: -12.83333, Y: -14.83333}, B: graphmap.Vector{X: 0, Y: -24.33333}, C: graphmap.Vector{X: 0, Y: -15}, W: graphmap.Vector{X: -4.277778, Y: -18.05556}, Neighbors: map[int32]float32{5: 17.17305, 39: 13.83344}}
	m.Triangles[5] = graphmap.Triangle{A: graphmap.Vector{X: -12.83333, Y: -14.83333}, B: graphmap.Vector{X: -24.33333, Y: -24.33333}, C: graphmap.Vector{X: 0, Y: -24.33333}, W: graphmap.Vector{X: -12.38889, Y: -21.16667}, Neighbors: map[int32]float32{3: 6.25636, 4: 6.46453}}
	m.Triangles[6] = graphmap.Triangle{A: graphmap.Vector{X: -9.166666, Y: -9.5}, B: graphmap.Vector{X: -12.83333, Y: -12.66667}, C: graphmap.Vector{X: -8.5, Y: -9.5}, W: graphmap.Vector{X: -10.16667, Y: -10.55556}, Neighbors: map[int32]float32{7: 0.38889, 12: 1.78903}}
	m.Triangles[7] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -12.83333, Y: -12.66667}, C: graphmap.Vector{X: -9.166666, Y: -9.5}, W: graphmap.Vector{X: -15.44444, Y: -10.55556}, Neighbors: map[int32]float32{6: 4.88889, 8: 3.90078, 10: 6.36348}}
	m.Triangles[8] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -13.16667, Y: -13}, C: graphmap.Vector{X: -12.83333, Y: -12.66667}, W: graphmap.Vector{X: -16.77778, Y: -11.72222}, Neighbors: map[int32]float32{7: 6.33065, 9: 4.48626}}
	m.Triangles[9] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: -9.5}, B: graphmap.Vector{X: -13.16667, Y: -14.5}, C: graphmap.Vector{X: -13.16667, Y: -13}, W: graphmap.Vector{X: -16.88889, Y: -12.33333}, Neighbors: map[int32]float32{2: 3.99150, 8: 5.20268}}
	m.Triangles[10] = graphmap.Triangle{A: graphmap.Vector{X: -10.33333, Y: -8.666666}, B: graphmap.Vector{X: -24.33333, Y: -9.5}, C: graphmap.Vector{X: -9.166666, Y: -9.5}, W: graphmap.Vector{X: -14.61111, Y: -9.222222}, Neighbors: map[int32]float32{7: 4.26911, 11: 5.91008}}
	m.Triangles[11] = graphmap.Triangle{A: graphmap.Vector{X: -10.33333, Y: -8}, B: graphmap.Vector{X: -24.33333, Y: -9.5}, C: graphmap.Vector{X: -10.33333, Y: -8.666666}, W: graphmap.Vector{X: -15, Y: -8.722222}, Neighbors: map[int32]float32{1: 9.61111, 10: 5.79937}}
	m.Triangles[12] = graphmap.Triangle{A: graphmap.Vector{X: -0.5, Y: -12.66667}, B: graphmap.Vector{X: -8.5, Y: -9.5}, C: graphmap.Vector{X: -12.83333, Y: -12.66667}, W: graphmap.Vector{X: -7.277778, Y: -11.61111}, Neighbors: map[int32]float32{6: 3.44355, 13: 4.01387}}
	m.Triangles[13] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: -11.66667}, B: graphmap.Vector{X: -8.5, Y: -9.5}, C: graphmap.Vector{X: -0.5, Y: -12.66667}, W: graphmap.Vector{X: -3, Y: -11.27778}, Neighbors: map[int32]float32{12: 8.61756, 14: 5.85446}}
	m.Triangles[14] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: -11.66667}, B: graphmap.Vector{X: -2.666666, Y: 0}, C: graphmap.Vector{X: -8.5, Y: -9.5}, W: graphmap.Vector{X: -3.722222, Y: -7.055555}, Neighbors: map[int32]float32{13: 8.65526, 15: 3.17105}}
	m.Triangles[15] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: -11.66667}, B: graphmap.Vector{X: 0, Y: 0}, C: graphmap.Vector{X: -2.666666, Y: 0}, W: graphmap.Vector{X: -0.8888887, Y: -3.888889}, Neighbors: map[int32]float32{14: 6.93221, 57: 3.98918, 62: 4.53382}}
	m.Triangles[16] = graphmap.Triangle{A: graphmap.Vector{X: 19.83333, Y: -20.16667}, B: graphmap.Vector{X: 24.5, Y: -24.33333}, C: graphmap.Vector{X: 24.5, Y: -20.16667}, W: graphmap.Vector{X: 22.94444, Y: -21.55556}, Neighbors: map[int32]float32{17: 44.50000, 24: 42.81243}}
	m.Triangles[17] = graphmap.Triangle{A: graphmap.Vector{X: 18.66667, Y: -20.16667}, B: graphmap.Vector{X: 24.5, Y: -24.33333}, C: graphmap.Vector{X: 19.83333, Y: -20.16667}, W: graphmap.Vector{X: 21, Y: -21.55556}, Neighbors: map[int32]float32{16: 42.55556, 54: 40.86964}}
	m.Triangles[18] = graphmap.Triangle{A: graphmap.Vector{X: 3.833333, Y: -13.83333}, B: graphmap.Vector{X: 18.66667, Y: -20.16667}, C: graphmap.Vector{X: 17.33333, Y: -16}, W: graphmap.Vector{X: 13.27778, Y: -16.66667}, Neighbors: map[int32]float32{22: 28.22272, 53: 29.61299}}
	m.Triangles[19] = graphmap.Triangle{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 3.166667, Y: -13.83333}, C: graphmap.Vector{X: 1.5, Y: -13.83333}, W: graphmap.Vector{X: 1.777778, Y: -14.22222}, Neighbors: map[int32]float32{29: 15.29484, 52: 16.00000}}
	m.Triangles[20] = graphmap.Triangle{A: graphmap.Vector{X: 10.66667, Y: 0}, B: graphmap.Vector{X: 5.666667, Y: -2.666666}, C: graphmap.Vector{X: 6.333333, Y: -2.666666}, W: graphmap.Vector{X: 7.555556, Y: -1.777777}, Neighbors: map[int32]float32{21: 9.33333, 56: 8.49110}}
	m.Triangles[21] = graphmap.Triangle{A: graphmap.Vector{X: 10.66667, Y: 0}, B: graphmap.Vector{X: 6.333333, Y: -2.666666}, C: graphmap.Vector{X: 10.83333, Y: -2.666666}, W: graphmap.Vector{X: 9.277779, Y: -1.777777}, Neighbors: map[int32]float32{20: 11.05556, 51: 12.15702}}
	m.Triangles[22] = graphmap.Triangle{A: graphmap.Vector{X: 3.833333, Y: -13.83333}, B: graphmap.Vector{X: 17.33333, Y: -16}, C: graphmap.Vector{X: 17.33333, Y: -14.83333}, W: graphmap.Vector{X: 12.83333, Y: -14.88889}, Neighbors: map[int32]float32{18: 29.55352, 23: 27.00966}}
	m.Triangles[23] = graphmap.Triangle{A: graphmap.Vector{X: 3.833333, Y: -13.83333}, B: graphmap.Vector{X: 17.33333, Y: -14.83333}, C: graphmap.Vector{X: 20, Y: -13.83333}, W: graphmap.Vector{X: 13.72222, Y: -14.16667}, Neighbors: map[int32]float32{22: 28.62022, 34: 25.36280}}
	m.Triangles[24] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: -20.16667}, B: graphmap.Vector{X: 22.66667, Y: -19.16667}, C: graphmap.Vector{X: 19.83333, Y: -20.16667}, W: graphmap.Vector{X: 22.33333, Y: -19.83333}, Neighbors: map[int32]float32{16: 43.92267, 26: 40.74783}}
	m.Triangles[25] = graphmap.Triangle{A: graphmap.Vector{X: 22, Y: -15.83333}, B: graphmap.Vector{X: 24.5, Y: -13.83333}, C: graphmap.Vector{X: 21.16667, Y: -13.83333}, W: graphmap.Vector{X: 22.55556, Y: -14.5}, Neighbors: map[int32]float32{27: 39.22352, 32: 34.18744}}
	m.Triangles[26] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: -20.16667}, B: graphmap.Vector{X: 22, Y: -15.83333}, C: graphmap.Vector{X: 22.66667, Y: -19.16667}, W: graphmap.Vector{X: 23.05556, Y: -18.38889}, Neighbors: map[int32]float32{24: 42.91320, 27: 39.70649}}
	m.Triangles[27] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: -20.16667}, B: graphmap.Vector{X: 24.5, Y: -13.83333}, C: graphmap.Vector{X: 22, Y: -15.83333}, W: graphmap.Vector{X: 23.66667, Y: -16.61111}, Neighbors: map[int32]float32{25: 38.22501, 26: 42.09311}}
	m.Triangles[28] = graphmap.Triangle{A: graphmap.Vector{X: 1.5, Y: -13.83333}, B: graphmap.Vector{X: 2.333333, Y: -12.83333}, C: graphmap.Vector{X: 2, Y: -12.83333}, W: graphmap.Vector{X: 1.944445, Y: -13.16667}, Neighbors: map[int32]float32{29: 15.44804}}
	m.Triangles[29] = graphmap.Triangle{A: graphmap.Vector{X: 1.5, Y: -13.83333}, B: graphmap.Vector{X: 3.166667, Y: -13.83333}, C: graphmap.Vector{X: 2.333333, Y: -12.83333}, W: graphmap.Vector{X: 2.333333, Y: -13.5}, Neighbors: map[int32]float32{19: 16.57130, 28: 15.50358}}
	m.Triangles[30] = graphmap.Triangle{A: graphmap.Vector{X: 7.333333, Y: -9.5}, B: graphmap.Vector{X: 5.833333, Y: -6.833332}, C: graphmap.Vector{X: 5.166667, Y: -7.666666}, W: graphmap.Vector{X: 6.111111, Y: -8}, Neighbors: map[int32]float32{31: 13.83612}}
	m.Triangles[31] = graphmap.Triangle{A: graphmap.Vector{X: 7.333333, Y: -9.5}, B: graphmap.Vector{X: 15.16667, Y: -6.833332}, C: graphmap.Vector{X: 5.833333, Y: -6.833332}, W: graphmap.Vector{X: 9.444445, Y: -7.722221}, Neighbors: map[int32]float32{30: 17.44666, 38: 18.31152, 45: 15.20285}}
	m.Triangles[32] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: -6.833332}, B: graphmap.Vector{X: 21.16667, Y: -13.83333}, C: graphmap.Vector{X: 24.5, Y: -13.83333}, W: graphmap.Vector{X: 23.38889, Y: -11.5}, Neighbors: map[int32]float32{25: 38.00747, 36: 32.63906}}
	m.Triangles[33] = graphmap.Triangle{A: graphmap.Vector{X: 15.16667, Y: -6.833332}, B: graphmap.Vector{X: 20, Y: -13.83333}, C: graphmap.Vector{X: 16.66667, Y: -6.833332}, W: graphmap.Vector{X: 17.27778, Y: -9.166666}, Neighbors: map[int32]float32{34: 28.87222, 37: 28.87222}}
	m.Triangles[34] = graphmap.Triangle{A: graphmap.Vector{X: 15.16667, Y: -6.833332}, B: graphmap.Vector{X: 3.833333, Y: -13.83333}, C: graphmap.Vector{X: 20, Y: -13.83333}, W: graphmap.Vector{X: 13, Y: -11.5}, Neighbors: map[int32]float32{23: 27.29723, 33: 22.28913, 35: 23.30984}}
	m.Triangles[35] = graphmap.Triangle{A: graphmap.Vector{X: 15.16667, Y: -6.833332}, B: graphmap.Vector{X: 7.333333, Y: -10.16667}, C: graphmap.Vector{X: 3.833333, Y: -13.83333}, W: graphmap.Vector{X: 8.777778, Y: -10.27778}, Neighbors: map[int32]float32{34: 20.31458, 38: 17.67025}}
	m.Triangles[36] = graphmap.Triangle{A: graphmap.Vector{X: 16.66667, Y: -6.833332}, B: graphmap.Vector{X: 21.16667, Y: -13.83333}, C: graphmap.Vector{X: 24.5, Y: -6.833332}, W: graphmap.Vector{X: 20.77778, Y: -9.166666}, Neighbors: map[int32]float32{32: 32.36200, 37: 32.36200, 47: 27.60843}}
	m.Triangles[37] = graphmap.Triangle{A: graphmap.Vector{X: 16.66667, Y: -6.833332}, B: graphmap.Vector{X: 20, Y: -13.83333}, C: graphmap.Vector{X: 21.16667, Y: -13.83333}, W: graphmap.Vector{X: 19.27778, Y: -11.5}, Neighbors: map[int32]float32{33: 28.53999, 36: 28.53999}}
	m.Triangles[38] = graphmap.Triangle{A: graphmap.Vector{X: 7.333333, Y: -10.16667}, B: graphmap.Vector{X: 15.16667, Y: -6.833332}, C: graphmap.Vector{X: 7.333333, Y: -9.5}, W: graphmap.Vector{X: 9.944445, Y: -8.833333}, Neighbors: map[int32]float32{31: 17.70157, 35: 20.27374}}
	m.Triangles[39] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: -24.33333}, B: graphmap.Vector{X: 0.6666667, Y: -15}, C: graphmap.Vector{X: 0, Y: -15}, W: graphmap.Vector{X: 0.2222222, Y: -18.11111}, Neighbors: map[int32]float32{4: 18.27786, 55: 21.66895}}
	m.Triangles[40] = graphmap.Triangle{A: graphmap.Vector{X: 13.5, Y: -3.166666}, B: graphmap.Vector{X: 14.83333, Y: -6.5}, C: graphmap.Vector{X: 14.83333, Y: -3.166666}, W: graphmap.Vector{X: 14.38889, Y: -4.277777}, Neighbors: map[int32]float32{43: 18.66667}}
	m.Triangles[41] = graphmap.Triangle{A: graphmap.Vector{X: 7.333333, Y: -3.5}, B: graphmap.Vector{X: 11.16667, Y: -3.166666}, C: graphmap.Vector{X: 7, Y: -3.166666}, W: graphmap.Vector{X: 8.5, Y: -3.277777}, Neighbors: map[int32]float32{42: 12.17288, 50: 11.50335}}
	m.Triangles[42] = graphmap.Triangle{A: graphmap.Vector{X: 7.333333, Y: -4.333332}, B: graphmap.Vector{X: 11.16667, Y: -3.166666}, C: graphmap.Vector{X: 7.333333, Y: -3.5}, W: graphmap.Vector{X: 8.611112, Y: -3.666666}, Neighbors: map[int32]float32{41: 11.89525, 46: 13.43491}}
	m.Triangles[43] = graphmap.Triangle{A: graphmap.Vector{X: 11.16667, Y: -3.166666}, B: graphmap.Vector{X: 14.83333, Y: -6.5}, C: graphmap.Vector{X: 13.5, Y: -3.166666}, W: graphmap.Vector{X: 13.16667, Y: -4.277777}, Neighbors: map[int32]float32{40: 17.44444, 44: 18.70664}}
	m.Triangles[44] = graphmap.Triangle{A: graphmap.Vector{X: 11.16667, Y: -3.166666}, B: graphmap.Vector{X: 15.16667, Y: -6.833332}, C: graphmap.Vector{X: 14.83333, Y: -6.5}, W: graphmap.Vector{X: 13.72222, Y: -5.5}, Neighbors: map[int32]float32{43: 18.04145, 45: 19.33365}}
	m.Triangles[45] = graphmap.Triangle{A: graphmap.Vector{X: 11.16667, Y: -3.166666}, B: graphmap.Vector{X: 5.833333, Y: -6.833332}, C: graphmap.Vector{X: 15.16667, Y: -6.833332}, W: graphmap.Vector{X: 10.72222, Y: -5.61111}, Neighbors: map[int32]float32{31: 18.56487, 44: 16.22260, 46: 15.52238}}
	m.Triangles[46] = graphmap.Triangle{A: graphmap.Vector{X: 11.16667, Y: -3.166666}, B: graphmap.Vector{X: 7.333333, Y: -4.333332}, C: graphmap.Vector{X: 5.833333, Y: -6.833332}, W: graphmap.Vector{X: 8.111112, Y: -4.777777}, Neighbors: map[int32]float32{42: 11.83007, 45: 13.74750}}
	m.Triangles[47] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: -6.833332}, B: graphmap.Vector{X: 17, Y: -6.5}, C: graphmap.Vector{X: 16.66667, Y: -6.833332}, W: graphmap.Vector{X: 19.38889, Y: -6.722221}, Neighbors: map[int32]float32{36: 28.65999, 48: 23.94193}}
	m.Triangles[48] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 0}, B: graphmap.Vector{X: 17, Y: -6.5}, C: graphmap.Vector{X: 24.5, Y: -6.833332}, W: graphmap.Vector{X: 22, Y: -4.444444}, Neighbors: map[int32]float32{47: 28.81240, 49: 24.27377}}
	m.Triangles[49] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 0}, B: graphmap.Vector{X: 17, Y: 0}, C: graphmap.Vector{X: 17, Y: -6.5}, W: graphmap.Vector{X: 19.5, Y: -2.166667}, Neighbors: map[int32]float32{48: 24.05254, 98: 17.86679}}
	m.Triangles[50] = graphmap.Triangle{A: graphmap.Vector{X: 10.83333, Y: -2.666666}, B: graphmap.Vector{X: 7, Y: -3.166666}, C: graphmap.Vector{X: 11.16667, Y: -3.166666}, W: graphmap.Vector{X: 9.666667, Y: -2.999999}, Neighbors: map[int32]float32{41: 12.94742, 51: 12.50111}}
	m.Triangles[51] = graphmap.Triangle{A: graphmap.Vector{X: 10.83333, Y: -2.666666}, B: graphmap.Vector{X: 6.333333, Y: -2.666666}, C: graphmap.Vector{X: 7, Y: -3.166666}, W: graphmap.Vector{X: 8.055555, Y: -2.833333}, Neighbors: map[int32]float32{21: 9.88983, 50: 11.05681}}
	m.Triangles[52] = graphmap.Triangle{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 3.833333, Y: -13.83333}, C: graphmap.Vector{X: 3.166667, Y: -13.83333}, W: graphmap.Vector{X: 2.555556, Y: -14.22222}, Neighbors: map[int32]float32{19: 16.77778, 53: 19.00649}}
	m.Triangles[53] = graphmap.Triangle{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 18.66667, Y: -20.16667}, C: graphmap.Vector{X: 3.833333, Y: -13.83333}, W: graphmap.Vector{X: 7.722223, Y: -16.33333}, Neighbors: map[int32]float32{18: 24.39117, 52: 22.04576, 54: 27.77695}}
	m.Triangles[54] = graphmap.Triangle{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 24.5, Y: -24.33333}, C: graphmap.Vector{X: 18.66667, Y: -20.16667}, W: graphmap.Vector{X: 14.61111, Y: -19.83333}, Neighbors: map[int32]float32{17: 36.20765, 53: 31.14175, 55: 35.86024}}
	m.Triangles[55] = graphmap.Triangle{A: graphmap.Vector{X: 0.6666667, Y: -15}, B: graphmap.Vector{X: 0, Y: -24.33333}, C: graphmap.Vector{X: 24.5, Y: -24.33333}, W: graphmap.Vector{X: 8.388889, Y: -21.22222}, Neighbors: map[int32]float32{39: 26.68200, 54: 28.25638}}
	m.Triangles[56] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 5.666667, Y: -2.666666}, C: graphmap.Vector{X: 10.66667, Y: 0}, W: graphmap.Vector{X: 5.444445, Y: -0.8888887}, Neighbors: map[int32]float32{20: 7.27672, 57: 10.93697, 96: 5.29675}}
	m.Triangles[57] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 0, Y: -11.66667}, C: graphmap.Vector{X: 5.666667, Y: -2.666666}, W: graphmap.Vector{X: 1.888889, Y: -4.777777}, Neighbors: map[int32]float32{15: 5.84575, 56: 4.77907}}
	m.Triangles[58] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: 2.833333}, B: graphmap.Vector{X: -5.333332, Y: 0}, C: graphmap.Vector{X: -3.5, Y: 2.833333}, W: graphmap.Vector{X: -11.05556, Y: 1.888889}, Neighbors: map[int32]float32{59: 12.03711, 69: 13.98026}}
	m.Triangles[59] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: 2.833333}, B: graphmap.Vector{X: -24.33333, Y: 0}, C: graphmap.Vector{X: -5.333332, Y: 0}, W: graphmap.Vector{X: -18, Y: 0.9444445}, Neighbors: map[int32]float32{0: 15.75282, 58: 19.91130}}
	m.Triangles[60] = graphmap.Triangle{A: graphmap.Vector{X: -8.666666, Y: 18.83333}, B: graphmap.Vector{X: -24.33333, Y: 24.5}, C: graphmap.Vector{X: -24.33333, Y: 18.83333}, W: graphmap.Vector{X: -19.11111, Y: 20.72222}, Neighbors: map[int32]float32{61: 41.76496, 77: 37.88616}}
	m.Triangles[61] = graphmap.Triangle{A: graphmap.Vector{X: -8.666666, Y: 18.83333}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: -24.33333, Y: 24.5}, W: graphmap.Vector{X: -11, Y: 22.61111}, Neighbors: map[int32]float32{60: 31.77841, 63: 31.77841}}
	m.Triangles[62] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: -1.833332, Y: 1.333333}, C: graphmap.Vector{X: -2.666666, Y: 0}, W: graphmap.Vector{X: -1.499999, Y: 0.4444444}, Neighbors: map[int32]float32{15: 4.94819, 82: 2.76274}}
	m.Triangles[63] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 24.5}, B: graphmap.Vector{X: -8.666666, Y: 18.83333}, C: graphmap.Vector{X: -7.166666, Y: 18.83333}, W: graphmap.Vector{X: -5.277777, Y: 20.72222}, Neighbors: map[int32]float32{61: 27.95278, 87: 25.88913}}
	m.Triangles[64] = graphmap.Triangle{A: graphmap.Vector{X: -5, Y: 6.666667}, B: graphmap.Vector{X: 0, Y: 7.5}, C: graphmap.Vector{X: -0.666666, Y: 7.833333}, W: graphmap.Vector{X: -1.888889, Y: 7.333333}, Neighbors: map[int32]float32{65: 7.57269, 78: 9.16684}}
	m.Triangles[65] = graphmap.Triangle{A: graphmap.Vector{X: -5, Y: 6.666667}, B: graphmap.Vector{X: -1.833332, Y: 2.166667}, C: graphmap.Vector{X: 0, Y: 7.5}, W: graphmap.Vector{X: -2.277777, Y: 5.444445}, Neighbors: map[int32]float32{64: 9.79497, 66: 6.35984, 83: 5.93197}}
	m.Triangles[66] = graphmap.Triangle{A: graphmap.Vector{X: -5, Y: 6.666667}, B: graphmap.Vector{X: -2.666666, Y: 2.833333}, C: graphmap.Vector{X: -1.833332, Y: 2.166667}, W: graphmap.Vector{X: -3.166666, Y: 3.888889}, Neighbors: map[int32]float32{65: 8.75048, 81: 7.28117}}
	m.Triangles[67] = graphmap.Triangle{A: graphmap.Vector{X: -1, Y: 11.5}, B: graphmap.Vector{X: -5, Y: 7.333333}, C: graphmap.Vector{X: -1, Y: 10}, W: graphmap.Vector{X: -2.333333, Y: 9.611111}, Neighbors: map[int32]float32{68: 11.56210, 79: 10.79166}}
	m.Triangles[68] = graphmap.Triangle{A: graphmap.Vector{X: -1, Y: 11.5}, B: graphmap.Vector{X: -6.833332, Y: 8.833334}, C: graphmap.Vector{X: -5, Y: 7.333333}, W: graphmap.Vector{X: -4.277777, Y: 9.222222}, Neighbors: map[int32]float32{67: 13.89433, 84: 15.13570}}
	m.Triangles[69] = graphmap.Triangle{A: graphmap.Vector{X: -9.166666, Y: 3}, B: graphmap.Vector{X: -24.33333, Y: 2.833333}, C: graphmap.Vector{X: -3.5, Y: 2.833333}, W: graphmap.Vector{X: -12.33333, Y: 2.888889}, Neighbors: map[int32]float32{58: 14.25733, 70: 15.27788, 71: 15.27788}}
	m.Triangles[70] = graphmap.Triangle{A: graphmap.Vector{X: -9.166666, Y: 3}, B: graphmap.Vector{X: -24.33333, Y: 3}, C: graphmap.Vector{X: -24.33333, Y: 2.833333}, W: graphmap.Vector{X: -19.27778, Y: 2.944444}, Neighbors: map[int32]float32{69: 22.16674, 75: 23.07061}}
	m.Triangles[71] = graphmap.Triangle{A: graphmap.Vector{X: -9.166666, Y: 3}, B: graphmap.Vector{X: -3.5, Y: 2.833333}, C: graphmap.Vector{X: -8.5, Y: 3}, W: graphmap.Vector{X: -7.055555, Y: 2.944444}, Neighbors: map[int32]float32{69: 9.94460, 80: 11.28858}}
	m.Triangles[72] = graphmap.Triangle{A: graphmap.Vector{X: -9, Y: 18.5}, B: graphmap.Vector{X: -11.66667, Y: 6}, C: graphmap.Vector{X: -9, Y: 8.833334}, W: graphmap.Vector{X: -9.888888, Y: 11.11111}, Neighbors: map[int32]float32{73: 19.15450}}
	m.Triangles[73] = graphmap.Triangle{A: graphmap.Vector{X: -9, Y: 18.5}, B: graphmap.Vector{X: -24.33333, Y: 3}, C: graphmap.Vector{X: -11.66667, Y: 6}, W: graphmap.Vector{X: -15, Y: 9.166667}, Neighbors: map[int32]float32{72: 26.18341, 74: 28.76431, 76: 20.25890}}
	m.Triangles[74] = graphmap.Triangle{A: graphmap.Vector{X: -9, Y: 18.5}, B: graphmap.Vector{X: -24.33333, Y: 18.83333}, C: graphmap.Vector{X: -24.33333, Y: 3}, W: graphmap.Vector{X: -19.22222, Y: 13.44444}, Neighbors: map[int32]float32{73: 28.70938, 77: 38.30974}}
	m.Triangles[75] = graphmap.Triangle{A: graphmap.Vector{X: -11.66667, Y: 5.333333}, B: graphmap.Vector{X: -24.33333, Y: 3}, C: graphmap.Vector{X: -9.166666, Y: 3}, W: graphmap.Vector{X: -15.05556, Y: 3.777778}, Neighbors: map[int32]float32{70: 18.01928, 76: 19.85853}}
	m.Triangles[76] = graphmap.Triangle{A: graphmap.Vector{X: -11.66667, Y: 6}, B: graphmap.Vector{X: -24.33333, Y: 3}, C: graphmap.Vector{X: -11.66667, Y: 5.333333}, W: graphmap.Vector{X: -15.88889, Y: 4.777778}, Neighbors: map[int32]float32{73: 25.43704, 75: 19.69207}}
	m.Triangles[77] = graphmap.Triangle{A: graphmap.Vector{X: -24.33333, Y: 18.83333}, B: graphmap.Vector{X: -9, Y: 18.5}, C: graphmap.Vector{X: -8.666666, Y: 18.83333}, W: graphmap.Vector{X: -14, Y: 18.72222}, Neighbors: map[int32]float32{60: 34.77977, 74: 27.94732}}
	m.Triangles[78] = graphmap.Triangle{A: graphmap.Vector{X: -0.666666, Y: 7.833333}, B: graphmap.Vector{X: -5, Y: 7.333333}, C: graphmap.Vector{X: -5, Y: 6.666667}, W: graphmap.Vector{X: -3.555555, Y: 7.277778}, Neighbors: map[int32]float32{64: 10.88903, 79: 11.99601}}
	m.Triangles[79] = graphmap.Triangle{A: graphmap.Vector{X: -0.666666, Y: 7.833333}, B: graphmap.Vector{X: -1, Y: 10}, C: graphmap.Vector{X: -5, Y: 7.333333}, W: graphmap.Vector{X: -2.222222, Y: 8.388889}, Neighbors: map[int32]float32{67: 11.89629, 78: 9.56476}}
	m.Triangles[80] = graphmap.Triangle{A: graphmap.Vector{X: -3.5, Y: 2.833333}, B: graphmap.Vector{X: -5, Y: 6.666667}, C: graphmap.Vector{X: -8.5, Y: 3}, W: graphmap.Vector{X: -5.666667, Y: 4.166667}, Neighbors: map[int32]float32{71: 8.69742, 81: 9.77794}}
	m.Triangles[81] = graphmap.Triangle{A: graphmap.Vector{X: -3.5, Y: 2.833333}, B: graphmap.Vector{X: -2.666666, Y: 2.833333}, C: graphmap.Vector{X: -5, Y: 6.666667}, W: graphmap.Vector{X: -3.722222, Y: 4.111111}, Neighbors: map[int32]float32{66: 7.61436, 80: 7.88909}}
	m.Triangles[82] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: -1.833332, Y: 2.166667}, C: graphmap.Vector{X: -1.833332, Y: 1.333333}, W: graphmap.Vector{X: -1.222221, Y: 1.166667}, Neighbors: map[int32]float32{62: 1.81642, 83: 4.89677}}
	m.Triangles[83] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 0, Y: 7.5}, C: graphmap.Vector{X: -1.833332, Y: 2.166667}, W: graphmap.Vector{X: -0.6111107, Y: 3.222222}, Neighbors: map[int32]float32{65: 6.45043, 82: 2.71768, 97: 6.09619}}
	m.Triangles[84] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 12}, B: graphmap.Vector{X: -6.833332, Y: 8.833334}, C: graphmap.Vector{X: -1, Y: 11.5}, W: graphmap.Vector{X: -2.611111, Y: 10.77778}, Neighbors: map[int32]float32{68: 11.93514, 85: 15.89442}}
	m.Triangles[85] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 12}, B: graphmap.Vector{X: -6.833332, Y: 18.5}, C: graphmap.Vector{X: -6.833332, Y: 8.833334}, W: graphmap.Vector{X: -4.555555, Y: 13.11111}, Neighbors: map[int32]float32{84: 15.50985, 86: 23.47707}}
	m.Triangles[86] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 12}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: -6.833332, Y: 18.5}, W: graphmap.Vector{X: -2.277777, Y: 18.33333}, Neighbors: map[int32]float32{85: 16.25083, 87: 23.00195, 109: 19.84858}}
	m.Triangles[87] = graphmap.Triangle{A: graphmap.Vector{X: -6.833332, Y: 18.5}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: -7.166666, Y: 18.83333}, W: graphmap.Vector{X: -4.666666, Y: 20.61111}, Neighbors: map[int32]float32{63: 25.38913, 86: 23.11251}}
	m.Triangles[88] = graphmap.Triangle{A: graphmap.Vector{X: 4.833333, Y: 12}, B: graphmap.Vector{X: 8.666667, Y: 12.5}, C: graphmap.Vector{X: 4.333333, Y: 12.5}, W: graphmap.Vector{X: 5.944445, Y: 12.33333}, Neighbors: map[int32]float32{89: 5.22842, 103: 7.27014}}
	m.Triangles[89] = graphmap.Triangle{A: graphmap.Vector{X: 5.166667, Y: 8.5}, B: graphmap.Vector{X: 8.666667, Y: 12.5}, C: graphmap.Vector{X: 4.833333, Y: 12}, W: graphmap.Vector{X: 6.222222, Y: 11}, Neighbors: map[int32]float32{88: 6.25487, 90: 3.40116}}
	m.Triangles[90] = graphmap.Triangle{A: graphmap.Vector{X: 5.166667, Y: 8.5}, B: graphmap.Vector{X: 14.83333, Y: 5.666667}, C: graphmap.Vector{X: 8.666667, Y: 12.5}, W: graphmap.Vector{X: 9.555556, Y: 8.888889}, Neighbors: map[int32]float32{89: 2.55797, 91: 5.83307}}
	m.Triangles[91] = graphmap.Triangle{A: graphmap.Vector{X: 5.166667, Y: 8.5}, B: graphmap.Vector{X: 11, Y: 1.166667}, C: graphmap.Vector{X: 14.83333, Y: 5.666667}, W: graphmap.Vector{X: 10.33333, Y: 5.111111}, Neighbors: map[int32]float32{90: 4.04451, 92: 7.02069, 93: 8.04693}}
	m.Triangles[92] = graphmap.Triangle{A: graphmap.Vector{X: 5.166667, Y: 8.5}, B: graphmap.Vector{X: 10.66667, Y: 0.8333334}, C: graphmap.Vector{X: 11, Y: 1.166667}, W: graphmap.Vector{X: 8.944445, Y: 3.5}, Neighbors: map[int32]float32{91: 4.15814, 95: 5.84628}}
	m.Triangles[93] = graphmap.Triangle{A: graphmap.Vector{X: 14.83333, Y: 5.666667}, B: graphmap.Vector{X: 11, Y: 1.166667}, C: graphmap.Vector{X: 14.83333, Y: 1.166667}, W: graphmap.Vector{X: 13.55556, Y: 2.666667}, Neighbors: map[int32]float32{91: 8.79113}}
	m.Triangles[94] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 5.166667, Y: 8.5}, C: graphmap.Vector{X: 4.166667, Y: 8}, W: graphmap.Vector{X: 3.111111, Y: 5.5}, Neighbors: map[int32]float32{95: 2.38889, 97: 2.08241}}
	m.Triangles[95] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 10.66667, Y: 0.8333334}, C: graphmap.Vector{X: 5.166667, Y: 8.5}, W: graphmap.Vector{X: 5.277778, Y: 3.111111}, Neighbors: map[int32]float32{92: 1.81982, 94: 2.39920, 96: 5.74698}}
	m.Triangles[96] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 10.66667, Y: 0}, C: graphmap.Vector{X: 10.66667, Y: 0.8333334}, W: graphmap.Vector{X: 7.111111, Y: 0.2777778}, Neighbors: map[int32]float32{56: 8.08462, 95: 4.90181}}
	m.Triangles[97] = graphmap.Triangle{A: graphmap.Vector{X: 0, Y: 0}, B: graphmap.Vector{X: 4.166667, Y: 8}, C: graphmap.Vector{X: 0, Y: 7.5}, W: graphmap.Vector{X: 1.388889, Y: 5.166667}, Neighbors: map[int32]float32{83: 2.67245, 94: 4.12460}}
	m.Triangles[98] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 0}, B: graphmap.Vector{X: 17, Y: 6.5}, C: graphmap.Vector{X: 17, Y: 0}, W: graphmap.Vector{X: 19.5, Y: 2.166667}, Neighbors: map[int32]float32{49: 22.09575, 99: 13.10381}}
	m.Triangles[99] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 16.16667}, B: graphmap.Vector{X: 17, Y: 6.5}, C: graphmap.Vector{X: 24.5, Y: 0}, W: graphmap.Vector{X: 22, Y: 7.555556}, Neighbors: map[int32]float32{98: 20.55240, 100: 10.53770}}
	m.Triangles[100] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 16.16667}, B: graphmap.Vector{X: 8.333334, Y: 16.16667}, C: graphmap.Vector{X: 17, Y: 6.5}, W: graphmap.Vector{X: 16.61111, Y: 12.94445}, Neighbors: map[int32]float32{99: 10.53770, 101: 6.43774}}
	m.Triangles[101] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 24.5}, B: graphmap.Vector{X: 8.333334, Y: 16.16667}, C: graphmap.Vector{X: 24.5, Y: 16.16667}, W: graphmap.Vector{X: 19.11111, Y: 18.94444}, Neighbors: map[int32]float32{100: 8.60394, 102: 3.81234}}
	m.Triangles[102] = graphmap.Triangle{A: graphmap.Vector{X: 24.5, Y: 24.5}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: 8.333334, Y: 16.16667}, W: graphmap.Vector{X: 10.94444, Y: 21.72222}, Neighbors: map[int32]float32{101: 8.46853, 105: 8.46853}}
	m.Triangles[103] = graphmap.Triangle{A: graphmap.Vector{X: 6.833333, Y: 14.5}, B: graphmap.Vector{X: 4.333333, Y: 12.5}, C: graphmap.Vector{X: 8.666667, Y: 12.5}, W: graphmap.Vector{X: 6.611111, Y: 13.16667}, Neighbors: map[int32]float32{88: 5.78258, 104: 6.55556}}
	m.Triangles[104] = graphmap.Triangle{A: graphmap.Vector{X: 6.833333, Y: 14.5}, B: graphmap.Vector{X: 3.833333, Y: 12.5}, C: graphmap.Vector{X: 4.333333, Y: 12.5}, W: graphmap.Vector{X: 5, Y: 13.16667}, Neighbors: map[int32]float32{103: 8.16667, 106: 9.15993}}
	m.Triangles[105] = graphmap.Triangle{A: graphmap.Vector{X: 8.333334, Y: 16.16667}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: 7.833333, Y: 16.16667}, W: graphmap.Vector{X: 5.388889, Y: 18.94444}, Neighbors: map[int32]float32{102: 16.56786, 109: 12.24568}}
	m.Triangles[106] = graphmap.Triangle{A: graphmap.Vector{X: 6.833333, Y: 15.33333}, B: graphmap.Vector{X: 3.833333, Y: 12.5}, C: graphmap.Vector{X: 6.833333, Y: 14.5}, W: graphmap.Vector{X: 5.833333, Y: 14.11111}, Neighbors: map[int32]float32{104: 7.39390, 107: 7.49094}}
	m.Triangles[107] = graphmap.Triangle{A: graphmap.Vector{X: 6.833333, Y: 15.33333}, B: graphmap.Vector{X: 0, Y: 12}, C: graphmap.Vector{X: 3.833333, Y: 12.5}, W: graphmap.Vector{X: 3.555556, Y: 13.27778}, Neighbors: map[int32]float32{106: 10.58840, 108: 11.01248}}
	m.Triangles[108] = graphmap.Triangle{A: graphmap.Vector{X: 7.833333, Y: 16.16667}, B: graphmap.Vector{X: 0, Y: 12}, C: graphmap.Vector{X: 6.833333, Y: 15.33333}, W: graphmap.Vector{X: 4.888889, Y: 14.5}, Neighbors: map[int32]float32{107: 8.47746, 109: 13.03000}}
	m.Triangles[109] = graphmap.Triangle{A: graphmap.Vector{X: 7.833333, Y: 16.16667}, B: graphmap.Vector{X: 0, Y: 24.5}, C: graphmap.Vector{X: 0, Y: 12}, W: graphmap.Vector{X: 2.611111, Y: 17.55556}, Neighbors: map[int32]float32{86: 15.74145, 105: 16.39228, 108: 12.27526}}

	return m
}
