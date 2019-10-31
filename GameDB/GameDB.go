package GameDB

type Database struct {
	Spells				map[int64]*Spell
	Effects				map[int64]*Effect
	Procs				map[int64]*Proc
	Items				map[int64]*Item

	Classes				map[string]*Class
}

func (g *Database) SearchSpellByName(name string) (*Spell) {
	for _, spell := range g.Spells { 
		if spell.Name == name {
			return spell;
		}
	}
	return nil
}

func (g *Database) SearchEffectByName(name string) (*Effect) {
	for _, effect := range g.Effects { 
		if effect.Name == name {
			return effect;
		}
	}
	return nil
}

func (g *Database) SearchItemByName(name string) (*Item) {
	for _, item := range g.Items { 
		if item.Name == name {
			return item;
		}
	}
	return nil
}

func (g *Database) SearchProcByName(name string) (*Proc) {
	for _, proc := range g.Procs { 
		if proc.Name == name {
			return proc;
		}
	}
	return nil
}