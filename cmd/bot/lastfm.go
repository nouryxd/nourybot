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
	u, err := app.Models.LastFMUsers.Get(twitchLogin)
	if err != nil {
		sugar.Errorw("No LastFM account registered for: ",
			"twitchLogin:", twitchLogin,
		)
		reply := "No lastfm account registered in my database. Use ()register lastfm <username> to register. (Not yet implemented) Otherwise use ()lastfm <username> without registering."
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

	target := message.Channel
	user := u.LastFMUser
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
		"user:", user,
	)

	commands.LastFmUserRecent(target, user, app.TwitchClient)
}
