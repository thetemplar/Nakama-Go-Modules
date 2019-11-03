package main

import (
	"github.com/heroiclabs/nakama-common/runtime"
)

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
	CastingTargeted 			string

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
	Cooldowns     				map[string]*int64
}

type Act func(state *MatchState, p *InternalInteractable)

type PlayerStats struct{
	MovementSpeedModifier		float32
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
func (p *InternalInteractable) startCastTimer(spellId int64, endTick int64, targetId string){
	p.CastingSpellId = spellId
	p.CastingTickEnd = endTick
	p.CastingTargeted = targetId
}
func (p *InternalInteractable) stopCastTimer(){
	p.CastingSpellId = -1
	p.CastingTickEnd = 0
	p.CastingTargeted = ""
}
