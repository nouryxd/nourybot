package main

import (
	"fmt"

	"github.com/lyx0/nourybot/pkg/ivr"
	"github.com/nicklaw5/helix/v2"
)

const (
	eventSubResponseOK           = 200
	eventSubResponseAccepted     = 202
	eventSubResponseNoContent    = 204
	eventSubResponseBadRequest   = 400
	eventSubResponseUnauthorized = 401
	eventSubResponseForbidden    = 403
)

func (app *application) createLiveSubscription(target, channel string) string {
	// client, err := helix.NewClient(&helix.Options{
	// 	ClientID:       app.Config.twitchClientId,
	// 	AppAccessToken: app.HelixClient.GetAppAccessToken(),
	// })
	// if err != nil {
	// 	app.Log.Errorw("Error creating new helix client",
	// 		"err", err,
	// 	)
	// }

	uid := ivr.IDByUsername(channel)
	if uid == "xd" {
		return fmt.Sprintf("Could not find user with name %v", channel)
	}

	resp, err := app.HelixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
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
		app.Log.Errorw("Error creating EventSub subscription",
			"resp", resp,
			"err", err,
		)
	}

	app.Log.Infof("%+v\n", resp)
	return fmt.Sprintf("Created subscription for channel %v; uid=%v", channel, uid)
}

func (app *application) deleteLiveSubscription(target, channel string) string {
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
	if uid == "xd" {
		return fmt.Sprintf("Could not find user with name %v", channel)
	}

	resp, err := client.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{
		UserID: uid,
	})
	if err != nil {
		app.Log.Errorw("Error getting EventSub subscriptions",
			"resp", resp,
			"err", err,
		)
		return "Something went wrong FeelsBadMan"
	}

	app.Log.Infof("%+v\n", resp)

	if resp.StatusCode != eventSubResponseOK {
		return "Subscription was not found FeelsBadMan"
	}
	if len(resp.Data.EventSubSubscriptions) == 0 {
		return "No such subscription found"
	}
	eventSubID := resp.Data.EventSubSubscriptions[0].ID

	resp2, err := client.RemoveEventSubSubscription(eventSubID)
	if err != nil {
		app.Log.Errorw("Error getting EventSub subscriptions",
			"resp", resp,
			"err", err,
		)
		return "Something went wrong FeelsBadMan"
	}

	if resp2.StatusCode != eventSubResponseNoContent {
		return "Subscription was not found FeelsBadMan"
	}

	app.Log.Infof("%+v\n", resp2)
	return fmt.Sprintf("Successfully deleted live notification for channel %s; id=%v", channel, uid)
}
