package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/commands"
	"github.com/lyx0/nourybot/internal/common"
	"go.uber.org/zap"
)

func (app *Application) CheckLastFM(message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	twitchLogin := message.User.Name
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
	)
	lastfmUser, err := app.Models.Users.GetLastFM(twitchLogin)
	if err != nil {
		sugar.Errorw("No LastFM account registered for: ",
			"twitchLogin:", twitchLogin,
		)
		reply := "No lastfm account registered in my database. Use ()register lastfm <username> to register. (Not yet implemented) Otherwise use ()lastfm <username> without registering."
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

	target := message.Channel
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
		"user:", lastfmUser,
	)

	commands.LastFmUserRecent(target, lastfmUser, app.TwitchClient)
}

//func (app *Application) SetLastFMUser(lastfmUser string, message twitch.PrivateMessage) {
//	sugar := zap.NewExample().Sugar()
//	defer sugar.Sync()
//
//	user := &data.LastFMUser{
//		TwitchLogin: message.User.Name,
//		TwitchID:    message.User.ID,
//		LastFMUser:  lastfmUser,
//	}
//	sugar.Infow("User:: ",
//		"user:", user,
//	)
//
//	err := app.Models.LastFMUsers.Insert(user)
//	if err != nil {
//		if err != nil {
//			reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
//			common.Send(message.Channel, reply, app.TwitchClient)
//			return
//		}
//	} else {
//		reply := fmt.Sprintf("Successfully set your lastfm account as %v", lastfmUser)
//		common.Send(message.Channel, reply, app.TwitchClient)
//		return
//	}
//}
//
