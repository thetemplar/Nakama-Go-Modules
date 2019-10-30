package main

import (
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
	self.getInternalPlayer(state).MoveMessageCountThisFrame = 1
	fmt.Printf("[OGRE] Act %v (%v|%v) -> Target: %v (%v|%v) = dir: (%v|%v)\n", self.Id, self.Position.X, self.Position.Y, target.Id, target.Position.X, target.Position.Y, direction.X, direction.Y)
	
	self.PerformMovement(state, direction.X, direction.Y, 0)
}