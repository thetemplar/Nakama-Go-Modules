package GameDB

type Item struct {
	Id					int64
	Name 				string
	Description 		string

	Type				Item_Type
	Slot				Item_Slot

	DamageMin			int32
	DamageMax			int32
	AttackSpeed			float32
	Range 				float32

	BlockValue			int32
}

type Item_Type int
const (
	Item_Type_Weapon_OneHand = 0;
	Item_Type_Weapon_MainHand = 1;
	Item_Type_Weapon_OffHand = 2;
	Item_Type_Weapon_TwoHand = 3;
	Item_Type_Weapon_Shield = 4;
	Item_Type_Weapon_Bow = 5;
	Item_Type_Weapon_Wand = 6;
)

type Item_Slot int
const (
	Item_Slot_Weapon_MainHand = 0;
	Item_Slot_Weapon_OffHand = 1;
	Item_Slot_Weapon_BothHands = 2;
)
