package main

import (
	"math"
	"fmt"
)

func ActGeneric(state *MatchState, p *InternalPlayer) {
	fmt.Printf("Act %v -> Target: %v\n", p.Id, p.getPublicPlayer(state).Target)
}

func Act_Ogre(state *MatchState, p *InternalPlayer) {
	self := p.getPublicPlayer(state)
	target := state.PublicMatchState.Interactable[self.Target]
	
	direction := PublicMatchState_Vector2Df {
		X: target.Position.X - self.Position.X,
		Y: target.Position.Y - self.Position.Y,
	}
	length := float32(math.Sqrt(math.Pow(float64(direction.X), 2) + math.Pow(float64(direction.Y), 2)))
	if(length > 2) {		
		self.RotateTowardsTogarget(state, target.Position)
		self.PerformMovement(state, direction.X, direction.Y, 0)
	}

	if(self.Character.CurrentHealth < 100) {
		enrage := state.GameDB.SearchSpellByName("Enrage");
		self.startCast(state, enrage)
	}
}