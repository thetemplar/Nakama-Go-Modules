package GameDB

type Database struct {
	Spells				map[int64]*Spell
	Effects				map[int64]*Effect
	Procs				map[int64]*Proc
	Items				map[int64]*Item

	Classes				map[string]*Class
}

func (g *Database) searchSpellByName(name string) (*Spell) {
	for _, spell := range g.Spells { 
		if spell.Name == name {
			return spell;
		}
	}
	return nil
}

func (g *Database) searchEffectByName(name string) (*Effect) {
	for _, effect := range g.Effects { 
		if effect.Name == name {
			return effect;
		}
	}
	return nil
}

func (g *Database) searchItemByName(name string) (*Item) {
	for _, item := range g.Items { 
		if item.Name == name {
			return item;
		}
	}
	return nil
}

func (g *Database) searchProcByName(name string) (*Proc) {
	for _, proc := range g.Procs { 
		if proc.Name == name {
			return proc;
		}
	}
	return nil
}