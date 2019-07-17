package main

import (
	"context"
	"time"
	"database/sql"
	"github.com/heroiclabs/nakama/runtime"
	"github.com/golang/protobuf/proto"
	"fmt"
	"strconv"
	"math"
)

type MatchState struct {
	PublicMatchState    PublicMatchState
	EmptyCounter        int
	Debug               bool
	TickRate			int

	InternalPlayer      map[string]*InternalPlayer
	PresenceList      	map[string]*runtime.Presence

	ProjectileCounter	int64
	NpcCounter			int64
	
	GameDB				*GameDB
	Map					*Map
	
	runtimeSet			[]int64
	runtimeSetIndex		int
}  

func (ms *MatchState) GetClassFromDB(char *Character) *GameDB_Class {
	return ms.GameDB.Classes[char.Classname]
}

type Match struct{
}


func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	logger.Print(" >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> MatchInit <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	for _, entry := range params { 
		logger.Printf("%+v\n", entry)
	}
	
	tickRate := 10
	label := ""

	state := &MatchState{
		Debug: false,
		EmptyCounter : 0,
		PublicMatchState : PublicMatchState{
			Interactable: make(map[string]*PublicMatchState_Interactable),
			Projectile: make(map[string]*PublicMatchState_Projectile),
			Combatlog: make([]*PublicMatchState_CombatLogEntry, 0),
		},

		
		InternalPlayer: make(map[string]*InternalPlayer),
		PresenceList: make(map[string]*runtime.Presence),
		runtimeSet: make([]int64, 20),
		TickRate: tickRate,
	}
	
	//create spellbook
	state.GameDB = init_db()
	state.Map = map_init()

	//create map npcs:
	enemy := &PublicMatchState_Interactable{
		Id: "npc_" + strconv.FormatInt(state.NpcCounter, 16),
		Type: PublicMatchState_Interactable_NPC,
		CharacterId: 2,
		//Position: currentPlayerPublic.Position,
		Position: &PublicMatchState_Vector2Df {
			X: 15,
			Y: 15,
		},
		Auras: make([]*PublicMatchState_Aura, 0),
		Character: &Character {
			Classname: "Mage",
			EquippedItemMainhandId: 1,
			EquippedItemOffhandId: 2,
			CurrentHealth: 100,
			CurrentPower: 10,
		},
	}
	state.PublicMatchState.Interactable[enemy.Id] = enemy
	enemyInternal := &InternalPlayer{
		Id: "npc_" + strconv.FormatInt(state.NpcCounter, 16),
		Presence: nil,
		BasePlayerStats: PlayerStats {
			MovementSpeed: 20.0,
		},
		StatModifiers: PlayerStats {},
	}
	state.InternalPlayer[enemyInternal.Id] = enemyInternal
	state.NpcCounter++

	if state.Debug {
		logger.Printf("match init, starting with debug: %v at tickrate %v", state.Debug, tickRate)
	} 
	return state, tickRate, label
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	if state.(*MatchState).Debug {
		logger.Printf("match join attempt username %v user_id %v session_id %v node %v with metadata %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId(), metadata)
	}
	return state, true, ""
}

func SpawnPlayer(state *MatchState, userId string, classname string) {
	if state.PublicMatchState.Interactable[userId] != nil || state.InternalPlayer[userId] != nil || state.GameDB.Classes[classname] == nil {
		return
	}

	state.PublicMatchState.Interactable[userId] = &PublicMatchState_Interactable{
		Id: userId,
		Type: PublicMatchState_Interactable_Player,
		Position: &PublicMatchState_Vector2Df { 
			X: 0.1,
			Y: 0.1,
		},
		Character: &Character {
			Classname: classname,
			EquippedItemMainhandId: 1,
			EquippedItemOffhandId: 2,
			CurrentHealth: 100,
			CurrentPower: 100,
		},
	}
	
	state.InternalPlayer[userId] = &InternalPlayer{
		Id: userId,
		Presence: *state.PresenceList[userId],
		BasePlayerStats: PlayerStats {
			MovementSpeed: 20.0,
		},			
		TriangleIndex: -1,
		StatModifiers: PlayerStats {},
	}

	fmt.Printf("new character %v spawn @ %v\n", userId, state.PublicMatchState.Interactable[userId].Position)
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	for _, presence := range presences {		
		logger.Printf("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
		state.(*MatchState).PresenceList[presence.GetUserId()] = &presence
	}

	return state
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	for _, presence := range presences {		
		state.(*MatchState).PublicMatchState.Interactable[presence.GetUserId()] = nil
		state.(*MatchState).InternalPlayer[presence.GetUserId()] = nil
		state.(*MatchState).PresenceList[presence.GetUserId()] = nil
		delete(state.(*MatchState).PublicMatchState.Interactable, presence.GetUserId())
		delete(state.(*MatchState).InternalPlayer, presence.GetUserId())
		delete(state.(*MatchState).PresenceList, presence.GetUserId())

		logger.Printf("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
	}

	return state
}

func PerformMovement(logger runtime.Logger, state interface{}, playerId string, xAxis, yAxis, rotation float32, tickrate int) {


	currentPlayerInternal := state.(*MatchState).InternalPlayer[playerId];
	currentPlayerPublic   := state.(*MatchState).PublicMatchState.Interactable[playerId];
	

	add := PublicMatchState_Vector2Df {
		X: xAxis / float32(currentPlayerInternal.MessageCountThisFrame) * ((currentPlayerInternal.BasePlayerStats.MovementSpeed - currentPlayerInternal.StatModifiers.MovementSpeed) / float32(tickrate)),
		Y: yAxis / float32(currentPlayerInternal.MessageCountThisFrame) * ((currentPlayerInternal.BasePlayerStats.MovementSpeed - currentPlayerInternal.StatModifiers.MovementSpeed) / float32(tickrate)),
	}
	
	if math.IsNaN(float64(add.X)) || math.IsNaN(float64(add.Y)) {
		return
	}

	if currentPlayerInternal.CastingSpellId > 0 && state.(*MatchState).GameDB.Spells[currentPlayerInternal.CastingSpellId].InterruptedBy != GameDB_Interrupt_Type_None && (xAxis != 0 || yAxis != 0) {
		fmt.Printf("cancelCast %v\n", currentPlayerInternal.CastingSpellId)
		currentPlayerPublic.cancelCast(state.(*MatchState))
	}

	rotatedAdd := add.rotate(rotation)
		

	currentPlayerPublic.Position.X += rotatedAdd.X;
	currentPlayerPublic.Position.Y += rotatedAdd.Y;
	currentPlayerPublic.Rotation = rotation;
	

	//am i still in my triangle?	
	if currentPlayerInternal.TriangleIndex >= 0 {
		isItIn, _, _, _ := state.(*MatchState).Map.Triangles[currentPlayerInternal.TriangleIndex].isInTriangle(currentPlayerPublic.Position)
		if !isItIn {
			currentPlayerInternal.TriangleIndex = -1
		}
	}

	//no current triangle_index?
	if currentPlayerInternal.TriangleIndex < 0 {
		//find triangle I am in
		found := false
		for i, triangle := range state.(*MatchState).Map.Triangles {
			isItIn, _, _, _ := triangle.isInTriangle(currentPlayerPublic.Position)
			if isItIn {
				currentPlayerInternal.TriangleIndex = int64(i)
				found = true
				break;
			}
		}
		if !found {
			currentPlayerPublic.Position.X -= rotatedAdd.X;
			currentPlayerPublic.Position.Y -= rotatedAdd.Y;
		} 
	}	
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	if state.(*MatchState).Debug {
		logger.Printf("match loop match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Printf("match loop match_id %v message count %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), len(messages))
	}
    start := time.Now()
	state.(*MatchState).PublicMatchState.Tick = tick	
	state.(*MatchState).PublicMatchState.Combatlog = make([]*PublicMatchState_CombatLogEntry, 0)
	tickrate := ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int);

	//clear for new loop (finish cast & substract gcd)
	for _, player := range state.(*MatchState).InternalPlayer {	
		currentPlayerPublic := player.getPublicPlayer(state.(*MatchState));
		
		//finish casts
		if player.CastingTickEnd <= tick && player.CastingSpellId > 0 {
			currentPlayerPublic.finishCast(state.(*MatchState), state.(*MatchState).GameDB.Spells[player.CastingSpellId], player.CastingTargeted)
			player.stopCastTimer()
		}
		
		//finish swing mainhand 
		if player.Autoattacking && player.AutoattackMainhandTickEnd <= tick && player.AutoattackMainhandTickEnd > 0{
			currentPlayerPublic.finishAutoattack(state.(*MatchState), GameDB_Item_Slot_Weapon_MainHand, player.AutoattackTargeted)
			
			//queue next swing!
			player.AutoattackMainhandTickEnd = int64(state.(*MatchState).GameDB.Items[currentPlayerPublic.Character.EquippedItemMainhandId].AttackSpeed * state.(*MatchState).GetClassFromDB(currentPlayerPublic.Character).getMeeleAttackSpeed(currentPlayerPublic.Character) * float32(tickrate)) + tick
		}
		//finish swing offhand
		if player.Autoattacking && player.AutoattackOffhandTickEnd <= tick && player.AutoattackOffhandTickEnd > 0{
			currentPlayerPublic.finishAutoattack(state.(*MatchState), GameDB_Item_Slot_Weapon_OffHand, player.AutoattackTargeted)
			
			//queue next swing!
			player.AutoattackOffhandTickEnd = int64(state.(*MatchState).GameDB.Items[currentPlayerPublic.Character.EquippedItemOffhandId].AttackSpeed * state.(*MatchState).GetClassFromDB(currentPlayerPublic.Character).getMeeleAttackSpeed(currentPlayerPublic.Character) * float32(tickrate)) + tick
		}
			
		//substract GCD
		player.getPublicPlayer(state.(*MatchState)).GlobalCooldown -= float32(1)/float32(ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int));
	}

	//get new input-counts
	for _, message := range messages { 
		if state.(*MatchState).InternalPlayer[message.GetUserId()] == nil {
			continue
		}
		if(message.GetOpCode() == 0) {
			state.(*MatchState).InternalPlayer[message.GetUserId()].MessageCountThisFrame++
		}
	}

	//get new inputs
	for _, message := range messages { 
		if (state.(*MatchState).InternalPlayer[message.GetUserId()] == nil && message.GetOpCode() != 100) || message.GetOpCode() == 255 {
			continue
		}
		//logger.Printf("message from %v with opcode %v", message.GetUserId(), message.GetOpCode())
		//entry.UserID, entry.SessionId, entry.Username, entry.Node, entry.OpCode, entry.Data, entry.ReceiveTime
		currentPlayerInternal := state.(*MatchState).InternalPlayer[message.GetUserId()];
		currentPlayerPublic   := state.(*MatchState).PublicMatchState.Interactable[message.GetUserId()];
		if message.GetOpCode() == 0 {
			/*if state.(*MatchState).InternalPlayer[message.GetUserId()] == nil || state.(*MatchState).PublicMatchState.Interactable[message.GetUserId()] == nil {
				return
			}*/
			currentPlayerInternal.LastMessage = message
			currentPlayerInternal.LastMessageServerTick = tick
			currentPlayerInternal.MissingCount = 0

			msg := &Client_Character{}
			if err := proto.Unmarshal(message.GetData(), msg); err != nil {
				logger.Printf("Failed to parse incoming SendPackage Client_Character:", err)
			}
			
			if currentPlayerPublic.Target != msg.Target {
				currentPlayerInternal.stopAutoattackTimer();
			}
			currentPlayerPublic.Target = msg.Target;

			PerformMovement(logger, state, message.GetUserId(), msg.XAxis, msg.YAxis, msg.Rotation, tickrate)
			
			currentPlayerPublic.LastProcessedClientTick = msg.ClientTick
		} else if message.GetOpCode() == 1 {
			msg := &Client_Cast{}
			if err := proto.Unmarshal(message.GetData(), msg); err != nil {
				logger.Printf("Failed to parse incoming SendPackage Client_Cast:", err)
			}
			//is the spell in his spellbook?
			for _, spell := range state.(*MatchState).GetClassFromDB(currentPlayerPublic.Character).Spells {
				if (spell.Id == msg.SpellId) {
					currentPlayerPublic.startCast(state.(*MatchState), state.(*MatchState).GameDB.Spells[msg.SpellId])
					break
				}
			}
		} else if message.GetOpCode() == 2 {
			msg := &Client_Autoattack{}
			if err := proto.Unmarshal(message.GetData(), msg); err != nil {
				logger.Printf("Failed to parse incoming SendPackage Client_Autoattack:", err)
			}
			if !currentPlayerInternal.Autoattacking && currentPlayerPublic.Target != "" {
				fmt.Printf("startAutoattack %v > %v\n", currentPlayerPublic.Id, currentPlayerPublic.Target)
				currentPlayerPublic.startAutoattack(state.(*MatchState), msg.Attacktype)
			}
		} else if message.GetOpCode() == 3 {
			currentPlayerInternal.stopAutoattackTimer();
			currentPlayerInternal.stopCastTimer();
		} else if message.GetOpCode() == 100 {
			msg := &Client_SelectCharacter{}
			if err := proto.Unmarshal(message.GetData(), msg); err != nil {
				logger.Printf("Failed to parse incoming SendPackage Client_SelectCharacter:", err)
			}
			SpawnPlayer(state.(*MatchState), message.GetUserId(), msg.Classname)
		}  
	}
	
	//did a player not send an package? then re-do his last
	for _, player := range state.(*MatchState).InternalPlayer {		
		if player.LastMessageServerTick != tick {
			player.MissingCount++
			if player.MissingCount > 1 && player.LastMessage != nil {
				player.MessageCountThisFrame = 1
				logger.Printf("2nd missing Package from player %v in a row, inserting last known package.", player.Id)
	
				msg := &Client_Character{}
				if err := proto.Unmarshal(player.LastMessage.GetData(), msg); err != nil {
					logger.Printf("Failed to parse incoming SendPackage Client_Character:", err)
				}
	
				PerformMovement(logger, state, player.LastMessage.GetUserId(), msg.XAxis, msg.YAxis, msg.Rotation, tickrate)
			}
		}
	}

	//auras
	for _, interactable := range state.(*MatchState).PublicMatchState.Interactable {		
		i := 0
		doRecalc := false
		for _, aura := range interactable.Auras {
			effect := state.(*MatchState).GameDB.Effects[aura.EffectId]

			switch effect.Type.(type) {
			case *GameDB_Effect_Apply_Aura_Periodic_Damage:
				if int64(float32(aura.AuraTickCount + 1) * effect.Type.(*GameDB_Effect_Apply_Aura_Periodic_Damage).Intervall * float32(tickrate)) + aura.CreatedAtTick < tick {
					aura.AuraTickCount++					
					interactable.applyAbilityDamage(state.(*MatchState), effect, aura.Creator)
				}
			}			

			//is it depleted?
			if int64(effect.Duration * float32(tickrate)) + aura.CreatedAtTick < tick {
				clEntry := &PublicMatchState_CombatLogEntry {
					Timestamp: tick,
					SourceId: aura.Creator,
					DestinationId: interactable.Id,
					SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{aura.EffectId},
					Source: PublicMatchState_CombatLogEntry_Spell,
					Type: &PublicMatchState_CombatLogEntry_Aura{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Aura{
						Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Aura_Depleted,
					}},
				}
				state.(*MatchState).PublicMatchState.Combatlog = append(state.(*MatchState).PublicMatchState.Combatlog, clEntry)

				fmt.Printf("auras run off > %v\n", aura)

				switch effect.Type.(type) {
				case *GameDB_Effect_Apply_Aura_Mod:
					doRecalc = true //cant do "recalcStats" here, since its not "deleted" yes. only after the loop is complete and Auras[:i] is called!
				}
			} else { //stays in the list			
				interactable.Auras[i] = aura
				i++				
			}
		}
		interactable.Auras = interactable.Auras[:i]

		if doRecalc {			
			interactable.recalcStats(state.(*MatchState))
		}
	}
	
	//calculate game/npcs/objects
	for _, projectile := range state.(*MatchState).PublicMatchState.Projectile {		
		if projectile == nil || projectile.CreatedAtTick == tick {
			continue
		}
		projectile.Run(state.(*MatchState), projectile, tickrate)	
	}

	//regen
	for _, player := range state.(*MatchState).InternalPlayer {
		fmt.Printf("regen > %v\n", player.Id)
		if(player.LastRegenTick + int64(tickrate) < tick) {
			player.LastRegenTick = tick
			regenPercHP := float64((tick - player.LastHealthDrainTick) / 10) //secs since last dmg
			regenPercHP = math.Max(0, math.Min(1, regenPercHP / 10.0)) //0-100%
			
			regenPercPower := float64((tick - player.LastPowerDrainTick) / 10) //secs since last dmg
			regenPercPower = math.Max(0, math.Min(1, regenPercPower / 10.0)) //0-100%

			player.getPublicPlayer(state.(*MatchState)).Regen(state.(*MatchState), regenPercHP, regenPercPower);
		}
	}

	//send new game state (by creating protobuf message)
	for _, player := range state.(*MatchState).InternalPlayer {		
		if player.Presence == nil {
			continue
		}
		player.MessageCountThisFrame = 0


		out, err := proto.Marshal(&state.(*MatchState).PublicMatchState)
		if err != nil {
				logger.Printf("Failed to encode PublicMatchState:", err)
		}
		//currentPlayerPublic := state.(*MatchState).PublicMatchState.Interactable[player.Id];
		//fmt.Printf("%v @ %v | %v  GCD: %v -- hp: %v/%v mana: %v/%v -- bytes: %v kB/s\n", player.Id, currentPlayerPublic.Position.X, currentPlayerPublic.Position.Y, currentPlayerPublic.GlobalCooldown, currentPlayerPublic.Character.CurrentHealth, state.(*MatchState).GetClassFromDB(currentPlayerPublic.Character).getMaxHp(currentPlayerPublic.Character), currentPlayerPublic.Character.CurrentPower, state.(*MatchState).GetClassFromDB(currentPlayerPublic.Character).getMaxMana(currentPlayerPublic.Character), float64(len(out) * tickrate) / 1000.0)
		dispatcher.BroadcastMessage(1, out, []runtime.Presence { player.Presence }, nil)
	}	
	
	//save for history
	//historyCopy := state.(*MatchState).PublicMatchState
	//state.(*MatchState).OldMatchState[tick] = historyCopy

	//end if no ones sending smth (all dc'ed)
	if true {
		if len(messages) == 0 {
			state.(*MatchState).EmptyCounter = state.(*MatchState).EmptyCounter + 1;
		} else {
			state.(*MatchState).EmptyCounter = 0
		}
		
		if state.(*MatchState).EmptyCounter == 10 {
			return nil
		}
	}

	//calc loop runtime
	if false {
		state.(*MatchState).runtimeSet[state.(*MatchState).runtimeSetIndex] = int64(time.Since(start))
		avg := int64(0)
		for _, time := range state.(*MatchState).runtimeSet {
			avg += time
		}
		avg /= 20
		state.(*MatchState).runtimeSetIndex = (state.(*MatchState).runtimeSetIndex + 1) % 20
		fmt.Printf(" - - duration %v - avg:  %vµs - - \n", time.Since(start), avg/1000.0)
		fmt.Printf(" _ _ _ _ _ new tick %v _ _ _ _ _\n", tick+1)
	}

	return state
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	if state.(*MatchState).Debug {
		logger.Printf("match terminate match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Printf("match terminate match_id %v grace seconds %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), graceSeconds)
	}

	return state
}

