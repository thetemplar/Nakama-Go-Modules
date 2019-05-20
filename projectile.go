package main

import (
	"math"
	"fmt"
)

func (p PublicMatchState_Projectile) Run(state *MatchState, projectile *PublicMatchState_Projectile, tickrate int) {
	target := state.PublicMatchState.Interactable[projectile.Target]					
	distance := float32(math.Sqrt(math.Pow(float64(projectile.Position.X - target.Position.X), 2) + math.Pow(float64(projectile.Position.Y - target.Position.Y), 2)))	
	direction := PublicMatchState_Vector2Df {
		X: target.Position.X - projectile.Position.X,
		Y: target.Position.Y - projectile.Position.Y,
	}
	direction.X /= distance
	direction.Y /= distance

	if distance <= (projectile.Speed / float32(tickrate)) {
		//impact
		fmt.Printf("%v impact\n", projectile.Id)
		spell := state.GameDB.Spells[projectile.SpellId]	
		projectile.Hit(state, target, projectile, spell)
		delete(state.PublicMatchState.Projectile, projectile.Id)
		projectile = nil
		return
	}

	projectile.Position.X = projectile.Position.X + (direction.X * (projectile.Speed / float32(tickrate)))
	projectile.Position.Y = projectile.Position.Y + (direction.Y * (projectile.Speed / float32(tickrate)))

	fmt.Printf("%v @ %v | %v\n", projectile.Id, projectile.Position.X, projectile.Position.Y)
}


func (p PublicMatchState_Projectile) Hit(state *MatchState, target *PublicMatchState_Interactable, projectile *PublicMatchState_Projectile, spell *GameDB_Spell) {
	for _, effect := range spell.Effect { 
		fmt.Printf("Apply Effect on Hit %v\n", effect)
		if effect.Duration > 0 {
			i := target.containsEffectId(effect.Id, projectile.Creator)
			if i == -1 {
				aura := &PublicMatchState_Aura{
					CreatedAtTick: state.PublicMatchState.Tick,
					EffectId: effect.Id,
					Creator: projectile.Creator,
				}
				target.Auras = append(target.Auras, aura)

				clEntry := &PublicMatchState_CombatLogEntry {
					Timestamp: state.PublicMatchState.Tick,
					SourceId: aura.Creator,
					DestinationId: target.Id,
					SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
					Source: PublicMatchState_CombatLogEntry_Spell,
					Type: &PublicMatchState_CombatLogEntry_Aura{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Aura{
						Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Aura_Applied,
					}},
				}
				state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

				switch effect.Type.(type) {
				case *GameDB_Effect_Apply_Aura_Mod:
					target.recalcStats(state)
				}
			} else {
				target.Auras[i].CreatedAtTick = state.PublicMatchState.Tick
				target.Auras[i].AuraTickCount = 0

				clEntry := &PublicMatchState_CombatLogEntry {
					Timestamp: state.PublicMatchState.Tick,
					SourceId: target.Auras[i].Creator,
					DestinationId: target.Id,
					SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
					Source: PublicMatchState_CombatLogEntry_Spell,
					Type: &PublicMatchState_CombatLogEntry_Aura{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Aura{
						Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Aura_Refreshed,
					}},
				}
				state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
			}
		} else {
			switch effect.Type.(type) {
			case *GameDB_Effect_Damage:
				dmg := randomInt(effect.Type.(*GameDB_Effect_Damage).ValueMin, effect.Type.(*GameDB_Effect_Damage).ValueMax);
				dmgCrit := int32(0)
				if randomInt(0, 100) > 20 {
					dmgCrit = dmg
				}
				overkill := target.applyDamage(dmg + dmgCrit)
			
				clEntry := &PublicMatchState_CombatLogEntry {
					Timestamp: state.PublicMatchState.Tick,
					SourceId: projectile.Creator,
					DestinationId: target.Id,
					SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
					Source: PublicMatchState_CombatLogEntry_Spell,
					Type: &PublicMatchState_CombatLogEntry_Damage{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Damage{
						Amount: dmg,
						Critical: dmgCrit,
						Overkill: overkill,
					}},
				}
				state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
			}
		}
	}

}