package main

func init_db() *GameDB {

	GameDB := &GameDB {
		Spells: make(map[int64]*GameDB_Spell),
		Effects: make(map[int64]*GameDB_Effect),
		Procs: make(map[int64]*GameDB_Proc),
		Spellbook: make([]*GameDB_Spell, 0),
	}
		
	fireball_dmg := &GameDB_Effect{
		Id				:1,
		Name 			:"Fireball",
		Description 	:"Fireball",
		Visible			:true,
		Dispellable		:false,
		Duration 		:0,
		EffectID		:0,
		Type 			: &GameDB_Effect_Damage {
			Type: GameDB_Spell_DamageType_Fire,
			ValueMin: 20,
			ValueMax: 30,
		},
	}
	GameDB.Effects[1] = fireball_dmg
	
	frostbolt_dmg := &GameDB_Effect{
		Id				:2,
		Name 			:"Frostbolt",
		Description 	:"Frostbolt",
		Visible			:true,
		Dispellable		:false,
		Duration 		:0,
		EffectID		:0,
		Type 			: &GameDB_Effect_Damage {
			Type: GameDB_Spell_DamageType_Frost,
			ValueMin: 10,
			ValueMax: 15,
		},
	}
	GameDB.Effects[2] = frostbolt_dmg

	chilled := &GameDB_Effect{
		Id				:3,
		Name 			:"Chilled",
		Description 	:"Chilled",
		Visible			:true,
		Dispellable		:true,
		Duration 		:5,
		EffectID		:1,
		Type 			: &GameDB_Effect_Apply_Aura_Mod {
			Stat: GameDB_Stat_Speed,
			Value: 15,
		},
	}
	GameDB.Effects[3] = chilled

	sunburn_dmg := &GameDB_Effect{
		Id				:4,
		Name 			:"Sunburn",
		Description 	:"Sunburn Initial Damage",
		Visible			:true,
		Dispellable		:false,
		Duration 		:0,
		EffectID		:0,
		Type 			: &GameDB_Effect_Damage {
			Type: GameDB_Spell_DamageType_Fire,
			ValueMin: 5,
			ValueMax: 10,
		},
	}
	GameDB.Effects[4] = sunburn_dmg

	sunburn_dot := &GameDB_Effect{
		Id				:5,
		Name 			:"Sunburn",
		Description 	:"Sunburn DoT",
		Visible			:true,
		Dispellable		:false,
		Duration 		:10,
		EffectID		:2,
		Type 			: &GameDB_Effect_Apply_Aura_Periodic_Damage {
			Type: GameDB_Spell_DamageType_Fire,
			ValueMin: 3,
			ValueMax: 5,
			Intervall: 2,
		},
	}
	GameDB.Effects[5] = sunburn_dot

	fireball := &GameDB_Spell{
		Id					:1,
		Name 				:"Fireball",
		Description 		:"Fireball",
		Visible				:true,
		ThreadModifier		:1,
		Cooldown			:0,
		GlobalCooldown		:1.5,
	
		MissileID			:1,
		EffectID			:0,
		IconID		 		:1,
		Speed				:60,
		ApplicationType     :GameDB_Spell_ApplicationType_Missile,
	
		BaseCost			:10,
		CostPerSec			:0,
		CostPercentage  	:0,
		
		CastTime 			:2,
		Range				:100,
		FacingFront			:true,
	
		TargetAuraRequired 	:0,
		CasterAuraRequired 	:0,
		
		Target				:GameDB_Spell_Target_Unit,
	
		Effect			    :[]*GameDB_Effect { GameDB.Effects[1] },
		
		InterruptedBy		:GameDB_Interrupt_OnMovement,
	}
	GameDB.Spells[1] = fireball
	
	frostbolt := &GameDB_Spell{
		Id					:2,
		Name 				:"Frostbolt",
		Description 		:"Frostbolt",
		Visible				:true,
		ThreadModifier		:1,
		Cooldown			:0,
		GlobalCooldown		:1.5,
	
		MissileID			:2,
		EffectID			:0,
		IconID		 		:2,
		Speed				:40,
		ApplicationType     :GameDB_Spell_ApplicationType_Missile,
	
		BaseCost			:5,
		CostPerSec			:0,
		CostPercentage  	:0,
		
		CastTime 			:0,
		Range				:70,
		FacingFront			:true,
	
		TargetAuraRequired 	:0,
		CasterAuraRequired 	:0,
		
		Target				:GameDB_Spell_Target_Unit,
	
		Effect			    :[]*GameDB_Effect { GameDB.Effects[2], GameDB.Effects[3] },
		
		InterruptedBy		:GameDB_Interrupt_None,
	}
	GameDB.Spells[2] = frostbolt

	sunburn := &GameDB_Spell{
		Id					:3,
		Name 				:"Sunburn",
		Description 		:"Sunburn",
		Visible				:true,
		ThreadModifier		:1,
		Cooldown			:0,
		GlobalCooldown		:1.5,
	
		MissileID			:3,
		EffectID			:0,
		IconID		 		:3,
		Speed				:1000,
		ApplicationType     :GameDB_Spell_ApplicationType_Beam,
	
		BaseCost			:5,
		CostPerSec			:0,
		CostPercentage  	:0,
		
		CastTime 			:0,
		Range				:70,
		FacingFront			:true,
	
		TargetAuraRequired 	:0,
		CasterAuraRequired 	:0,
		
		Target				:GameDB_Spell_Target_Unit,
	
		Effect			    :[]*GameDB_Effect { GameDB.Effects[4], GameDB.Effects[5] },
		
		InterruptedBy		:GameDB_Interrupt_None,
	}
	GameDB.Spells[3] = sunburn

	return GameDB
}