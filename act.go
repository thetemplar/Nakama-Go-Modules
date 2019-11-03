package main

import (
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
	if(direction.length() > float32(4)) {		
		self.rotateTowardsTarget(target.Position)
		self.performMovement(state, direction, 0, self.getClass(state).MovementSpeed)
	}

	if(self.CurrentHealth < 100) {
		enrage := state.GameDB.SearchSpellByName("Enrage");
		self.startCast(state, enrage)
	}
}