package main

import (
	"fmt"

	"github.com/lyx0/nourybot/pkg/ivr"
	"github.com/nicklaw5/helix/v2"
)

func (app *application) createLiveSubscription(target, channel string) string {
	client, err := helix.NewClient(&helix.Options{
		ClientID:       app.Config.twitchClientId,
		AppAccessToken: app.HelixClient.GetAppAccessToken(),
	})
	if err != nil {
		app.Log.Errorw("Error creating new helix client",
			"err", err,
		)
	}

	uid := ivr.IDByUsername(channel)

	resp, err := client.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeStreamOnline,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: uid,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: "https://bot.noury.li/eventsub",
			Secret:   app.Config.eventSubSecret,
		},
	})
	if err != nil {
		app.Log.Errorw("Error creating EvenSub subscription",
			"resp", resp,
			"err", err,
		)
	}

	app.Log.Infof("%+v\n", resp)
	return fmt.Sprintf("Created subscription for channel %v; uid=%v", channel, uid)
}
