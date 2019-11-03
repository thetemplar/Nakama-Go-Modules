package main

import (
	"context"
	"time"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/golang/protobuf/proto"
	"fmt"
	"strconv"
	"math"
	"Nakama-Go-Modules/GameDB"
)


type MatchState struct {
	PublicMatchState    PublicMatchState
	EmptyCounter        int
	Debug               bool
	TickRate			int

	Player      		map[string]*InternalInteractable
	PresenceList      	map[string]*runtime.Presence

	ProjectileCounter	int64
	NpcCounter			int64
	
	GameDB				*GameDB.Database
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

		
		Player: make(map[string]*InternalInteractable),
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
		Classname: "Ogre",
		CurrentHealth: 100,
		CurrentPower: 1,
	}
	state.PublicMatchState.Interactable[enemy.Id] = enemy

	enemyInternal := &InternalInteractable{
		PublicMatchState_Interactable: enemy,
		Presence: nil,
		StatModifiers: PlayerStats {},
	}
	enemyInternal.Act = Act_Ogre
	state.Player[enemyInternal.Id] = enemyInternal
	enemyInternal.recalcStats(state)
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
	if state.PublicMatchState.Interactable[userId] != nil || state.Player[userId] != nil || state.GameDB.Classes[classname] == nil {
		return
	}

	state.PublicMatchState.Interactable[userId] = &PublicMatchState_Interactable{
		Id: userId,
		Type: PublicMatchState_Interactable_Player,
		Position: &PublicMatchState_Vector2Df { 
			X: 0.1,
			Y: 0.1,
		},
		Classname: classname,
		CurrentHealth: 100,
		CurrentPower: 100,
	}
	
	state.Player[userId] = &InternalInteractable{
		PublicMatchState_Interactable: state.PublicMatchState.Interactable[userId],
		Presence: *state.PresenceList[userId],	
		TriangleIndex: -1,
		StatModifiers: PlayerStats {},
	}
	state.Player[userId].recalcStats(state)
	fmt.Printf("new character %v spawn @ %v\n", userId, state.Player[userId].Position)
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
		state.(*MatchState).Player[presence.GetUserId()] = nil
		state.(*MatchState).PresenceList[presence.GetUserId()] = nil
		delete(state.(*MatchState).PublicMatchState.Interactable, presence.GetUserId())
		delete(state.(*MatchState).Player, presence.GetUserId())
		delete(state.(*MatchState).PresenceList, presence.GetUserId())

		logger.Printf("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
	}

	return state
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
	for _, player := range state.(*MatchState).Player {	
		
		//finish casts
		if player.CastingTickEnd <= tick && player.CastingSpellId > 0 {
			player.finishCast(state.(*MatchState), state.(*MatchState).GameDB.Spells[player.CastingSpellId], player.CastingTargeted)
			player.stopCastTimer()
		}
		
		//finish swing mainhand 
		if player.Autoattacking && player.AutoattackMainhandTickEnd <= tick && player.AutoattackMainhandTickEnd > 0{
			player.finishAutoattack(state.(*MatchState), GameDB.Item_Slot_Weapon_MainHand, player.AutoattackTargeted)
			
			//queue next swing!
			player.AutoattackMainhandTickEnd = int64(state.(*MatchState).GameDB.Classes[player.Classname].Mainhand.AttackSpeed * player.getMeeleAttackSpeed(state.(*MatchState).GameDB.Classes[player.Classname]) * float32(tickrate)) + tick
		}
		//finish swing offhand
		if player.Autoattacking && player.AutoattackOffhandTickEnd <= tick && player.AutoattackOffhandTickEnd > 0{
			player.finishAutoattack(state.(*MatchState), GameDB.Item_Slot_Weapon_OffHand, player.AutoattackTargeted)
			
			//queue next swing!
			player.AutoattackOffhandTickEnd = int64(state.(*MatchState).GameDB.Classes[player.Classname].Mainhand.AttackSpeed * player.getMeeleAttackSpeed(state.(*MatchState).GameDB.Classes[player.Classname]) * float32(tickrate)) + tick
		}
			
		//substract GCD
		player.GlobalCooldown -= float32(1)/float32(ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int));
	}

	//get new input-counts
	for _, message := range messages { 
		if state.(*MatchState).Player[message.GetUserId()] == nil {
			continue
		}
		if(message.GetOpCode() == 2) {
			state.(*MatchState).Player[message.GetUserId()].MoveMessageCountThisFrame++
		}
	}

	//get new inputs
	for _, message := range messages { 
		if (state.(*MatchState).Player[message.GetUserId()] == nil && message.GetOpCode() != 100) || message.GetOpCode() == 255 {
			continue
		}
		//logger.Printf("message from %v with opcode %v", message.GetUserId(), message.GetOpCode())
		//entry.UserID, entry.SessionId, entry.Username, entry.Node, entry.OpCode, entry.Data, entry.ReceiveTime
		player := state.(*MatchState).Player[message.GetUserId()];

		msg := &Client_Message{}
		if err := proto.Unmarshal(message.GetData(), msg); err != nil {
			logger.Printf("Failed to parse incoming SendPackage Client_Message:", err)
		}

		switch t := msg.Type.(type) {
		case *Client_Message_Character:
			player.LastMessageServerTick = tick
			player.MissingCount = 0
			
			if player.Target != msg.GetCharacter().Target {
				player.stopAutoattackTimer();
			}
			player.Target = msg.GetCharacter().Target;
		case *Client_Message_Cast:
			//is the spell in his spellbook?
			for _, spell := range state.(*MatchState).GameDB.Classes[player.Classname].Spells {
				if (spell.Id == msg.GetCast().SpellId) {
					player.startCast(state.(*MatchState), state.(*MatchState).GameDB.Spells[msg.GetCast().SpellId])
					break
				}
			}
		case *Client_Message_AutoAttack:
			if !player.Autoattacking && player.Target != "" {
				fmt.Printf("startAutoattack %v > %v\n", player.Id, player.Target)
				player.startAutoattack(state.(*MatchState), msg.GetAutoAttack().Attacktype)
			}
		case *Client_Message_CancelAttack:
			player.stopAutoattackTimer();
			player.stopCastTimer();
		case *Client_Message_Move:
			if(msg.GetMove().AbsoluteCoordinates) {
				continue;
			}
			player.LastMovement = msg.GetMove()
			player.stopAutoattackTimer();
			player.stopCastTimer();
			player.performMovement(state.(*MatchState), msg.GetMove().XAxis, msg.GetMove().YAxis, msg.GetMove().Rotation)
		case *Client_Message_SelectChar:
			SpawnPlayer(state.(*MatchState), message.GetUserId(), msg.GetSelectChar().Classname)
			player = state.(*MatchState).Player[message.GetUserId()];
		default:
			fmt.Printf("Unknown Client_Message_Character %v\n", t);
		}  
		player.LastProcessedClientTick = msg.ClientTick
	}			
	
	//did a player not send an package? then re-do his last
	for _, player := range state.(*MatchState).Player {		
		if player.LastMessageServerTick != tick {
			player.MissingCount++
			if player.MissingCount > 1 && player.LastMovement != nil {
				player.MoveMessageCountThisFrame = 1
				logger.Printf("2nd missing Package from player %v in a row, inserting last known package.", player.Id)
		
				player.performMovement(state.(*MatchState), player.LastMovement.XAxis, player.LastMovement.YAxis, player.LastMovement.Rotation)
			}
		}
	}

	//auras
	for _, player := range state.(*MatchState).Player {		
		i := 0
		doRecalc := false
		for _, aura := range player.Auras {
			effect := state.(*MatchState).GameDB.Effects[aura.EffectId]

			switch effect.Type.(type) {
			case *GameDB.Effect_Apply_Aura_Periodic_Damage:
				if int64(float32(aura.AuraTickCount + 1) * effect.Type.(*GameDB.Effect_Apply_Aura_Periodic_Damage).Intervall * float32(tickrate)) + aura.CreatedAtTick < tick {
					aura.AuraTickCount++					
					player.applyAbilityDamage(state.(*MatchState), effect, aura.Creator)
				}
			}			

			//is it depleted?
			if int64(effect.Duration * float32(tickrate)) + aura.CreatedAtTick < tick {
				clEntry := &PublicMatchState_CombatLogEntry {
					Timestamp: tick,
					SourceId: aura.Creator,
					DestinationId: player.Id,
					SourceSpellEffectId: &PublicMatchState_CombatLogEntry_SourceEffectId{aura.EffectId},
					Source: PublicMatchState_CombatLogEntry_Spell,
					Type: &PublicMatchState_CombatLogEntry_Aura{ &PublicMatchState_CombatLogEntry_CombatLogEntry_Aura{
						Event: PublicMatchState_CombatLogEntry_CombatLogEntry_Aura_Depleted,
					}},
				}
				state.(*MatchState).PublicMatchState.Combatlog = append(state.(*MatchState).PublicMatchState.Combatlog, clEntry)

				fmt.Printf("auras run off > %v\n", aura)

				switch effect.Type.(type) {
				case *GameDB.Effect_Apply_Aura_Mod:
					doRecalc = true //cant do "recalcStats" here, since its not "deleted" yes. only after the loop is complete and Auras[:i] is called!
				}
			} else { //stays in the list			
				player.Auras[i] = aura
				i++				
			}
		}
		player.Auras = player.Auras[:i]

		if doRecalc {			
			player.recalcStats(state.(*MatchState))
		}
	}
	
	//calculate game/npcs/objects
	for _, projectile := range state.(*MatchState).PublicMatchState.Projectile {		
		if projectile == nil || projectile.CreatedAtTick == tick {
			continue
		}
		projectile.Run(state.(*MatchState), projectile, tickrate)	
	}
	for _, player := range state.(*MatchState).Player {
		if player.Presence != nil || !player.IsEngaged || player.CurrentHealth <= 0 {
			continue
		}
		player.Act(state.(*MatchState), player)
	}

	//regen
	for _, player := range state.(*MatchState).Player {
		if(player.LastRegenTick + int64(tickrate) < tick) {
			player.LastRegenTick = tick
			regenPercHP := float64((tick - player.LastHealthDrainTick) / 10) //secs since last dmg
			regenPercHP = math.Max(0, math.Min(1, regenPercHP / 10.0)) //0-100%
			
			regenPercPower := float64((tick - player.LastPowerDrainTick) / 10) //secs since last dmg
			regenPercPower = math.Max(0, math.Min(1, regenPercPower / 10.0)) //0-100%

			player.regen(state.(*MatchState), regenPercHP, regenPercPower);
		}
	}

	//send new game state (by creating protobuf message)
	for _, player := range state.(*MatchState).Player {		
		if player.Presence == nil {
			continue
		}
		player.MoveMessageCountThisFrame = 0


		out, err := proto.Marshal(&state.(*MatchState).PublicMatchState)
		if err != nil {
				logger.Printf("Failed to encode PublicMatchState:", err)
		}
		dispatcher.BroadcastMessage(1, out, []runtime.Presence { player.Presence }, nil, true)
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

