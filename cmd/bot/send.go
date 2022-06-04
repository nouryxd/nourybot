package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// banphraseResponse is the data we receive back from
// the banphrase API
type banphraseResponse struct {
	Banned        bool          `json:"banned"`
	InputMessage  string        `json:"input_message"`
	BanphraseData banphraseData `json:"banphrase_data"`
}

// banphraseData contains details about why a message
// was banphrased.
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

	// {"message": "AHAHAHAHA LUL"}
	reqBody, err := json.Marshal(map[string]string{
		"message": text,
	})
	if err != nil {
		app.logger.Error(err)
	}

	resp, err := http.Post(banPhraseUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		app.logger.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.logger.Error(err)
	}

	var responseObject banphraseResponse
	json.Unmarshal(body, &responseObject)

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

// Send is used to send twitch replies and contains the necessary
// safeguards and logic for that.
func (app *application) Send(target, message string) {
	// Message we are trying to send is empty.
	if len(message) == 0 {
		return
	}

	// Since messages starting with `.` or `/` are used for special actions
	// (ban, whisper, timeout) and so on, we place a emote infront of it so
	// the actions wouldn't execute. `!` and `$` are common bot prefixes so we
	// don't allow them either.
	if message[0] == '.' || message[0] == '/' || message[0] == '!' || message[0] == '$' {
		message = ":tf: " + message
	}

	// check the message for bad words before we say it
	messageBanned, banReason := app.checkMessage(message)
	if messageBanned {
		// Bad message, replace message and log it.
		app.twitchClient.Say(target, "[BANPHRASED] monkaS")
		app.logger.Info("Banned message detected: ", banReason)

		return
	} else {
		// In case the message we are trying to send is longer than the
		// maximum allowed message length on twitch. We split the message
		// in two parts.
		if len(message) > 500 {
			firstMessage := message[0:499]
			secondMessage := message[499:]

			app.twitchClient.Say(target, firstMessage)
			app.twitchClient.Say(target, secondMessage)

			return
		}
		// Message was fine.
		app.twitchClient.Say(target, message)
		return
	}
}
