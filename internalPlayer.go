package main

import (
	"github.com/heroiclabs/nakama-common/runtime"
)

type InternalPlayer struct {
	Presence                	runtime.Presence
	Id                      	string

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
	BasePlayerStats				PlayerStats
	StatModifiers				PlayerStats

	//regen
	LastRegenTick				int64
	LastHealthDrainTick			int64
	LastPowerDrainTick			int64

	//fights
	Act Act
}

type Act func(state *MatchState, p *InternalPlayer)

type PlayerStats struct{
	MovementSpeed				float32
}

func (p *InternalPlayer) getPublicPlayer(state *MatchState) (*PublicMatchState_Interactable) {
	return state.PublicMatchState.Interactable[p.Id];
}

//timer
func (p *InternalPlayer) startAutoattackTimer(endMainhand int64, endOffhand int64, targetId string){
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
func (p *InternalPlayer) stopAutoattackTimer(){
	p.Autoattacking = false
	p.AutoattackMainhandTickEnd = 0
	p.AutoattackOffhandTickEnd = 0
	p.AutoattackTargeted = ""
}
func (p *InternalPlayer) startCastTimer(spellId int64, endTick int64, targetId string){
	p.CastingSpellId = spellId
	p.CastingTickEnd = endTick
	p.CastingTargeted = targetId
}
func (p *InternalPlayer) stopCastTimer(){
	p.CastingSpellId = -1
	p.CastingTickEnd = 0
	p.CastingTargeted = ""
}
