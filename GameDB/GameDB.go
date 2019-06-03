package main

type GameDB struct {
	Spells				map[int64]*GameDB_Spell
	Effects				map[int64]*GameDB_Effect
	Procs				map[int64]*GameDB_Proc
	Items				map[int64]*GameDB_Item

	Classes				map[string]*GameDB_Class
}

func (g *GameDB) searchSpellByName(name string) (*GameDB_Spell) {
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

func (g *GameDB) searchItemByName(name string) (*GameDB_Item) {
	for _, item := range g.Items { 
		if item.Name == name {
			return item;
		}
	}
	return nil
}

func (g *GameDB) searchProcByName(name string) (*GameDB_Proc) {
	for _, proc := range g.Procs { 
		if proc.Name == name {
			return proc;
		}
	}
	return nil
}