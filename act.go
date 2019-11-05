package main

import (
	"fmt"
)

func ActGeneric(state *MatchState, p *InternalInteractable) {
	fmt.Printf("Act %v -> Target: %v\n", p.Id, p.Target)
}

type NpcState_Ogre struct {
	enraged 	bool 
}

func Act_Ogre(state *MatchState, self *InternalInteractable) {
	target := state.Player[self.Target]

	
	direction := Vector2Df {
		X: target.Position.X - self.Position.X,
		Y: target.Position.Y - self.Position.Y,
	}
	if direction.length() > float32(3) {		
		self.rotateTowardsTarget(target.Position)
		self.performMovement(state, direction, 0, self.getClass(state).MovementSpeed)
	} else {		
		self.startAutoattack(state, Client_Message_Client_Autoattack_Meele)
	}

	if self.CurrentHealth < 100 && !self.npcState.(*NpcState_Ogre).enraged {
		enrage := state.GameDB.SearchSpellByName("Enrage");
		self.startCast(state, enrage, &Vector2Df{})
		
		self.npcState.(*NpcState_Ogre).enraged = true
	}
}