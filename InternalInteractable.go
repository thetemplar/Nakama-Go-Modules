package main

import (
	"github.com/heroiclabs/nakama-common/runtime"
	"fmt"
	"strconv"
	"math"
	"Nakama-Go-Modules/GameDB"	
)

//InternalInteractable is an extended struct from the protobuf PublicMatchState_Interactable struct
type InternalInteractable struct {
	*PublicMatchState_Interactable

	//played by user
	Presence                	runtime.Presence

	//messages from client
	LastMovement             	*Client_Message_Client_Movement
	LastMessageServerTick   	int64
	LastMessageClientTick   	int64
	MissingCount				int
	MoveMessageCountThisFrame  	int

	//movement	
	TriangleIndex 				int64
	
	//casts
	CastingSpellId 				int64
	CastingTickEnd	 			int64
	CastingTargeted 			interface{}

	//autoattacks
	Autoattacking				bool
	AutoattackMainhandTickEnd	int64
	AutoattackOffhandTickEnd	int64
	AutoattackTargeted			string

	//playerstats
	StatModifiers				PlayerStats

	//regen
	LastRegenTick				int64
	LastHealthDrainTick			int64
	LastPowerDrainTick			int64

	//fights
	Act Act

	//cooldowns	
	Cooldowns     				map[int64]int64
	
	//npc
	npcState 					interface{}
}

type Act func(state *MatchState, p *InternalInteractable)

type PlayerStats struct{
	MovementSpeedModifier		float32
}

func (p *InternalInteractable) getClass(state *MatchState) *GameDB.Class {
	return state.GameDB.Classes[p.Classname]
}

//timer
func (p *InternalInteractable) startAutoattackTimer(endMainhand int64, endOffhand int64, targetId string){
	if endMainhand > 0 {
		p.AutoattackMainhandTickEnd = endMainhand
		p.Autoattacking = true
	}
	if endOffhand > 0 {
		p.AutoattackOffhandTickEnd = endOffhand
		p.Autoattacking = true
	}
	p.AutoattackTargeted = targetId
}
func (p *InternalInteractable) stopAutoattackTimer(){
	p.Autoattacking = false
	p.AutoattackMainhandTickEnd = 0
	p.AutoattackOffhandTickEnd = 0
	p.AutoattackTargeted = ""
}
func (p *InternalInteractable) startCastTimer(spellId int64, endTick int64, targetId interface{}){
	p.CastingSpellId = spellId
	p.CastingTickEnd = endTick
	p.CastingTargeted = targetId
}
func (p *InternalInteractable) stopCastTimer(){
	p.CastingSpellId = -1
	p.CastingTickEnd = 0
	p.CastingTargeted = ""
}


//movement
func (p *InternalInteractable) performMovement(state *MatchState, vector Vector2Df, rotation float32, movementSpeed float32) {
	p.Rotation = rotation;
	//clamb [-1..1]
	if vector.length() == 0 {
		return
	}
	if vector.length() > 1 {
		fmt.Printf("* * %v * *", vector.length())
		vector.X /= vector.length()
		vector.Y /= vector.length()
	}

	vector.X *= (movementSpeed / float32(100))
	vector.Y *= (movementSpeed / float32(100))

	mod := p.StatModifiers.MovementSpeedModifier
	vector.X *= mod
	vector.Y *= mod

	moveMsgCount := p.MoveMessageCountThisFrame
	if(p.Type == PublicMatchState_Interactable_NPC)	{
		moveMsgCount = 1
	}

	add := Vector2Df {
		X: vector.X / float32(moveMsgCount) * ((20) / float32(state.TickRate)),
		Y: vector.Y / float32(moveMsgCount) * ((20) / float32(state.TickRate)),
	}
	
	if math.IsNaN(float64(add.X)) || math.IsNaN(float64(add.Y)) {
		return
	}

	if p.CastingSpellId > 0 && state.GameDB.Spells[p.CastingSpellId].InterruptedBy != GameDB.Interrupt_Type_None && (vector.X != 0 || vector.Y != 0) {
		fmt.Printf("cancelCast %v\n", p.CastingSpellId)
		p.cancelCast(state)
	}

	rotatedAdd := add.rotate(rotation)		

	p.Position.X += rotatedAdd.X;
	p.Position.Y += rotatedAdd.Y;
	p.Rotation = rotation;

	//am i still in my triangle?	
	if p.TriangleIndex >= 0 {
		isItIn, _, _, _ := state.Map.Triangles[p.TriangleIndex].IsInTriangle(p.Position.X, p.Position.Y)
		if !isItIn {
			p.TriangleIndex = -1
		}
	}

	//no current triangle_index?
	if p.TriangleIndex < 0 {
		//find triangle I am in
		found := false
		for i, triangle := range state.Map.Triangles {
			isItIn, _, _, _ := triangle.IsInTriangle(p.Position.X, p.Position.Y)
			if isItIn {
				p.TriangleIndex = int64(i)
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

func (p *InternalInteractable) rotateTowardsTarget(targetPos *Vector2Df) {
	p.Rotation = float32(math.Atan2(float64(targetPos.X), float64(targetPos.Y)) * 57.2957795131);
}

//pathfinding


//regen
func (p *InternalInteractable) regen(state *MatchState, hpPercent, powerPercent float64) {	
	thisClass := state.GameDB.Classes[p.Classname]

	if p.CurrentHealth <= 0 {
		return
	}

	p.CurrentHealth += float32(math.Max(float64(p.getHpRegen(thisClass)) * hpPercent, 0));
	p.CurrentPower += float32(math.Max(float64(p.getManaRegen(thisClass)) * powerPercent, 0));

	p.CurrentHealth = float32(math.Min(float64(p.CurrentHealth), float64(p.getMaxHp(thisClass))))
	p.CurrentPower = float32(math.Min(float64(p.CurrentPower), float64(p.getMaxMana(thisClass))))
}

//fight
func (p *InternalInteractable) applyAutoattackDamage(state *MatchState, creator string, slot GameDB.Item_Slot) {
	source :=  state.Player[creator]
	sourceClass := state.GameDB.Classes[source.Classname]
	
	thisClass := state.GameDB.Classes[p.Classname]
	if(!p.IsEngaged || p.Target == "") {
		p.Target = source.Id
	}
	p.IsEngaged = true

	miss := (1 - source.getMeeleHitChance(sourceClass))
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
	fmt.Printf("\napplyAutoattackDamage (%v) %v+%v to unit %v\n", slot, dmgInput, source.getMeeleAttackPower(sourceClass), p.Id)
	dmgInput += source.getMeeleAttackPower(sourceClass)

	if slot == GameDB.Item_Slot_Weapon_OffHand {
		miss *= 2
	}
	
	roll := randomPercentage()
	dodge := p.getDodgeChance(thisClass)
	parry := p.getParryChance(thisClass)
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
	
	armor := p.getArmor(thisClass)
	dmgInput -= armor

	block := p.getBlockPercentage(thisClass)
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
	crit := source.getMeeleCritChance(sourceClass)
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgBlocked = dmgBlocked * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - p.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;
	p.LastHealthDrainTick = state.PublicMatchState.Tick
	fmt.Printf("applyDamage to %v: %v -> now:  %v\n", p.Id, (dmgInput + dmgInputCrit) - overkill, p.CurrentHealth)

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

func (p *InternalInteractable) applyAbilityDamage(state *MatchState, effect *GameDB.Effect, creator string) {
	source :=  state.Player[creator]
	sourceClass := state.GameDB.Classes[p.Classname]
	
	thisClass := state.GameDB.Classes[p.Classname]
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
		miss = (1 - source.getSpellHitChance(sourceClass))
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

	resist := p.getResistance(thisClass, effect.School)
	dmgResisted := dmgInput * (resist / 100)
	dmgInput = dmgInput - dmgResisted
	fmt.Printf("magical reduction by resistance (%v) by %v -> %v\n", resist, dmgResisted, dmgInput)

	roll = randomPercentage()
	crit := source.getSpellCritChance(sourceClass)
	dmgInputCrit := float32(0)
	if roll <= crit {
		dmgResisted = dmgResisted * 2
		dmgInputCrit = dmgInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v from %v\n", roll, crit, p.Id, dmgInput, dmgInputCrit, effect)
		dmgInput = 0
	} 

	overkill := float32(0)
	overkill = (dmgInput + dmgInputCrit) - p.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.CurrentHealth -= (dmgInput + dmgInputCrit) - overkill;	
	p.LastHealthDrainTick = state.PublicMatchState.Tick
	fmt.Printf("applyDamage to %v: %v from %v  -> now:  %v\n\n", p.Id, (dmgInput + dmgInputCrit) - overkill, effect, p.CurrentHealth)

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


func (p *InternalInteractable) applyAreaDamage(state *MatchState, effect *GameDB.Effect, creator string) {
	source :=  state.Player[creator]
	
	thisClass := state.GameDB.Classes[p.Classname]
	if(!p.IsEngaged || p.Target == "") {
		p.Target = source.Id
	}
	p.IsEngaged = true

	dmgInput := randomFloatInt(effect.ValueMin, effect.ValueMax)
	fmt.Printf("\napplyAreaDamage %v from effect %v to unit %v\n", dmgInput, effect, p.Id)

	resist := p.getResistance(thisClass, effect.School)
	dmgResisted := dmgInput * (resist / 100)
	dmgInput = dmgInput - dmgResisted
	fmt.Printf("magical reduction by resistance (%v) by %v -> %v\n", resist, dmgResisted, dmgInput)

	overkill := float32(0)
	overkill = dmgInput - p.CurrentHealth
	if overkill <= 0 {
		overkill = 0
	}
	p.CurrentHealth -= dmgInput - overkill;	
	p.LastHealthDrainTick = state.PublicMatchState.Tick
	fmt.Printf("applyDamage to %v: %v from %v  -> now:  %v\n\n", p.Id, dmgInput - overkill, effect, p.CurrentHealth)

	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: creator,
		DestinationId: p.Id,
		SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
		Source: PublicMatchState_CombatLogEntry_AoE,
		Type: &PublicMatchState_CombatLogEntry_Damage{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Damage{
			Amount: dmgInput,
			Resisted: dmgResisted,
			Blocked: 0,
			Absorbed: 0,
			Critical: 0,
			Overkill: overkill,			
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
}

func (p *InternalInteractable) applyAbilityHeal(state *MatchState, effect *GameDB.Effect, creator string) {
	source :=  state.Player[creator]
	sourceClass := state.GameDB.Classes[p.Classname]

	healInput := randomFloatInt(effect.ValueMin, effect.ValueMax)
	fmt.Printf("\applyAbilityHeal %v from effect %v to unit %v\n", healInput, effect, p.Id)
	
	roll := randomPercentage()
	crit := source.getSpellCritChance(sourceClass)
	healInputCrit := float32(0)
	if roll <= crit {
		healInputCrit = healInput * 2
		fmt.Printf("crit (%v/%v) damage to %v: %v -> %v from %v\n", roll, crit, p.Id, healInput, healInputCrit, effect)
		healInput = 0
	} 

	overheal := float32(0)
	overheal = (healInput + healInputCrit) - (p.getMaxHp(p.getClass(state)) - p.CurrentHealth)
	if overheal <= 0 {
		overheal = 0
	}
	p.CurrentHealth += (healInput + healInputCrit) - overheal;	
	fmt.Printf("applyDamage to %v: %v from %v  -> now:  %v\n\n", p.Id, (healInput + healInputCrit) - overheal, effect, p.CurrentHealth)

	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: creator,
		DestinationId: p.Id,
		SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
		Source: PublicMatchState_CombatLogEntry_Spell,
		Type: &PublicMatchState_CombatLogEntry_Heal{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Heal{
			Amount: healInput,
			Absorbed: 0,
			Critical: healInputCrit,
			Overheal: overheal,			
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)
}

func (p *InternalInteractable) containsEffectId(id int64, creator string) int64 {
    for i, a := range p.Auras {
        if a.EffectId == id && a.Creator == creator{
            return int64(i)
        }
    }
    return -1
}

//autoattack
func (p *InternalInteractable) startAutoattack(state *MatchState, attacktype Client_Message_Client_Autoattack_Type) {
	if p.Autoattacking == true && p.AutoattackTargeted == p.Target {
		return
	}
	thisClass := state.GameDB.Classes[p.Classname]

	failedMessage := ""

	if p.CastingSpellId > 0 && state.GameDB.Spells[p.CastingSpellId].IgnoresWeaponswing == false {
		failedMessage = "Cannot do that now!"
		
	}

	targetId := ""
	if p.Target == "" || p.Target == "Player" || p.Target == p.Id {
		failedMessage = "No Target!"
	} else {		
		targetId = p.Target
		target := state.Player[targetId]
		distance := p.Position.distance(target.Position)	

		if target.Team == p.Team {
			failedMessage = "Not an Enemy!"
		}	
		
		if distance > thisClass.Mainhand.Range {
			failedMessage = "Out of Range!"
		}

		if distance > 1 {
			if IntersectingBorders(p.Position, target.Position, state.Map) {
				failedMessage = "Not in Line of Sight!"
			}
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

		p.stopAutoattackTimer()

		return
	}
	mhEnd := int64(0)
	if (thisClass.Mainhand != nil && thisClass.Mainhand.AttackSpeed > 0) {
		mhEnd = int64(thisClass.Mainhand.AttackSpeed * p.getMeeleAttackSpeed(thisClass) * float32(state.TickRate)) + state.PublicMatchState.Tick
		
		fmt.Printf("swinging mainhand at  %v\n", mhEnd)
	}
	ohEnd := int64(0)
	if (thisClass.Offhand != nil && thisClass.Offhand.AttackSpeed > 0) {
		ohEnd = int64(thisClass.Offhand.AttackSpeed * p.getMeeleAttackSpeed(thisClass) * float32(state.TickRate)) + state.PublicMatchState.Tick
	
		fmt.Printf("swinging offhand at  %v\n", ohEnd)
	}
	p.startAutoattackTimer(mhEnd, ohEnd, targetId)
}

func (p *InternalInteractable) finishAutoattack(state *MatchState, slot GameDB.Item_Slot, targetId string) {
	target := state.Player[targetId]
	distance := p.Position.distance(target.Position)	

	if distance > state.GameDB.Classes[p.Classname].Mainhand.Range {
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

		p.stopAutoattackTimer()

		return
	}
	target.applyAutoattackDamage(state, p.Id, slot)
}

//casts
func (p *InternalInteractable) startCast(state *MatchState, spell *GameDB.Spell, targetPos *Vector2Df) {		
	fmt.Printf("startCast: %v %v\n", spell.Id, spell.Name)
	thisClass := state.GameDB.Classes[p.Classname]
	failedMessage := ""

	if (p.GlobalCooldown > 0 || spell.IgnoresGCD == true) || p.CastingSpellId > 0 || p.Cooldowns[spell.Id] > state.PublicMatchState.Tick {
		failedMessage = "Cannot do that now!"
	}
	
	manacost := float32(spell.BaseCost) + float32(spell.CostPercentage) * p.getMaxMana(thisClass);
	if manacost > p.CurrentPower {
		failedMessage = "Not enough Mana!"
	}

	targetId := "" 
	if (spell.Target_Type == GameDB.Spell_Target_Type_Self || p.Target == "Player" || p.Target == p.Id) && spell.Target_Type != GameDB.Spell_Target_Type_AoE{
		targetId = p.Id
		if spell.Target_Type != GameDB.Spell_Target_Type_Ally {
			failedMessage = "Not an Ally!"
		}
	} else if (spell.Target_Type == GameDB.Spell_Target_Type_None) {
			//nothing so far
	} else {
		if p.Target == "" && spell.Target_Type != GameDB.Spell_Target_Type_AoE {
			failedMessage = "No Target!"
		} else {
			targetId = ""
			target := p
			distance := float32(999999)	
			if spell.Target_Type == GameDB.Spell_Target_Type_AoE {
				distance = p.Position.distance(targetPos)	
			} else {
				targetId = p.Target
				target = state.Player[targetId]	
				distance = p.Position.distance(target.Position)	
					
				switch spell.Target_Type {
				case GameDB.Spell_Target_Type_Ally:
					if target.Team != p.Team {
						failedMessage = "Not an Ally!"
					}
				case GameDB.Spell_Target_Type_Enemy:
					if target.Team == p.Team {
						failedMessage = "Not an Enemy!"
					}
				}
			}
			
			
			behind := target.Position.isFacedBy(p.Position, p.Rotation)
			if spell.FacingFront && !behind && failedMessage == "" {
				failedMessage = "Is not facing target!"
			}
	
			if distance > spell.Range && failedMessage == "" {	
				fmt.Printf("Out of Range: %v > %v\n", distance, spell.Range)
				failedMessage = "Out of Range!"
			}
	
			if IntersectingBorders(p.Position, target.Position, state.Map) && failedMessage == "" {
				failedMessage = "Not in Line of Sight!"
			}
		}

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
		p.stopAutoattackTimer()
	}
	
	p.LastPowerDrainTick = state.PublicMatchState.Tick
	if spell.CastTime == 0 {
		switch spell.Target_Type {
		case GameDB.Spell_Target_Type_None:
			p.finishCast(state, spell, nil)
		case GameDB.Spell_Target_Type_AoE:
			p.finishCast(state, spell, targetPos)
		default:
			p.finishCast(state, spell, targetId)
		}
	} else {
		end := int64(spell.CastTime * (p.getSpellAttackSpeed(thisClass)) * float32(state.TickRate)) + state.PublicMatchState.Tick

		switch spell.Target_Type {
		case GameDB.Spell_Target_Type_None:
			p.startCastTimer(spell.Id, end, nil)
		case GameDB.Spell_Target_Type_AoE:
			p.startCastTimer(spell.Id, end, targetPos)
		default:
			p.startCastTimer(spell.Id, end, targetId)
		}
	}
}

func (p *InternalInteractable) cancelCast(state *MatchState) {	
	clEntry := &PublicMatchState_CombatLogEntry {
		Timestamp: state.PublicMatchState.Tick,
		SourceId: p.Id,
		SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceSpellId{p.CastingSpellId},
		Source: PublicMatchState_CombatLogEntry_Spell,
		Type: &PublicMatchState_CombatLogEntry_Cast{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Cast{
			Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Cast_Failed,
			FailedMessage: "Cast canceled by Movement!",
		}},
	}
	state.PublicMatchState.Combatlog = append(state.PublicMatchState.Combatlog, clEntry)

	p.stopCastTimer()
}

func (p *InternalInteractable) finishCast(state *MatchState, spell *GameDB.Spell, target interface{}) {
	thisClass := state.GameDB.Classes[p.Classname]
	
	p.Cooldowns[spell.Id] = state.PublicMatchState.Tick + int64(spell.Cooldown * float32(state.TickRate))

	var pos *Vector2Df
	
	if spell.Target_Type == GameDB.Spell_Target_Type_AoE {
		pos = target.(*Vector2Df)
	} else {
		pos = state.Player[target.(string)].Position
	}
	
	if !spell.IgnoresGCD {
		p.GlobalCooldown = spell.GlobalCooldown
	}

	if !IntersectingBorders(p.Position, pos, state.Map) {
		p.CurrentPower -= float32(spell.BaseCost);
		p.CurrentPower -= float32(spell.CostPercentage) * p.getMaxMana(thisClass);
		p.LastPowerDrainTick = state.PublicMatchState.Tick	

		fmt.Printf("finish cast spell: %v (mana now: %v)\n", spell.Name, p.CurrentPower)

		switch spell.Application_Type {
		case GameDB.Spell_Application_Type_WeaponSwing:
		case GameDB.Spell_Application_Type_Instant:
			SpellHit(state, state.Player[target.(string)], p.Id, spell)
		case GameDB.Spell_Application_Type_Missile:	
			proj := &PublicMatchState_Projectile{
				Id: "p_" + strconv.FormatInt(state.ProjectileCounter, 16),
				SpellId: spell.Id,
				Position: &Vector2Df {
					X: p.Position.X,
					Y: p.Position.Y,
				},
				Rotation: p.Rotation,
				CreatedAtTick: state.PublicMatchState.Tick,
				Target: target.(string),
				Speed: spell.Speed,
				Creator: p.Id,
			}
			state.PublicMatchState.Projectile[proj.Id] = proj
			state.ProjectileCounter++	
		case GameDB.Spell_Application_Type_Beam:
			SpellHit(state, state.Player[target.(string)], p.Id, spell)
		case GameDB.Spell_Application_Type_AoE:
			CreateAoEArea(state, spell, pos, p.Id)
		case GameDB.Spell_Application_Type_Cone:
		case GameDB.Spell_Application_Type_Summon:	
		case GameDB.Spell_Application_Type_Teleport:			
		}
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