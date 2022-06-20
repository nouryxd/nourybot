package bot

import "github.com/gempir/go-twitch-irc/v3"

func (bot *Bot) handleWhisperMessage(message twitch.WhisperMessage) {
	// Print the whisper message for now.
	// TODO: Implement a basic whisper handler.
	bot.logger.Infof("[#whisper]:%s: %s", message.User.DisplayName, message.Message)
}
