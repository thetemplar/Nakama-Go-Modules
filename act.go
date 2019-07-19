package main

import (
	"fmt"
)

func ActGeneric(state *MatchState, p *InternalPlayer) {
	fmt.Printf("Act %v -> Target: %v\n", p.Id, p.getPublicPlayer(state).Target)
}

func Act_Ogre(state *MatchState, p *InternalPlayer) {
	fmt.Printf("[OGRE] Act %v -> Target: %v\n", p.Id, p.getPublicPlayer(state).Target)
}