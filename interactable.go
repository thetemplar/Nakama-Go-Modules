package main

import (
	"fmt"
	"strconv"
	"math"
)

func (p PublicMatchState_Interactable) getInternalPlayer(state *MatchState) (*InternalPlayer) {
	return state.InternalPlayer[p.Id];
}

func (p PublicMatchState_Interactable) applyDamage(dmg int32) (overkill int32){	
	overkill = dmg -  p.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.CurrentHealth -= dmg - overkill;
	return overkill
}

func (p PublicMatchState_Interactable) startCast(state *MatchState, spellId int64) {
	currentPlayerInternal := p.getInternalPlayer(state)
	failedMessage := ""

	if p.GlobalCooldown > 0 || currentPlayerInternal.CastingSpellId > 0 {
		failedMessage = "Cannot do that now!"
	}

	if state.GameDB.Spells[spellId].Target != GameDB_Spell_Target_None && p.Target == "" {
		failedMessage = "No Target!"
	}

	targetId := p.Target
	target := state.PublicMatchState.Interactable[targetId]
	distance := float32(math.Sqrt(math.Pow(float64(p.Position.X - target.Position.X), 2) + math.Pow(float64(p.Position.Y - target.Position.Y), 2)))	
	
	if distance > state.GameDB.Spells[spellId].Range {
		failedMessage = "Out of Range!"
	}

	if IntersectingBorders(p.Position, target.Position, state.Map) {
		failedMessage = "Not in Line of Sight!"
	}

	if failedMessage != "" {
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{spellId},
			Source: PublicMatchState_CombatLogEntry_Spell,
			Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
				Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
				FailedMessage: failedMessage,
			}},
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

		return
	}
	
	if state.GameDB.Spells[spellId].CastTime == 0 {
		p.finishCast(state, spellId, targetId)
	} else {
		currentPlayerInternal.CastingSpellId = spellId
		currentPlayerInternal.CastingTickStarted = state.PublicMatchState.Tick
		if state.GameDB.Spells[spellId].Target != GameDB_Spell_Target_None {
			currentPlayerInternal.CastingTargeted = targetId
		} else {
			currentPlayerInternal.CastingTargeted = ""
		}
	}
}

func (p PublicMatchState_Interactable) cancelCast(state *MatchState) {	

	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: p.Id,
		SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{p.getInternalPlayer(state).CastingSpellId},
		Source: PublicMatchState_CombatLogEntry_Spell,
		Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
			Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
			FailedMessage: "Cast canceled by Movement!",
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

	p.getInternalPlayer(state).CastingSpellId = -1
	p.getInternalPlayer(state).CastingTickStarted = 0
	p.getInternalPlayer(state).CastingTargeted = ""
}

func (p PublicMatchState_Interactable) finishCast(state *MatchState, spellId int64, targetId string) {
	if !IntersectingBorders(p.Position, state.PublicMatchState.Interactable[targetId].Position, state.Map) {		
		fmt.Printf("finish cast spell: %v\n", spellId)
		p.GlobalCooldown = state.GameDB.Spells[spellId].GlobalCooldown
		proj := &PublicMatchState_Projectile{
			Id: "p_" + strconv.FormatInt(state.ProjectileCounter, 16),
			SpellId: spellId,
			Position: &PublicMatchState_Vector2Df {
				X: p.Position.X,
				Y: p.Position.Y,
			},
			Rotation: p.Rotation,
			CreatedAtTick: state.PublicMatchState.Tick,
			Target: targetId,
			Speed: state.GameDB.Spells[spellId].Speed,
		}
		state.PublicMatchState.Projectile[proj.Id] = proj
		state.ProjectileCounter++	
	} else {
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{spellId},
			Source: PublicMatchState_CombatLogEntry_Spell,
			Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
				Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
				FailedMessage: "Not in Line of Sight!",
			}},
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
	}		
}


func (p PublicMatchState_Interactable) recalcStats(state *MatchState) {
	p.getInternalPlayer(state).StatModifiers = PlayerStats {}
	for _, aura := range p.Auras {
		effect := state.GameDB.Effects[aura.EffectId]
		
		switch effect.Type.(type) {
		case *GameDB_Effect_Apply_Aura_Mod:
			if effect.Type.(*GameDB_Effect_Apply_Aura_Mod).Stat == GameDB_Stat_Speed && effect.Type.(*GameDB_Effect_Apply_Aura_Mod).Value > p.getInternalPlayer(state).StatModifiers.MovementSpeed {
				p.getInternalPlayer(state).StatModifiers.MovementSpeed = effect.Type.(*GameDB_Effect_Apply_Aura_Mod).Value
			}
		}
	}
}