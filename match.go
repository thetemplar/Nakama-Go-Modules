package main

import (
	"context"
	"math"
	"time"
	"database/sql"
	"github.com/heroiclabs/nakama/runtime"
	"github.com/golang/protobuf/proto"
	"fmt"
	"strconv"
)

type MatchState struct {
	PublicMatchState    PublicMatchState
	EmptyCounter        int
	Debug               bool

	InternalPlayer      map[string]*InternalPlayer

	ProjectileCounter	int64
	NpcCounter			int64
	
	GameDB				*GameDB
	Map					*Map
	
	runtimeSet			[]int64
	runtimeSetIndex		int
}  

type Match struct{
}


func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	logger.Print(" >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> MatchInit <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	for _, entry := range params { 
		logger.Printf("%+v\n", entry)
	}

	state := &MatchState{
		Debug: false,
		EmptyCounter : 0,
		PublicMatchState : PublicMatchState{
			Interactable: make(map[string]*PublicMatchState_Interactable),
			Projectile: make(map[string]*PublicMatchState_Projectile),
			Combatlog: make([]*PublicMatchState_CombatLogEntry, 0),
		},
		InternalPlayer: make(map[string]*InternalPlayer),
		runtimeSet: make([]int64, 20),
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
		CurrentHealth: 100,
		CurrentPower: 0,
		MaxHealth: 100,
		MaxPower: 0,
		Auras: make([]*PublicMatchState_Aura, 0),
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

	tickRate := 10
	label := ""

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

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	for _, presence := range presences {
		state.(*MatchState).PublicMatchState.Interactable[presence.GetUserId()] = &PublicMatchState_Interactable{
			Id: presence.GetUserId(),
			Type: PublicMatchState_Interactable_Player,
			Position: &PublicMatchState_Vector2Df {
				X: 22,
				Y: 7.555557,
			},
			CurrentHealth: 100,
			CurrentPower: 100,
			MaxHealth: 100,
			MaxPower: 100,
		}
		
		state.(*MatchState).InternalPlayer[presence.GetUserId()] = &InternalPlayer{
			Id: presence.GetUserId(),
			Presence: presence,
			BasePlayerStats: PlayerStats {
				MovementSpeed: 20.0,
			},			
			TriangleIndex: -1,
			StatModifiers: PlayerStats {},
		}
		
		logger.Printf("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
	}

	return state
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	for _, presence := range presences {		
		state.(*MatchState).PublicMatchState.Interactable[presence.GetUserId()] = nil
		state.(*MatchState).InternalPlayer[presence.GetUserId()] = nil

		logger.Printf("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
	}

	return state
}

func PublicMatchState_Vector2Df_Rotate(v PublicMatchState_Vector2Df, degrees float32) PublicMatchState_Vector2Df {
	ca := float32(math.Cos(float64(360 - degrees) * 0.01745329251)); //0.01745329251
	sa := float32(math.Sin(float64(360 - degrees) * 0.01745329251));

	vec := PublicMatchState_Vector2Df {
		X: ca * v.X - sa * v.Y,
		Y: sa * v.X + ca * v.Y,
	}

	return vec
}


func PerformInputs(logger runtime.Logger, state interface{}, message runtime.MatchData, tickrate int) {
	if state.(*MatchState).InternalPlayer[message.GetUserId()] == nil || state.(*MatchState).PublicMatchState.Interactable[message.GetUserId()] == nil {
		return
	}
	currentPlayerInternal := state.(*MatchState).InternalPlayer[message.GetUserId()];
	currentPlayerPublic   := state.(*MatchState).PublicMatchState.Interactable[message.GetUserId()];
	
	msg := &Client_Character{}
	if err := proto.Unmarshal(message.GetData(), msg); err != nil {
		logger.Printf("Failed to parse incoming SendPackage Client_Character:", err)
	}

	//ClientState := state.(*MatchState).OldMatchState[msg.ServerTickPerformingOn]
	add := PublicMatchState_Vector2Df {
		X: msg.XAxis / float32(currentPlayerInternal.MessageCountThisFrame) * ((currentPlayerInternal.BasePlayerStats.MovementSpeed - currentPlayerInternal.StatModifiers.MovementSpeed) / float32(tickrate)),
		Y: msg.YAxis / float32(currentPlayerInternal.MessageCountThisFrame) * ((currentPlayerInternal.BasePlayerStats.MovementSpeed - currentPlayerInternal.StatModifiers.MovementSpeed) / float32(tickrate)),
	}

	if currentPlayerInternal.CastingSpellId > 0 && state.(*MatchState).GameDB.Spells[currentPlayerInternal.CastingSpellId].InterruptedBy != GameDB_Interrupt_None && (msg.YAxis != 0 || msg.XAxis != 0) {
		fmt.Printf("cancelCast %v\n", currentPlayerInternal.CastingSpellId)
		currentPlayerPublic.cancelCast(state.(*MatchState))
	}

	rotatedAdd := add.rotate(msg.Rotation)
	currentPlayerPublic.Position.X += rotatedAdd.X;
	currentPlayerPublic.Position.Y += rotatedAdd.Y;
	currentPlayerPublic.Rotation = msg.Rotation;
	
	currentPlayerPublic.Target = msg.Target;

	//am i still in my triangle?	
	if currentPlayerInternal.TriangleIndex >= 0 {
		isItIn, w1, w2, w3 := state.(*MatchState).Map.Triangles[currentPlayerInternal.TriangleIndex].isInTriangle(currentPlayerPublic.Position)
		fmt.Printf("am i still in my triangle?: %v %v (w1:%v  w2:%v  w3:%v)\n", currentPlayerInternal.TriangleIndex, isItIn, w1, w2, w3)
		if !isItIn {
			currentPlayerInternal.TriangleIndex = -1
		}
	}

	//no current triangle_index?
	if currentPlayerInternal.TriangleIndex < 0 {
		//find triangle I am in
		found := false
		for i, triangle := range state.(*MatchState).Map.Triangles {
			isItIn, w1, w2, w3 := triangle.isInTriangle(currentPlayerPublic.Position)
			if isItIn {
				currentPlayerInternal.TriangleIndex = int64(i)
				found = true
				fmt.Printf("triangle.isInTriangle: %v  (w1:%v  w2:%v  w3:%v)\n", i, w1, w2, w3)
				break;
			}
		}
		if !found {
			fmt.Printf("ERROR @ triangle.isInTriangle %v|%v\n", currentPlayerPublic.Position.X, currentPlayerPublic.Position.Y )
			currentPlayerPublic.Position.X -= rotatedAdd.X;
			currentPlayerPublic.Position.Y -= rotatedAdd.Y;
		} 
	}

	currentPlayerPublic.LastProcessedClientTick = msg.ClientTick
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	if state.(*MatchState).Debug {
		logger.Printf("match loop match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Printf("match loop match_id %v message count %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), len(messages))
	}
    start := time.Now()
	fmt.Printf(" _ _ _ _ _ new tick %v _ _ _ _ _\n", tick)
	state.(*MatchState).PublicMatchState.Tick = tick	
	state.(*MatchState).PublicMatchState.Combatlog = make([]*PublicMatchState_CombatLogEntry, 0)
	tickrate := ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int);

	
	for _, player := range state.(*MatchState).InternalPlayer {		
		if player == nil || player.CastingSpellId <= 0 {
			continue
		}
		
		//finish casts
		if int64(state.(*MatchState).GameDB.Spells[player.CastingSpellId].CastTime * float32(tickrate)) + player.CastingTickStarted < tick {
			state.(*MatchState).PublicMatchState.Interactable[player.Id].finishCast(state.(*MatchState), player.CastingSpellId, player.CastingTargeted)
			player.CastingSpellId = -1
			player.CastingTickStarted = 0
			player.CastingTargeted = ""
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
		//logger.Printf("message from %v with opcode %v", message.GetUserId(), message.GetOpCode())
		//entry.UserID, entry.SessionId, entry.Username, entry.Node, entry.OpCode, entry.Data, entry.ReceiveTime
		if state.(*MatchState).InternalPlayer[message.GetUserId()] == nil {
			continue
		}
		currentPlayerInternal := state.(*MatchState).InternalPlayer[message.GetUserId()];
		currentPlayerPublic   := state.(*MatchState).PublicMatchState.Interactable[message.GetUserId()];

		if message.GetOpCode() == 0 {
			currentPlayerInternal.LastMessage = message
			currentPlayerInternal.LastMessageServerTick = tick
			currentPlayerInternal.MissingCount = 0
			
			PerformInputs(logger, state, currentPlayerInternal.LastMessage, tickrate)
		} else if message.GetOpCode() == 1 {
			msg := &Client_Cast{}
			if err := proto.Unmarshal(message.GetData(), msg); err != nil {
				logger.Printf("Failed to parse incoming SendPackage Client_Cast:", err)
			}
			currentPlayerPublic.startCast(state.(*MatchState), msg.SpellId)
		}
	}
	
	//did a player not send an package? then re-do his last
	for _, player := range state.(*MatchState).InternalPlayer {		
		if player == nil {
			continue
		}
		if player.LastMessageServerTick != tick {
			player.MissingCount++
			if player.MissingCount > 1 && player.LastMessage != nil {
				player.MessageCountThisFrame = 1
				logger.Printf("2nd missing Package from player %v in a row, inserting last known package.", player.Id)
				PerformInputs(logger, state, player.LastMessage, tickrate)
			}
		}
	}

	//auras
	for _, interactable := range state.(*MatchState).PublicMatchState.Interactable {		
		if interactable == nil{
			continue
		}
		i := 0
		doRecalc := false
		for _, aura := range interactable.Auras {
			effect := state.(*MatchState).GameDB.Effects[aura.EffectId]

			switch effect.Type.(type) {
			case *GameDB_Effect_Apply_Aura_Periodic_Damage:
				if int64(float32(aura.AuraTickCount + 1) * effect.Type.(*GameDB_Effect_Apply_Aura_Periodic_Damage).Intervall * float32(tickrate)) + aura.CreatedAtTick < tick {
					aura.AuraTickCount++
					dmg := randomInt(effect.Type.(*GameDB_Effect_Apply_Aura_Periodic_Damage).ValueMin, effect.Type.(*GameDB_Effect_Apply_Aura_Periodic_Damage).ValueMax);
					interactable.CurrentHealth -= dmg;
				
					clEntry := &PublicMatchState_CombatLogEntry {
						Timestamp: state.(*MatchState).PublicMatchState.Tick,
						SourceId: aura.Creator,
						DestinationId: interactable.Id,
						SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{effect.Id},
						Source: PublicMatchState_CombatLogEntry_Spell,
						Type: &PublicMatchState_CombatLogEntry_Damage{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Damage{
							Amount: dmg,
						}},
					}
					state.(*MatchState).PublicMatchState.Combatlog = append(state.(*MatchState).PublicMatchState.Combatlog, clEntry)
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
		fmt.Printf("calc proj %v\n", projectile)
		projectile.Run(state.(*MatchState), projectile, tickrate)		
	}

	for _, npc := range state.(*MatchState).PublicMatchState.Interactable {		
		if npc == nil || npc.Type == PublicMatchState_Interactable_Player {
			continue
		}

	}

	//send new game state (by creating protobuf message)
	for _, player := range state.(*MatchState).InternalPlayer {		
		if player == nil || player.Presence == nil {
			continue
		}
		player.MessageCountThisFrame = 0

		currentPlayerPublic := state.(*MatchState).PublicMatchState.Interactable[player.Id];

		out, err := proto.Marshal(&state.(*MatchState).PublicMatchState)
		if err != nil {
				logger.Printf("Failed to encode PublicMatchState:", err)
		}
		fmt.Printf("%v @ %v | %v  GCD: %v | bytes: %v kB/s\n", player.Id, currentPlayerPublic.Position.X, currentPlayerPublic.Position.Y, currentPlayerPublic.GlobalCooldown, float64(len(out) * tickrate) / 1000.0)
		dispatcher.BroadcastMessage(1, out, []runtime.Presence { player.Presence }, nil)
	}	
	
	//save for history
	//historyCopy := state.(*MatchState).PublicMatchState
	//state.(*MatchState).OldMatchState[tick] = historyCopy



	//end if no ones sending smth (all dc'ed)
	if len(messages) == 0 {
		state.(*MatchState).EmptyCounter = state.(*MatchState).EmptyCounter + 1;
	} else {
		state.(*MatchState).EmptyCounter = 0
	}
	
	if state.(*MatchState).EmptyCounter == 10 {
		return nil
	}
	state.(*MatchState).runtimeSet[state.(*MatchState).runtimeSetIndex] = int64(time.Since(start))
	avg := int64(0)
	for _, time := range state.(*MatchState).runtimeSet {
		avg += time
	}
	avg /= 20
	state.(*MatchState).runtimeSetIndex = (state.(*MatchState).runtimeSetIndex + 1) % 20
	fmt.Printf(" - - duration %v - avg:  %vµs - - \n", time.Since(start), avg/1000.0)
	
	return state
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	if state.(*MatchState).Debug {
		logger.Printf("match terminate match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Printf("match terminate match_id %v grace seconds %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), graceSeconds)
	}

	return state
}

