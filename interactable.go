package main

import (
	"fmt"
	"strconv"
)

//helper
func (p PublicMatchState_Interactable) getInternalPlayer(state *MatchState) (*InternalPlayer) {
	return state.InternalPlayer[p.Id];
}

//fight
func (p *PublicMatchState_Interactable) applyAutoattackDamage(state *MatchState, creator string) {
	source :=  state.PublicMatchState.Interactable[creator]

	//behind?
	behind := source.Position.isBehind(p.Position, p.Rotation)

	dmgInput := float32(1)//randomFloatInt(effect.Type.(*GameDB_Effect_Damage).ValueMin, effect.Type.(*GameDB_Effect_Damage).ValueMax)
	fmt.Printf("\napplyAutoattackDamage %v to unit %v\n", dmgInput, p.Id)
	
	roll := randomPercentage()
	miss := (1 - source.Character.getMeeleHitChance())
	dodge := p.Character.getDodgeChance()
	parry := p.Character.getParryChance()
	if behind {
		dodge = 0
		parry = 0
	}

	fail := PublicMatchState_CombatLogEntry_CombatLogEntry_Missed(-1)
	if roll <= miss {
		fmt.Printf("miss (%v/%v) damage to %v: %v\n", roll, miss, p.Id, dmgInput)
		fail = PublicMatchState_CombatLogEntry_Missed
	} 
	if roll <= miss + dodge {
		fmt.Printf("dodge (%v/%v) damage to %v: %v\n", roll, parry, p.Id, dmgInput)
		fail = PublicMatchState_CombatLogEntry_Dodged
	} 
	if roll <= miss + dodge + parry {
		fmt.Printf("parry (%v/%v) damage to %v: %v\n", roll, parry, p.Id, dmgInput)
		fail = PublicMatchState_CombatLogEntry_Parried
	} 	
	if fail > -1 {		
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: creator,
			DestinationId: p.Id,
			//SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
			Source: PublicMatchState_CombatLogEntry_Spell,
			Type: &PublicMatchState_CombatLogEntry_MissedType{ fail },
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
	}
	

	block := p.Character.getBlockPercentage()
	if behind {
		block = 0
	}
	armor := p.Character.getArmor() / (p.Character.getArmor() + 40)
	dmgBlocked := float32(dmgInput) * (1 - (block + armor))
	dmgInput = dmgInput - dmgBlocked
	fmt.Printf("physical reduction by block (%v behind:%v) and armor (%v) by %v -> %v\n", block, behind, armor, dmgBlocked, dmgInput)

	roll = randomPercentage()
	crit := source.Character.getMeeleCritChance()
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgBlocked = dmgBlocked * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - p.Character.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.Character.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;	
	fmt.Printf("applyDamage to %v: %v -> now:  %v\n", p.Id, (dmgInput + dmgInputCrit) - overkill, p.Character.CurrentHealth)

	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: creator,
		DestinationId: p.Id,
		//SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
		Source: PublicMatchState_CombatLogEntry_Spell,
		Type: &PublicMatchState_CombatLogEntry_Damage{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Damage{
			Amount: dmgInput,
			Resisted: 0,
			Blocked: dmgBlocked,
			Absorbed: 0,
			Critical: dmgInputCrit,
			Overkill: overkill,			
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
}

func (p *PublicMatchState_Interactable) applyAbilityDamage(state *MatchState, effect *GameDB_Effect, creator string) {
	source :=  state.PublicMatchState.Interactable[creator]

	dmgInput := randomFloatInt(effect.ValueMin, effect.ValueMax)
	fmt.Printf("\napplyAbilityDamage %v from effect %v to unit %v\n", dmgInput, effect, p.Id)
	
	roll := randomPercentage()
	miss := float32(0)
	switch effect.Type.(type) {
	case GameDB_Effect_Damage:
		miss = (1 - source.Character.getSpellHitChance())
	}
	if roll <= miss {
		fmt.Printf("miss (%v/%v) damage to %v: %v from %v\n", roll, miss, p.Id, dmgInput, effect)		
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: creator,
			DestinationId: p.Id,
			SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
			Source: PublicMatchState_CombatLogEntry_Spell,
			Type: &PublicMatchState_CombatLogEntry_MissedType{ PublicMatchState_CombatLogEntry_Missed },
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
	}

	resist := p.Character.getResistance(effect.School)
	dmgResisted := dmgInput * resist
	dmgInput = dmgInput - dmgResisted
	fmt.Printf("magical reduction by resistance (%v) by %v -> %v\n", resist, dmgResisted, dmgInput)

	roll = randomPercentage()
	crit := source.Character.getSpellCritChance()
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgResisted = dmgResisted * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v from %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit, effect)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - p.Character.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.Character.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;	
	fmt.Printf("applyDamage to %v: %v from %v  -> now:  %v\n\n", p.Id, (dmgInput + dmgInputCrit) - overkill, effect, p.Character.CurrentHealth)

	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: creator,
		DestinationId: p.Id,
		SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
		Source: PublicMatchState_CombatLogEntry_Spell,
		Type: &PublicMatchState_CombatLogEntry_Damage{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Damage{
			Amount: dmgInput,
			Resisted: dmgResisted,
			Blocked: 0,
			Absorbed: 0,
			Critical: dmgInputCrit,
			Overkill: overkill,			
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
}

func (p PublicMatchState_Interactable) containsEffectId(id int64, creator string) int64 {
    for i, a := range p.Auras {
        if a.EffectId == id && a.Creator == creator{
            return int64(i)
        }
    }
    return -1
}

//autoattack
func (p PublicMatchState_Interactable) startAutoattack(state *MatchState, attacktype Client_Autoattack_Type) {
	currentPlayerInternal := p.getInternalPlayer(state)	
	failedMessage := ""

	if currentPlayerInternal.CastingSpellId > 0 && state.GameDB.Spells[currentPlayerInternal.CastingSpellId].IgnoresWeaponswing == false {
		failedMessage = "Cannot do that now!"
		
	}

	if p.Target == "" {
		failedMessage = "No Target!"
	}

	targetId := p.Target
	target := state.PublicMatchState.Interactable[targetId]
	distance := p.Position.distance(target.Position)	
	
	if distance > state.GameDB.Items[p.Character.EquippedItemMainhandId].Range {
		failedMessage = "Out of Range!"
	}

	if distance > 1 {
		if IntersectingBorders(p.Position, target.Position, state.Map) {
			failedMessage = "Not in Line of Sight!"
		}
	}

	if failedMessage != "" {
		fmt.Printf("startAutoattack failed: %v", failedMessage)
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			Source: PublicMatchState_CombatLogEntry_Swing,
			Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
				Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
				FailedMessage: failedMessage,
			}},
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

		currentPlayerInternal.stopAutoattackTimer()

		return
	}
	mhEnd := int64(0)
	if (p.Character.EquippedItemMainhandId > 0 && state.GameDB.Items[p.Character.EquippedItemMainhandId].AttackSpeed > 0) {
		mhEnd = int64(state.GameDB.Items[p.Character.EquippedItemMainhandId].AttackSpeed * p.Character.getMeeleAttackSpeed() * float32(state.TickRate)) + state.PublicMatchState.Tick
		
		fmt.Printf("swinging mainhand at", mhEnd)
	}
	ohEnd := int64(0)
	if (p.Character.EquippedItemOffhandId > 0 && state.GameDB.Items[p.Character.EquippedItemOffhandId].AttackSpeed > 0) {
		ohEnd = int64(state.GameDB.Items[p.Character.EquippedItemOffhandId].AttackSpeed * p.Character.getMeeleAttackSpeed() * float32(state.TickRate)) + state.PublicMatchState.Tick
	
		fmt.Printf("swinging offhand at", ohEnd)
	}
	currentPlayerInternal.startAutoattackTimer(mhEnd, ohEnd, targetId)
}

func (p PublicMatchState_Interactable) finishAutoattack(state *MatchState, slot GameDB_Item_Slot, targetId string) {
	fmt.Printf("TODO: func (p PublicMatchState_Interactable) finishAutoattack(state *MatchState, slot GameDB_Item_Slot) %v\n", slot)
	state.PublicMatchState.Interactable[targetId].Character.CurrentHealth -= 10;	
}

//casts
func (p PublicMatchState_Interactable) startCast(state *MatchState, spellId int64) {
	currentPlayerInternal := p.getInternalPlayer(state)
	spell := state.GameDB.Spells[spellId]
	failedMessage := ""

	if (p.GlobalCooldown > 0 || spell.IgnoresGCD == true) || currentPlayerInternal.CastingSpellId > 0 {
		failedMessage = "Cannot do that now!"
	}

	if spell.Target != GameDB_Spell_Target_None && p.Target == "" {
		failedMessage = "No Target!"
	}

	targetId := p.Target
	target := state.PublicMatchState.Interactable[targetId]
	distance := p.Position.distance(target.Position)	
	
	if distance > spell.Range {
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

	if spell.IgnoresWeaponswing == false {
		currentPlayerInternal.stopAutoattackTimer()
	}
	
	if spell.CastTime == 0 {
		p.finishCast(state, spellId, targetId)
	} else {
		end := int64(spell.CastTime * p.Character.getSpellAttackSpeed() * float32(state.TickRate)) + state.PublicMatchState.Tick
		target := ""
		if spell.Target != GameDB_Spell_Target_None {
			target = targetId
		} 
		currentPlayerInternal.startCastTimer(spellId, end, target)
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

	p.getInternalPlayer(state).stopCastTimer()
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
			Creator: p.Id,
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

//stats
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