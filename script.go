
package main

import (
	"fmt"
	"context"
    "encoding/json"
	"database/sql"
    "github.com/heroiclabs/nakama/api"
	"github.com/heroiclabs/nakama/runtime"
	"github.com/gofrs/uuid"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Print("RUNNING IN GO: Script.go")
	if err := initializer.RegisterMatchmakerMatched(CreateMatchmakerMatch); err != nil {
		return err
	}
	if err := initializer.RegisterMatch("match", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
		return &Match{}, nil
	}); err != nil {
		return err
	}
	if err := initializer.RegisterRpc("createMatch", CreateManualMatch); err != nil {
		return err
	}	
	
	if err := initializer.RegisterRpc("getPlayers", GetPlayers); err != nil {
		return err
	}	


	//send note to blocked user & send refresh friendlist notification
	if err := initializer.RegisterAfterBlockFriends(NotificationBlockFriends); err != nil {
		return err
	}
	//send note to deleted user & send refresh friendlist notification
	if err := initializer.RegisterAfterDeleteFriends(NotificationDeleteFriends); err != nil {
		return err
	}
	//only send refresh friendlist notification
	if err := initializer.RegisterAfterAddFriends(NotificationAddFriends); err != nil {
		return err
	}
	
	//send note to refresh grouplist
	if err := initializer.RegisterAfterJoinGroup(NotificationJoinGroup); err != nil {
		return err
	}	
	//send note to refresh grouplist
	if err := initializer.RegisterAfterPromoteGroupUsers(NotificationPromoteGroupUsers); err != nil {
		return err
	}
	//send note to refresh grouplist
	if err := initializer.RegisterAfterKickGroupUsers(NotificationKickGroupUsers); err != nil {
		return err
	}
	//send note to refresh grouplist
	if err := initializer.RegisterAfterAddGroupUsers(NotificationAddGroupUsers); err != nil {
		return err
	}
	//send note to refresh grouplist
	if err := initializer.RegisterAfterLeaveGroup(NotificationLeaveGroup); err != nil {
		return err
	}
	//send note to refresh grouplist
	if err := initializer.RegisterAfterDeleteGroup(NotificationDeleteGroup); err != nil {
		return err
	}
	//send note to refresh grouplist
	if err := initializer.RegisterAfterUpdateGroup(RegisterAfterUpdateGroup); err != nil {
		return err
	}

	return nil
}

func CreateMatchmakerMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {	
	logger.Print(" >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> createMatchmakerMatch <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	for _, entry := range entries { 
		logger.Printf("%+v\n", entry)
	}
	
    params := map[string]interface{}{ "debug": "true" }
    matchID, err := nk.MatchCreate(ctx, "match", params)
    if err != nil {
        // Handle errors as you want.
        return "", err
    }
    return matchID, nil
}
	
func CreateManualMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Print(" >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> createManualMatch <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	
    params := map[string]interface{}{ "debug": "true" }
    matchID, err := nk.MatchCreate(ctx, "match", params)
	if err != nil {
        // Handle errors as you want.
        return "", err
    }
	return string(matchID), nil
}


func GetPlayers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	query := `
		SELECT id, username, display_name, avatar_url,
		lang_tag, location, timezone, metadata
		FROM users`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		logger.Printf("Error retrieving players.", (err))
		return "", err
	}
	defer rows.Close()

	players := make([]*api.User, 0)

	for rows.Next() {
		var id string
		var username sql.NullString
		var displayName sql.NullString
		var avatarURL sql.NullString
		var lang sql.NullString
		var location sql.NullString
		var timezone sql.NullString
		var metadata []byte

		if err = rows.Scan(&id, &username, &displayName, &avatarURL, &lang, &location, &timezone, &metadata); err != nil {
			logger.Printf("Error retrieving players.", (err))
			return "", err
		}

		playerID := uuid.FromStringOrNil(id)
/*
		online := false
		if tracker != nil {
			online = tracker.StreamExists(PresenceStream{Mode: StreamModeNotifications, Subject: playerID})
		}
*/
		user := &api.User{
			Id:          playerID.String(),
			Username:    username.String,
			DisplayName: displayName.String,
			AvatarUrl:   avatarURL.String,
			LangTag:     lang.String,
			Location:    location.String,
			Timezone:    timezone.String,
			Metadata:    string(metadata),
			//Online:      online,
		}

		players = append(players, user)
	}
	if err = rows.Err(); err != nil {
		logger.Printf("Error retrieving players.", (err))
		return "", err
	}
	
	b, err := json.Marshal(players)
	if err != nil {
		logger.Printf("Error converting players to json.", (err))
    }
	return string(b), nil
}


//friends
func NotificationBlockFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) error {	
	userID := ctx.Value("user_id").(string)
	//to the now-blocked user
	subject := fmt.Sprintf("%v blocked you", userID)	
    content := map[string]interface{}{ "username": userID }
	nk.NotificationSend(ctx, in.Ids[0], subject, content, 1, "" , true)
	
	//confirm to the sender
	subject = fmt.Sprintf("you blocked %v", userID)
    content = map[string]interface{}{ "username": in.Ids[0] }
	nk.NotificationSend(ctx, userID, subject, content, 2, "" , true)
	nk.NotificationSend(ctx, userID, subject, content, 1, "" , true)
	
	logger.Printf("Intercepted NotificationBlockFriends: %s blocked %s", userID, in.Ids[0])
	return nil
}

func NotificationDeleteFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) error {
	userID := ctx.Value("user_id").(string)
	
	//to the now-deleted user
	subject := fmt.Sprintf("%v deleted you from his friendlist", userID)
    content := map[string]interface{}{ "username": userID }
	nk.NotificationSend(ctx, in.Ids[0], subject, content, 3, "" , true)
	nk.NotificationSend(ctx, in.Ids[0], subject, content, 1, "" , true)
	
	//confirm to the sender
	subject = fmt.Sprintf("you deleted %v from your friendlist", userID)
    content = map[string]interface{}{ "username": in.Ids[0] }
	nk.NotificationSend(ctx, userID, subject, content, 1, "" , true)
	
	logger.Printf("Intercepted NotificationDeleteFriends: %s deleted %s from his list", userID, in.Ids[0])	
	return nil
}

func NotificationAddFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) error {
	userID := ctx.Value("user_id").(string)
	
	//to the now-deleted user
	subject := fmt.Sprintf("%v added you from to his friendlist", userID)
    content := map[string]interface{}{ "username": userID }
	nk.NotificationSend(ctx, in.Ids[0], subject, content, 1, "" , true)
	
	//confirm to the sender
	subject = fmt.Sprintf("you added %v to your friendlist", userID)
    content = map[string]interface{}{ "username": in.Ids[0] }
	nk.NotificationSend(ctx, userID, subject, content, 1, "" , true)
	
	logger.Printf("Intercepted NotificationAddFriends: %s added %s to his list", userID, in.Ids[0])	
	return nil
}

//groups
func NotificationJoinGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationJoinGroup")
	return nil
}

func NotificationPromoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationPromoteGroupUsers")
	return nil
}

func NotificationKickGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationKickGroupUsers")
	return nil
}

func NotificationAddGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) error {	
	logger.Printf("Intercepted NotificationAddGroupUsers")
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationAddGroupUsers")
	return nil
}

func NotificationLeaveGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationLeaveGroup")
	return nil
}

func NotificationDeleteGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted NotificationDeleteGroup")
	return nil
}

func RegisterAfterUpdateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) error {	
	query := "SELECT source_id, destination_id FROM group_edge WHERE source_id = $1::UUID"

	rows, err := db.QueryContext(ctx, query, in.GroupId)
	if err != nil {
		logger.Printf("Error retrieving group_edge. %s ", (err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var source_id string
		var destination_id string
		if err = rows.Scan(&source_id, &destination_id); err != nil {
			logger.Printf("Error retrieving group_edge. %s ", (err))
			return err
		}
		
		subject := fmt.Sprintf("NotificationJoinGroup %v", source_id)
		content := map[string]interface{}{ "group": source_id }
		nk.NotificationSend(ctx, destination_id, subject, content, 5, "" , true)
	}	
	
	logger.Printf("Intercepted RegisterAfterUpdateGroup")
	return nil
}



