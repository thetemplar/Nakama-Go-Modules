package main

import (
	"fmt"
	"strconv"
)

//helper
func (p *PublicMatchState_Interactable) getInternalPlayer(state *MatchState) (*InternalPlayer) {
	return state.InternalPlayer[p.Id];
}

//fight
func (p *PublicMatchState_Interactable) applyAutoattackDamage(state *MatchState, creator string, slot GameDB_Item_Slot) {
	source :=  state.PublicMatchState.Interactable[creator]
	sourceChar := source.Character
	sourceClass := state.GetClassFromDB(sourceChar)
	
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)

	miss := (1 - sourceClass.getMeeleHitChance(sourceChar))
	dmgInput := float32(0)
	if slot == GameDB_Item_Slot_Weapon_MainHand {
		weapon := sourceChar.EquippedItemMainhandId
		dmgInput = randomFloatInt(state.GameDB.Items[weapon].DamageMin, state.GameDB.Items[weapon].DamageMax)
		fmt.Printf("\napplyAutoattackDamage (%v) %v to unit %v\n", slot, dmgInput, p.Id)
		speed := state.GameDB.Items[weapon].AttackSpeed * sourceClass.getMeeleAttackSpeed(sourceChar)
		dmgInput += sourceClass.getMeeleAttackPower(sourceChar) / (1 / speed)
		fmt.Printf("\nadded AP to-> %v\n", dmgInput)
	}
	if slot == GameDB_Item_Slot_Weapon_OffHand {
		weapon := sourceChar.EquippedItemOffhandId
		dmgInput = randomFloatInt(state.GameDB.Items[weapon].DamageMin, state.GameDB.Items[weapon].DamageMax)
		fmt.Printf("\napplyAutoattackDamage (%v) %v to unit %v\n", slot, dmgInput, p.Id)
		speed := state.GameDB.Items[weapon].AttackSpeed * sourceClass.getMeeleAttackSpeed(sourceChar)
		dmgInput += sourceClass.getMeeleAttackPower(sourceChar) / (1 / speed)
		fmt.Printf("\nadded AP to-> %v\n", dmgInput)
		miss *= 2
	}

	
	roll := randomPercentage()
	dodge := thisClass.getDodgeChance(thisChar)
	parry := thisClass.getParryChance(thisChar)
	behind := source.Position.isBehind(p.Position, p.Rotation)
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
			Source: PublicMatchState_CombatLogEntry_Autoattack,
			Type: &PublicMatchState_CombatLogEntry_MissedType{ fail },
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
	}
	

	block := thisClass.getBlockPercentage(thisChar)
	if behind || state.GameDB.Items[sourceChar.EquippedItemOffhandId].Type != GameDB_Item_Type_Weapon_Shield {
		block = 0
	}
	armor := thisClass.getArmor(thisChar) / (thisClass.getArmor(thisChar) + 40)
	dmgBlocked := float32(dmgInput) * (1 - (block + armor))
	dmgInput = dmgInput - dmgBlocked
	fmt.Printf("physical reduction by block (%v behind:%v) and armor (%v) by %v -> %v\n", block, behind, armor, dmgBlocked, dmgInput)

	roll = randomPercentage()
	crit := sourceClass.getMeeleCritChance(sourceChar)
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgBlocked = dmgBlocked * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - thisChar.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	thisChar.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;	
	fmt.Printf("applyDamage to %v: %v -> now:  %v\n", p.Id, (dmgInput + dmgInputCrit) - overkill, thisChar.CurrentHealth)

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
	sourceChar :=  source.Character
	sourceClass := state.GetClassFromDB(sourceChar)
	
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)

	dmgInput := randomFloatInt(effect.ValueMin, effect.ValueMax)
	fmt.Printf("\napplyAbilityDamage %v from effect %v to unit %v\n", dmgInput, effect, p.Id)
	
	roll := randomPercentage()
	miss := float32(0)
	switch effect.Type.(type) {
	case GameDB_Effect_Damage:
		miss = (1 - sourceClass.getSpellHitChance(sourceChar))
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

	resist := thisClass.getResistance(thisChar, effect.School)
	dmgResisted := dmgInput * resist / 100
	dmgInput = dmgInput - dmgResisted
	fmt.Printf("magical reduction by resistance (%v) by %v -> %v\n", resist, dmgResisted, dmgInput)

	roll = randomPercentage()
	crit := sourceClass.getSpellCritChance(sourceChar)
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgResisted = dmgResisted * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v from %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit, effect)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - thisChar.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	thisChar.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;	
	fmt.Printf("applyDamage to %v: %v from %v  -> now:  %v\n\n", p.Id, (dmgInput + dmgInputCrit) - overkill, effect, thisChar.CurrentHealth)

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

func (p *PublicMatchState_Interactable) containsEffectId(id int64, creator string) int64 {
    for i, a := range p.Auras {
        if a.EffectId == id && a.Creator == creator{
            return int64(i)
        }
    }
    return -1
}

//autoattack
func (p *PublicMatchState_Interactable) startAutoattack(state *MatchState, attacktype Client_Autoattack_Type) {
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)

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
		fmt.Printf("startAutoattack failed: %v\n", failedMessage)
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			Source: PublicMatchState_CombatLogEntry_Autoattack,
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
		mhEnd = int64(state.GameDB.Items[p.Character.EquippedItemMainhandId].AttackSpeed * thisClass.getMeeleAttackSpeed(thisChar) * float32(state.TickRate)) + state.PublicMatchState.Tick
		
		fmt.Printf("swinging mainhand at", mhEnd)
	}
	ohEnd := int64(0)
	if (p.Character.EquippedItemOffhandId > 0 && state.GameDB.Items[p.Character.EquippedItemOffhandId].AttackSpeed > 0) {
		ohEnd = int64(state.GameDB.Items[p.Character.EquippedItemOffhandId].AttackSpeed *  thisClass.getMeeleAttackSpeed(thisChar) * float32(state.TickRate)) + state.PublicMatchState.Tick
	
		fmt.Printf("swinging offhand at", ohEnd)
	}
	currentPlayerInternal.startAutoattackTimer(mhEnd, ohEnd, targetId)
}

func (p *PublicMatchState_Interactable) finishAutoattack(state *MatchState, slot GameDB_Item_Slot, targetId string) {
	target := state.PublicMatchState.Interactable[targetId]
	distance := p.Position.distance(target.Position)	
	
	if distance > state.GameDB.Items[p.Character.EquippedItemMainhandId].Range {
		fmt.Printf("startAutoattack failed: Out of Range!\n")
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			Source: PublicMatchState_CombatLogEntry_Autoattack,
			Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
				Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
				FailedMessage: "Out of Range!",
			}},
		}
		state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

		p.getInternalPlayer(state).stopAutoattackTimer()

		return
	}
	target.applyAutoattackDamage(state, p.Id, slot)
}

//casts
func (p *PublicMatchState_Interactable) startCast(state *MatchState, spell *GameDB_Spell) {		
	fmt.Printf("startCast: %v %v\n", spell.Id, spell.Name)
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)
	currentPlayerInternal := p.getInternalPlayer(state)
	failedMessage := ""

	if (p.GlobalCooldown > 0 || spell.IgnoresGCD == true) || currentPlayerInternal.CastingSpellId > 0 {
		failedMessage = "Cannot do that now!"
	}

	if spell.Target_Type != GameDB_Spell_Target_Type_None && p.Target == "" {
		failedMessage = "No Target!"
	}

	targetId := p.Target
	target := state.PublicMatchState.Interactable[targetId]
	distance := p.Position.distance(target.Position)		
	
	behind := target.Position.isFacedBy(p.Position, p.Rotation)
	if spell.FacingFront && !behind {
		failedMessage = "Is not facing target!"
	}

	if distance > spell.Range {	
		fmt.Printf("Out of Range: %v > %v\n", distance, spell.Range)
		failedMessage = "Out of Range!"
	}

	if IntersectingBorders(p.Position, target.Position, state.Map) {
		failedMessage = "Not in Line of Sight!"
	}

	if failedMessage != "" {
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{spell.Id},
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
		p.finishCast(state, spell, targetId)
	} else {
		end := int64(spell.CastTime * thisClass.getSpellAttackSpeed(thisChar) * float32(state.TickRate)) + state.PublicMatchState.Tick
		target := ""
		if spell.Target_Type != GameDB_Spell_Target_Type_None {
			target = targetId
		} 
		currentPlayerInternal.startCastTimer(spell.Id, end, target)
	}
}

func (p *PublicMatchState_Interactable) cancelCast(state *MatchState) {	

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

func (p *PublicMatchState_Interactable) finishCast(state *MatchState, spell *GameDB_Spell, targetId string) {
	if !IntersectingBorders(p.Position, state.PublicMatchState.Interactable[targetId].Position, state.Map) {		
		fmt.Printf("finish cast spell: %v\n", spell.Id)
		p.GlobalCooldown = spell.GlobalCooldown
		proj := &PublicMatchState_Projectile{
			Id: "p_" + strconv.FormatInt(state.ProjectileCounter, 16),
			SpellId: spell.Id,
			Position: &PublicMatchState_Vector2Df {
				X: p.Position.X,
				Y: p.Position.Y,
			},
			Rotation: p.Rotation,
			CreatedAtTick: state.PublicMatchState.Tick,
			Target: targetId,
			Speed: spell.Speed,
			Creator: p.Id,
		}
		state.PublicMatchState.Projectile[proj.Id] = proj
		state.ProjectileCounter++	
	} else {
		clEntry := &PublicMatchState_CombatLogEntry {
			Timestamp: state.PublicMatchState.Tick,
			SourceId: p.Id,
			SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{spell.Id},
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
func (p *PublicMatchState_Interactable) recalcStats(state *MatchState) {
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