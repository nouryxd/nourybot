package main

import (
	"fmt"

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

// createLiveSubscription creates a stream.online twitch eventsub subscription for the specified channel
func (app *application) createLiveSubscription(target, channel string) string {
	uid, err := app.getUserID(channel)
	if err != nil {
		app.Log.Error(err)
		return fmt.Sprint(ErrGenericErrorMessage)
	}

	resp, err := app.HelixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeStreamOnline,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: uid,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: fmt.Sprintf("https://bot.noury.li/eventsub/%v", target),
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

// deleteLiveSubscription deletes a stream.live twitch eventsub subscription for the specified channel
func (app *application) deleteLiveSubscription(target, channel string) string {
	uid, err := app.getUserID(channel)
	if err != nil {
		app.Log.Error(err)
		return fmt.Sprint(ErrGenericErrorMessage)
	}

	resp, err := app.HelixClient.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{
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

	// Since we can't pass more than one variable to the api call we loop
	// over all the resp.Data.EventSubSubscriptions we get back and check
	// which one has the type we want for this event. Then use this ones ID.
	var eventSubID string
	for i := 0; i < len(resp.Data.EventSubSubscriptions); i++ {
		if resp.Data.EventSubSubscriptions[i].Type == "stream.online" &&
			resp.Data.EventSubSubscriptions[i].Transport.Callback == fmt.Sprintf("https://bot.noury.li/eventsub/%s", target) {
			eventSubID = resp.Data.EventSubSubscriptions[i].ID
		}
	}
	//eventSubID := resp.Data.EventSubSubscriptions[0].ID
	//app.SendNoContext("nouryxd", eventSubID)

	resp2, err := app.HelixClient.RemoveEventSubSubscription(eventSubID)
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

// createOfflineSubscription creates a stream.offline twitch eventsub
// subscription for the specified channel.
func (app *application) createOfflineSubscription(target, channel string) string {
	uid, err := app.getUserID(channel)
	if err != nil {
		app.Log.Error(err)
		return fmt.Sprint(ErrGenericErrorMessage)
	}

	resp, err := app.HelixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeStreamOffline,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: uid,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: fmt.Sprintf("https://bot.noury.li/eventsub/%v", target),
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
	return fmt.Sprintf("Created offline subscription for channel %v; uid=%v", channel, uid)
}

// deleteOfflineSubscription deletes a stream.offline twitch eventsub
// subscription for the specified channel
func (app *application) deleteOfflineSubscription(target, channel string) string {
	uid, err := app.getUserID(channel)
	if err != nil {
		app.Log.Error(err)
		return fmt.Sprint(ErrGenericErrorMessage)
	}

	resp, err := app.HelixClient.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{
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

	// Since we can't pass more than one variable to the api call we loop
	// over all the resp.Data.EventSubSubscriptions we get back and check
	// which one has the type we want for this event. Then use this ones ID.
	var eventSubID string
	for i := 0; i < len(resp.Data.EventSubSubscriptions); i++ {
		if resp.Data.EventSubSubscriptions[i].Type == "stream.offline" &&
			resp.Data.EventSubSubscriptions[i].Transport.Callback == fmt.Sprintf("https://bot.noury.li/eventsub/%s", target) {
			eventSubID = resp.Data.EventSubSubscriptions[i].ID
		}
	}
	//eventSubID = resp.Data.EventSubSubscriptions[0].ID
	//app.SendNoContext("nouryxd", eventSubID)

	resp2, err := app.HelixClient.RemoveEventSubSubscription(eventSubID)
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
