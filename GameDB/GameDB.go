package main

type GameDB struct {
	Spells				map[int64]*GameDB_Spell
	Effects				map[int64]*GameDB_Effect
	Procs				map[int64]*GameDB_Proc
	Items				map[int64]*GameDB_Item

	Classes				map[string]*GameDB_Class
}

func (g *GameDB) searchSpellsByName(name string) (*GameDB_Spell) {
	for _, spell := range g.Spells { 
		if spell.Name == name {
			return spell;
		}
	}
	return nil
}

func (g *GameDB) searchEffectByName(name string) (*GameDB_Effect) {
	for _, effect := range g.Effects { 
		if effect.Name == name {
			return effect;
		}
	}
	return nil
}