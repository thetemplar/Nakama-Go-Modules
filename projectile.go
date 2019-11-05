package main

import (
)

func (p PublicMatchState_Projectile) Run(state *MatchState, projectile *PublicMatchState_Projectile, tickrate int) {
	target := state.Player[projectile.Target]					
	distance := projectile.Position.distance(target.Position)	
	direction := Vector2Df {
		X: target.Position.X - projectile.Position.X,
		Y: target.Position.Y - projectile.Position.Y,
	}
	direction.X /= distance
	direction.Y /= distance

	if distance <= (projectile.Speed / float32(tickrate)) {
		//impact
		spell := state.GameDB.Spells[projectile.SpellId]	
		SpellHit(state, target, projectile.Creator, spell)
		delete(state.PublicMatchState.Projectile, projectile.Id)
		projectile = nil
		return
	}

	projectile.Position.X = projectile.Position.X + (direction.X * (projectile.Speed / float32(tickrate)))
	projectile.Position.Y = projectile.Position.Y + (direction.Y * (projectile.Speed / float32(tickrate)))
}
