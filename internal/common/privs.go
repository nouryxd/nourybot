package common

import "github.com/gempir/go-twitch-irc/v3"

// ElevatedPrivsMessage is checking a given message twitch.PrivateMessage
// if it came from a moderator/vip/or broadcaster and returns a bool
func ElevatedPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["vip"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}

// ModPrivsMessage is checking a given message twitch.PrivateMessage
// if it came from a moderator or broadcaster and returns a bool
func ModPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}
