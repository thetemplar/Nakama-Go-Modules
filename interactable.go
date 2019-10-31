package main

import (
	"fmt"
	"strconv"
	"math"
	"Nakama-Go-Modules/GameDB"
)

//helper
func (p *PublicMatchState_Interactable) getInternalPlayer(state *MatchState) (*InternalPlayer) {
	return state.InternalPlayer[p.Id];
}

//movement
func (p *PublicMatchState_Interactable) PerformMovement(state *MatchState, xAxis, yAxis, rotation float32) {
	p.Rotation = rotation;
	//clamb [-1..1]
	length := float32(math.Sqrt(math.Pow(float64(xAxis), 2) + math.Pow(float64(yAxis), 2)))
	if length == 0 {
		return
	}
	if length > 1 {
		xAxis /= length
		yAxis /= length
	}
	mod := p.getInternalPlayer(state).StatModifiers.MovementSpeedModifier
	xAxis *= mod
	yAxis *= mod

	currentPlayerInternal := p.getInternalPlayer(state)	

	moveMsgCount := currentPlayerInternal.MoveMessageCountThisFrame
	if(p.Type == PublicMatchState_Interactable_NPC)	{
		moveMsgCount = 1
	}

	add := PublicMatchState_Vector2Df {
		X: xAxis / float32(moveMsgCount) * ((20) / float32(state.TickRate)),
		Y: yAxis / float32(moveMsgCount) * ((20) / float32(state.TickRate)),
	}
	
	fmt.Printf("add %v %v > %v %v\n", xAxis, yAxis, add.X, add.Y)
	
	if math.IsNaN(float64(add.X)) || math.IsNaN(float64(add.Y)) {
		return
	}

	if currentPlayerInternal.CastingSpellId > 0 && state.GameDB.Spells[currentPlayerInternal.CastingSpellId].InterruptedBy != GameDB.Interrupt_Type_None && (xAxis != 0 || yAxis != 0) {
		fmt.Printf("cancelCast %v\n", currentPlayerInternal.CastingSpellId)
		p.cancelCast(state)
	}

	rotatedAdd := add.rotate(rotation)
		

	p.Position.X += rotatedAdd.X;
	p.Position.Y += rotatedAdd.Y;
	p.Rotation = rotation;
	

	//am i still in my triangle?	
	if currentPlayerInternal.TriangleIndex >= 0 {
		isItIn, _, _, _ := state.Map.Triangles[currentPlayerInternal.TriangleIndex].isInTriangle(p.Position)
		if !isItIn {
			currentPlayerInternal.TriangleIndex = -1
		}
	}

	//no current triangle_index?
	if currentPlayerInternal.TriangleIndex < 0 {
		//find triangle I am in
		found := false
		for i, triangle := range state.Map.Triangles {
			isItIn, _, _, _ := triangle.isInTriangle(p.Position)
			if isItIn {
				currentPlayerInternal.TriangleIndex = int64(i)
				found = true
				break;
			}
		}
		if !found {
			p.Position.X -= rotatedAdd.X;
			p.Position.Y -= rotatedAdd.Y;

			//better: find edge between this triangle and destination point, then move along the edge
		} 
	}	
}

func (p *PublicMatchState_Interactable) RotateTowardsTogarget(state *MatchState, targetPos *PublicMatchState_Vector2Df) {
	p.Rotation = float32(math.Atan2(float64(targetPos.X), float64(targetPos.Y)) * 57.2957795131);
}

//regen
func (p *PublicMatchState_Interactable) Regen(state *MatchState, hpPercent, powerPercent float64) {	
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)

	if thisChar.CurrentHealth <= 0 {
		return
	}

	thisChar.CurrentHealth += float32(math.Max(float64(thisChar.getHpRegen(thisClass)) * hpPercent, 0));
	thisChar.CurrentPower += float32(math.Max(float64(thisChar.getManaRegen(thisClass)) * powerPercent, 0));

	thisChar.CurrentHealth = float32(math.Min(float64(thisChar.CurrentHealth), float64(thisChar.getMaxHp(thisClass))))
	thisChar.CurrentPower = float32(math.Min(float64(thisChar.CurrentPower), float64(thisChar.getMaxMana(thisClass))))
}

//fight
func (p *PublicMatchState_Interactable) applyAutoattackDamage(state *MatchState, creator string, slot GameDB.Item_Slot) {
	source :=  state.PublicMatchState.Interactable[creator]
	sourceChar := source.Character
	sourceClass := state.GetClassFromDB(sourceChar)
	
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)
	if(!p.IsEngaged || p.Target == "") {
		p.Target = source.Id
	}
	p.IsEngaged = true

	miss := (1 - sourceChar.getMeeleHitChance(sourceClass))
	dmgInput := float32(0)
	
	var item *GameDB.Item
	if (slot == GameDB.Item_Slot_Weapon_MainHand || slot == GameDB.Item_Slot_Weapon_BothHands) {
		item = thisClass.Mainhand
		if (item == nil) {
			fmt.Printf("\nERROR NO WEAPON IN SLOT Item_Slot_Weapon_MainHand/Item_Slot_Weapon_BothHands\n")
			return
		}
	} else if (slot == GameDB.Item_Slot_Weapon_OffHand) {
		item = thisClass.Offhand
		if (item == nil) {
			fmt.Printf("\nERROR NO WEAPON IN SLOT Item_Slot_Weapon_OffHand\n")
			return
		}
	}
	dmgInput = randomFloatInt(item.DamageMin, item.DamageMax)
	fmt.Printf("\napplyAutoattackDamage (%v) %v+%v to unit %v\n", slot, dmgInput, sourceChar.getMeeleAttackPower(sourceClass), p.Id)
	dmgInput += sourceChar.getMeeleAttackPower(sourceClass)

	if slot == GameDB.Item_Slot_Weapon_OffHand {
		miss *= 2
	}
	
	roll := randomPercentage()
	dodge := thisChar.getDodgeChance(thisClass)
	parry := thisChar.getParryChance(thisClass)
	behind := source.Position.isBehind(p.Position, p.Rotation)
	if behind {
		dodge = 0
		parry = 0
	}


	fail := PublicMatchState_CombatLogEntry_CombatLogEntry_Missed(-1)
	if roll <= miss {
		fail = PublicMatchState_CombatLogEntry_Missed
	} else if roll <= miss + dodge {
		fail = PublicMatchState_CombatLogEntry_Dodged
	} else if roll <= miss + dodge + parry {
		fail = PublicMatchState_CombatLogEntry_Parried
	}
	
	fmt.Printf("rolled: %v - table: [miss: %v | dodge: %v | parry: %v ] - damage to %v: %v - failed: %v\n", roll, miss, miss + dodge, miss + dodge + parry, p.Id, dmgInput, fail)
		

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
		return;
	}
	
	armor := thisChar.getArmor(thisClass)
	dmgInput -= armor

	block := thisChar.getBlockPercentage(thisClass)
	if behind || (thisClass.Offhand != nil && thisClass.Offhand.Type != GameDB.Item_Type_Weapon_Shield) {
		block = 0
	}
	dmgBlocked := dmgInput * block
	dmgInput -= dmgBlocked
	fmt.Printf("physical reduction by blockpercentage %v (behind:%v) and armor (%v) by -> %v\n", block, behind, armor, dmgInput)

	if(dmgInput < 0) {
		dmgInput = 0
	}
	
	roll = randomPercentage()
	crit := sourceChar.getMeeleCritChance(sourceClass)
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
	p.getInternalPlayer(state).LastHealthDrainTick = state.PublicMatchState.Tick
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

func (p *PublicMatchState_Interactable) applyAbilityDamage(state *MatchState, effect *GameDB.Effect, creator string) {
	source :=  state.PublicMatchState.Interactable[creator]
	sourceChar :=  source.Character
	sourceClass := state.GetClassFromDB(sourceChar)
	
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)
	if(!p.IsEngaged || p.Target == "") {
		p.Target = source.Id
	}
	p.IsEngaged = true

	dmgInput := randomFloatInt(effect.ValueMin, effect.ValueMax)
	fmt.Printf("\napplyAbilityDamage %v from effect %v to unit %v\n", dmgInput, effect, p.Id)
	
	roll := randomPercentage()
	miss := float32(0)
	switch effect.Type.(type) {
	case GameDB.Effect_Damage:
		miss = (1 - sourceChar.getSpellHitChance(sourceClass))
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

	resist := thisChar.getResistance(thisClass, effect.School)
	dmgResisted := dmgInput * resist / 100
	dmgInput = dmgInput - dmgResisted
	fmt.Printf("magical reduction by resistance (%v) by %v -> %v\n", resist, dmgResisted, dmgInput)

	roll = randomPercentage()
	crit := sourceChar.getSpellCritChance(sourceClass)
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
	p.getInternalPlayer(state).LastHealthDrainTick = state.PublicMatchState.Tick
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
func (p *PublicMatchState_Interactable) startAutoattack(state *MatchState, attacktype Client_Message_Client_Autoattack_Type) {
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
	
	
	
	if distance > thisClass.Mainhand.Range {
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
	if (thisClass.Mainhand != nil && thisClass.Mainhand.AttackSpeed > 0) {
		mhEnd = int64(thisClass.Mainhand.AttackSpeed * thisChar.getMeeleAttackSpeed(thisClass) * float32(state.TickRate)) + state.PublicMatchState.Tick
		
		fmt.Printf("swinging mainhand at  %v\n", mhEnd)
	}
	ohEnd := int64(0)
	if (thisClass.Offhand != nil && thisClass.Offhand.AttackSpeed > 0) {
		ohEnd = int64(thisClass.Offhand.AttackSpeed * thisChar.getMeeleAttackSpeed(thisClass) * float32(state.TickRate)) + state.PublicMatchState.Tick
	
		fmt.Printf("swinging offhand at  %v\n", ohEnd)
	}
	currentPlayerInternal.startAutoattackTimer(mhEnd, ohEnd, targetId)
}

func (p *PublicMatchState_Interactable) finishAutoattack(state *MatchState, slot GameDB.Item_Slot, targetId string) {
	target := state.PublicMatchState.Interactable[targetId]
	distance := p.Position.distance(target.Position)	

	if distance > state.GetClassFromDB(p.Character).Mainhand.Range {
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
func (p *PublicMatchState_Interactable) startCast(state *MatchState, spell *GameDB.Spell) {		
	fmt.Printf("startCast: %v %v\n", spell.Id, spell.Name)
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)
	currentPlayerInternal := p.getInternalPlayer(state)
	failedMessage := ""

	if (p.GlobalCooldown > 0 || spell.IgnoresGCD == true) || currentPlayerInternal.CastingSpellId > 0 {
		failedMessage = "Cannot do that now!"
	}
	
	manacost := float32(spell.BaseCost) + float32(spell.CostPercentage) * thisChar.getMaxMana(thisClass);
	if manacost > thisChar.CurrentPower {
		failedMessage = "Not enough Mana!"
	}

	targetId := "" 
	if (spell.Target_Type == GameDB.Spell_Target_Type_Self) {
		targetId = p.Id
	} else if (spell.Target_Type == GameDB.Spell_Target_Type_None) {
		//nothing so far
	} else {
		if p.Target == "" {
			failedMessage = "No Target!"
		}

		targetId = p.Target
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
	}

	if spell.IgnoresWeaponswing == false {
		currentPlayerInternal.stopAutoattackTimer()
	}
	
	p.getInternalPlayer(state).LastPowerDrainTick = state.PublicMatchState.Tick
	if spell.CastTime == 0 {
		p.finishCast(state, spell, targetId)
	} else {
		end := int64(spell.CastTime * (float32(1) - thisChar.getSpellAttackSpeed(thisClass)) * float32(state.TickRate)) + state.PublicMatchState.Tick
		target := ""
		if spell.Target_Type != GameDB.Spell_Target_Type_None {
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

func (p *PublicMatchState_Interactable) finishCast(state *MatchState, spell *GameDB.Spell, targetId string) {
	thisChar := p.Character
	thisClass := state.GetClassFromDB(p.Character)

	if !IntersectingBorders(p.Position, state.PublicMatchState.Interactable[targetId].Position, state.Map) {
		thisChar.CurrentPower -= float32(spell.BaseCost);
		thisChar.CurrentPower -= float32(spell.CostPercentage) * thisChar.getMaxMana(thisClass);
		p.getInternalPlayer(state).LastPowerDrainTick = state.PublicMatchState.Tick	

		fmt.Printf("finish cast spell: %v (mana now: %v)\n", spell.Id, thisChar.CurrentPower)

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
	p.getInternalPlayer(state).StatModifiers.MovementSpeedModifier = 1
	for _, aura := range p.Auras {
		effect := state.GameDB.Effects[aura.EffectId]
		
		switch effect.Type.(type) {
		case *GameDB.Effect_Apply_Aura_Mod:
			if effect.Type.(*GameDB.Effect_Apply_Aura_Mod).Stat == GameDB.Stat_Speed {
				p.getInternalPlayer(state).StatModifiers.MovementSpeedModifier *= effect.Type.(*GameDB.Effect_Apply_Aura_Mod).Value
			}
		}
	}
	fmt.Printf("move: %v = %v\n", p.Id, p.getInternalPlayer(state).StatModifiers.MovementSpeedModifier)
}