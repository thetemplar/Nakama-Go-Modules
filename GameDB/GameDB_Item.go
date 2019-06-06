package main

type GameDB_Item struct {
	Id					int64
	Name 				string
	Description 		string

	Type				GameDB_Item_Type
	Slot				GameDB_Item_Slot

	DamageMin			int32
	DamageMax			int32
	AttackSpeed			float32
	Range 				float32

	BlockValue			int32
}

type GameDB_Item_Type int
const (
	GameDB_Item_Type_Weapon_OneHand = 0;
	GameDB_Item_Type_Weapon_MainHand = 1;
	GameDB_Item_Type_Weapon_OffHand = 2;
	GameDB_Item_Type_Weapon_TwoHand = 3;
	GameDB_Item_Type_Weapon_Shield = 4;
	GameDB_Item_Type_Weapon_Bow = 5;
	GameDB_Item_Type_Weapon_Wand = 6;
)

type GameDB_Item_Slot int
const (
	GameDB_Item_Slot_Weapon_MainHand = 0;
	GameDB_Item_Slot_Weapon_OffHand = 1;
	GameDB_Item_Slot_Weapon_BothHands = 2;
)
