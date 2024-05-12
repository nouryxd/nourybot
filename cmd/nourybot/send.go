package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
)

// banphraseResponse is the data we receive back from the banphrase API
type banphraseResponse struct {
	Banned        bool          `json:"banned"`
	InputMessage  string        `json:"input_message"`
	BanphraseData banphraseData `json:"banphrase_data"`
}

// banphraseData contains details about why a message was banphrased.
type banphraseData struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phrase    string `json:"phrase"`
	Length    int    `json:"length"`
	Permanent bool   `json:"permanent"`
}

var (
	banPhraseUrl = "https://pajlada.pajbot.com/api/v1/banphrases/test"
)

// CheckMessage checks a given message against the banphrase api.
// returns false, "okay" if a message is allowed
// returns true and a string with the reason if it was banned.
// More information:
// https://gist.github.com/pajlada/57464e519ba8d195a97ddcd0755f9715
func (app *application) checkMessage(text string) (bool, string) {
	pattern := `\bi(?:'|\s?a)?m\s?(?:\d|1[0-2])(?:\s?year|$)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		app.Log.Error("Error compiling regular expression:", err)
		return true, "could not check regex"
	}

	if re.MatchString(text) {
		return true, "regex matched"
	}

	// {"message": "AHAHAHAHA LUL"}
	reqBody, err := json.Marshal(map[string]string{
		"message": text,
	})
	if err != nil {
		app.Log.Error(err)
		return true, "could not check banphrase api"
	}

	resp, err := http.Post(banPhraseUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		app.Log.Error(err)
		return true, "could not check banphrase api"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Log.Error(err)
	}

	var responseObject banphraseResponse
	if err := json.Unmarshal(body, &responseObject); err != nil {
		app.Log.Error(err)
		return true, "could not check banphrase api"
	}

	// Bad Message
	//
	// {"phrase": "No gyazo allowed"}
	reason := responseObject.BanphraseData.Name
	if responseObject.Banned {
		return true, fmt.Sprint(reason)
	} else if !responseObject.Banned {
		// Good message
		return false, "okay"
	}

	// Couldn't contact api so assume it was a bad message
	return true, "Banphrase API couldn't be reached monkaS"
}

// SendNoContext is used to send twitch replies without the full twitch.PrivateMessage context to be logged.
func (app *application) SendNoContext(target, message string) {
	pattern := `\bi(?:'|\s?a)?m\s?(?:\d|1[0-2])(?:\s?year|$)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		app.Log.Error("Error compiling regular expression:", err)
		return
	}

	if re.MatchString(message) {
		return
	}

	// Message we are trying to send is empty.
	if len(message) == 0 {
		return
	}

	if target == "forsen" {
		return
	}

	identifier := uuid.NewString()
	go app.Models.SentMessagesLogs.Insert(
		target, message, "unavailable", "unavailable", "unavailable", "unavailable", identifier, "unavailable")

	// Since messages starting with `.` or `/` are used for special actions
	// (ban, whisper, timeout) and so on, we place an emote infront of it so
	// the actions wouldn't execute. `!` and `$` are common bot prefixes so we
	// don't allow them either.
	if message[0] == '.' || message[0] == '/' || message[0] == '!' || message[0] == '$' {
		message = ":tf: " + message
	}

	// check the message for bad words before we say it
	messageBanned, banReason := app.checkMessage(message)
	if !messageBanned {
		// In case the message we are trying to send is longer than the
		// maximum allowed message length on twitch we split the message in two parts.
		// Twitch has a maximum length for messages of 510 characters so to be safe
		// we split and check at 500 characters.
		// https://discuss.dev.twitch.tv/t/missing-client-side-message-length-check/21316
		// TODO: Make it so it splits at a space instead and not in the middle of a word.
		if len(message) > 500 {
			firstMessage := message[0:499]
			secondMessage := message[499:]

			app.TwitchClient.Say(target, firstMessage)
			app.TwitchClient.Say(target, secondMessage)

			return
		} else {
			// Message was fine.
			go app.TwitchClient.Say(target, message)
			return
		}
	} else {
		// Bad message, replace message and log it.
		app.TwitchClient.Say(target, "[BANPHRASED] monkaS")
		app.Log.Infow("banned message detected",
			"target channel", target,
			"message", message,
			"ban reason", banReason,
		)

		return
	}
}

// Send is used to send twitch replies and contains the necessary safeguards and logic for that.
// Send also logs the twitch.PrivateMessage contents into the database.
func (app *application) Send(target, message string, msgContext twitch.PrivateMessage) {
	pattern := `\bi(?:'|\s?a)?m\s?(?:\d|1[0-2])(?:\s?year|$)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		app.Log.Error("Error compiling regular expression:", err)
		return
	}

	if re.MatchString(message) {
		return
	}
	// Message we are trying to send is empty.
	if len(message) == 0 {
		return
	}

	if target == "forsen" {
		return
	}

	commandName := strings.ToLower(strings.SplitN(msgContext.Message, " ", 3)[0][2:])
	identifier := uuid.NewString()
	go app.Models.SentMessagesLogs.Insert(
		target, message, commandName, msgContext.User.Name, msgContext.User.ID, msgContext.Message, identifier, msgContext.Raw)

	// Since messages starting with `.` or `/` are used for special actions
	// (ban, whisper, timeout) and so on, we place an emote infront of it so
	// the actions wouldn't execute. `!` and `$` are common bot prefixes so we
	// don't allow them either.
	if message[0] == '.' || message[0] == '/' || message[0] == '!' || message[0] == '$' {
		message = ":tf: " + message
	}

	// check the message for bad words before we say it
	messageBanned, banReason := app.checkMessage(message)
	if !messageBanned {
		// In case the message we are trying to send is longer than the
		// maximum allowed message length on twitch we split the message in two parts.
		// Twitch has a maximum length for messages of 510 characters so to be safe
		// we split and check at 500 characters.
		// https://discuss.dev.twitch.tv/t/missing-client-side-message-length-check/21316
		// TODO: Make it so it splits at a space instead and not in the middle of a word.
		if len(message) > 500 {
			firstMessage := message[0:499]
			secondMessage := message[499:]

			app.TwitchClient.Say(target, firstMessage)
			app.TwitchClient.Say(target, secondMessage)

			return
		} else {
			// Message was fine.
			go app.TwitchClient.Say(target, message)
			return
		}
	} else {
		// Bad message, replace message and log it.
		app.TwitchClient.Say(target, "[BANPHRASED] monkaS")
		app.Log.Infow("banned message detected",
			"target channel", target,
			"message", message,
			"ban reason", banReason,
		)

		return
	}
}

// SendNoBanphrase does not check the banphrase before sending a twitch mesage.
func (app *application) SendNoBanphrase(target, message string) {
	pattern := `\bi(?:'|\s?a)?m\s?(?:\d|1[0-2])(?:\s?year|$)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		app.Log.Error("Error compiling regular expression:", err)
		return
	}

	if re.MatchString(message) {
		return
	}

	// Message we are trying to send is empty.
	if len(message) == 0 {
		return
	}

	if target == "forsen" {
		return
	}

	identifier := uuid.NewString()
	go app.Models.SentMessagesLogs.Insert(
		target, message, "unavailable", "unavailable", "unavailable", "unavailable", identifier, "unavailable")

	// Since messages starting with `.` or `/` are used for special actions
	// (ban, whisper, timeout) and so on, we place an emote infront of it so
	// the actions wouldn't execute. `!` and `$` are common bot prefixes so we
	// don't allow them either.
	if message[0] == '.' || message[0] == '/' || message[0] == '!' || message[0] == '$' {
		message = ":tf: " + message
	}

	// check the message for bad words before we say it
	// Message was fine.
	go app.TwitchClient.Say(target, message)
}

// SendNoLimit does not check for the maximum message size.
// Used in sending commands from the database since the command has to have
// been gotten in there somehow. So it fits. Still checks for banphrases.
func (app *application) SendNoLimit(target, message string) {
	pattern := `\bi(?:'|\s?a)?m\s?(?:\d|1[0-2])(?:\s?year|$)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		app.Log.Error("Error compiling regular expression:", err)
		return
	}

	if re.MatchString(message) {
		return
	}
	// Message we are trying to send is empty.
	if len(message) == 0 {
		return
	}

	// Since messages starting with `.` or `/` are used for special actions
	// (ban, whisper, timeout) and so on, we place an emote infront of it so
	// the actions wouldn't execute. `!` and `$` are common bot prefixes so we
	// don't allow them either.
	if message[0] == '.' || message[0] == '/' || message[0] == '!' || message[0] == '$' {
		message = ":tf: " + message
	}

	// check the message for bad words before we say it
	messageBanned, banReason := app.checkMessage(message)
	if messageBanned {
		// Bad message, replace message and log it.
		go app.TwitchClient.Say(target, "[BANPHRASED] monkaS")
		app.Log.Infow("banned message detected",
			"target channel", target,
			"message", message,
			"ban reason", banReason,
		)

		return
	} else {
		// In case the message we are trying to send is longer than the
		// maximum allowed message length on twitch we split the message in two parts.
		// Twitch has a maximum length for messages of 510 characters so to be safe
		// we split and check at 500 characters.
		// https://discuss.dev.twitch.tv/t/missing-client-side-message-length-check/21316
		// TODO: Make it so it splits at a space instead and not in the middle of a word.
		// Message was fine.
		identifier := uuid.NewString()
		go app.Models.SentMessagesLogs.Insert(
			target, message, "unavailable", "unavailable", "unavailable", "unavailable", identifier, "unavailable")
		go app.TwitchClient.Say(target, message)
		return
	}
}
