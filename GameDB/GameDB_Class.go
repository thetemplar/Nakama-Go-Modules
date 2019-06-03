package main

type GameDB_Class struct {
	Name 				string
	
	Spellbook			[]*GameDB_Spell
	Items				[]*GameDB_Item
	Procs				[]*GameDB_Proc
}