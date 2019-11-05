package main

import (
	"fmt"
	"strconv"
	"Nakama-Go-Modules/GameDB"
)

func SpellHit(state *MatchState, target *InternalInteractable, source string, spell *GameDB.Spell) {
	for _, effect := range spell.ApplyEffect { 
		fmt.Printf("Apply Effect on Hit %v\n", effect)
		if effect.Duration > 0 {
			i := target.containsEffectId(effect.Id, source)
			if i == -1 {
				aura := &PublicMatchState_Aura{
					CreatedAtTick: state.PublicMatchState.Tick,
					EffectId: effect.Id,
					Creator: source,
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
				case *GameDB.Effect_Apply_Aura_Mod:
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
			case *GameDB.Effect_Damage:
				target.applyAbilityDamage(state, effect, source)
			case *GameDB.Effect_Heal:
				target.applyAbilityHeal(state, effect, source)
			}
		}
	}

}

//AoE
func CreateAoEArea(state *MatchState, spell *GameDB.Spell, pos *Vector2Df, creator string) {
	for _, effect := range spell.ApplyEffect { 
		fmt.Printf("Apply Effect on AoE-Area %v\n", effect)
		area := &PublicMatchState_Area{
			Id: "a_" + strconv.FormatInt(state.ProjectileCounter, 16),
			EffectId: effect.Id,
			CreatedAtTick: state.PublicMatchState.Tick,
			Creator: creator,
			Position: pos,
			AreaTickCount: 0,
		}
		state.ProjectileCounter++

		state.PublicMatchState.Area[area.Id] = area
	}
}