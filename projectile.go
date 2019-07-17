package main

import (
	"fmt"
)

func (p PublicMatchState_Projectile) Run(state *MatchState, projectile *PublicMatchState_Projectile, tickrate int) {
	target := state.PublicMatchState.Interactable[projectile.Target]					
	distance := projectile.Position.distance(target.Position)	
	direction := PublicMatchState_Vector2Df {
		X: target.Position.X - projectile.Position.X,
		Y: target.Position.Y - projectile.Position.Y,
	}
	direction.X /= distance
	direction.Y /= distance

	if distance <= (projectile.Speed / float32(tickrate)) {
		//impact
		spell := state.GameDB.Spells[projectile.SpellId]	
		projectile.Hit(state, target, projectile, spell)
		delete(state.PublicMatchState.Projectile, projectile.Id)
		projectile = nil
		return
	}

	projectile.Position.X = projectile.Position.X + (direction.X * (projectile.Speed / float32(tickrate)))
	projectile.Position.Y = projectile.Position.Y + (direction.Y * (projectile.Speed / float32(tickrate)))
}


func (p PublicMatchState_Projectile) Hit(state *MatchState, target *PublicMatchState_Interactable, projectile *PublicMatchState_Projectile, spell *GameDB_Spell) {
	for _, effect := range spell.ApplyEffect { 
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
				target.applyAbilityDamage(state, effect, projectile.Creator)
			}
		}
	}

}