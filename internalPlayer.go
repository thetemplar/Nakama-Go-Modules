package main

import "github.com/heroiclabs/nakama/runtime"

type InternalPlayer struct {
	Presence                runtime.Presence
	Id                      string

	//messages from client
	LastMessage             runtime.MatchData
	LastMessageServerTick   int64
	LastMessageClientTick   int64
	MissingCount			int
	MessageCountThisFrame   int

	//movement	
	TriangleIndex 			int64
	
	//casts
	CastingSpellId 			int64
	CastingTickStarted 		int64
	CastingTargeted 		string

	//playerstats
	BasePlayerStats			PlayerStats
	StatModifiers			PlayerStats
}

type PlayerStats struct{
	MovementSpeed			float32
}

func (p InternalPlayer) getPublicPlayer(state *MatchState) (*PublicMatchState_Interactable) {
	return state.PublicMatchState.Interactable[p.Id];
}