package main

import (
	"fmt"
)

func ActGeneric(state *MatchState, p *InternalInteractable) {
	fmt.Printf("Act %v -> Target: %v\n", p.Id, p.Target)
}

type NpcState_Ogre struct {
	enraged 	bool
	lastSpell 	int64
}

func Act_Ogre(state *MatchState, self *InternalInteractable) {
	target := state.Player[self.Target]

	if self.npcState.(*NpcState_Ogre).lastSpell == 0 {
		self.npcState.(*NpcState_Ogre).lastSpell = state.PublicMatchState.Tick
	}

	if self.npcState.(*NpcState_Ogre).lastSpell + int64(15 * state.TickRate) < state.PublicMatchState.Tick {
		fire := state.GameDB.SearchSpellByName("Fire Zone");
		self.startCast(state, fire, target.Position)
	}
	
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