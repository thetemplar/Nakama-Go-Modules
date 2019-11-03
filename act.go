package main

import (
	"math"
	"fmt"
)

func ActGeneric(state *MatchState, p *InternalInteractable) {
	fmt.Printf("Act %v -> Target: %v\n", p.Id, p.Target)
}

func Act_Ogre(state *MatchState, self *InternalInteractable) {
	target := state.Player[self.Target]
	
	direction := PublicMatchState_Vector2Df {
		X: target.Position.X - self.Position.X,
		Y: target.Position.Y - self.Position.Y,
	}
	length := float32(math.Sqrt(math.Pow(float64(direction.X), 2) + math.Pow(float64(direction.Y), 2)))
	if(length > 2) {		
		self.rotateTowardsTarget(target.Position)
		self.performMovement(state, direction.X, direction.Y, 0)
	}

	if(self.CurrentHealth < 100) {
		enrage := state.GameDB.SearchSpellByName("Enrage");
		self.startCast(state, enrage)
	}
}