package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
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

// CheckMessage checks if a message contains
// banphrased content.
// If a message is allowed it returns
// false, "okay"
// If a message is not allowed it returns:
// true, "[banphrased] monkaS"
// More information:
// https://gist.github.com/pajlada/57464e519ba8d195a97ddcd0755f9715
func CheckMessage(text string) (bool, string) {
	log.Info("fn CheckMessage")

	// {"message": "AHAHAHAHA LUL"}
	reqBody, err := json.Marshal(map[string]string{
		"message": text,
	})
	if err != nil {
		log.Error(err)
	}

	resp, err := http.Post("https://pajlada.pajbot.com/api/v1/banphrases/test", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject banphraseResponse
	json.Unmarshal(body, &responseObject)

	// {"phrase": "No gyazo allowed"}
	reason := responseObject.BanphraseData.Name

	// log.Info("Bancheck: ", responseObject.Banned)
	// log.Info("Reason: ", reason)
	// log.Info("Bancheck: ", responseObject.Banned)

	// Bad message
	if responseObject.Banned {
		return true, fmt.Sprintf("Banphrased, reason: %s", reason)
	} else {
		// Good message
		return false, "okay"
	}

	// Couldn't contact api so assume it was a bad message
	return true, "Banphrase API couldn't be reached monkaS"
}
