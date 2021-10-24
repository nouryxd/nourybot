package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type currencyResponse struct {
	Amount int                `json:"amount"`
	Base   string             `json:"base"`
	Rates  map[string]float32 `json:"rates"`
}

// Currency returns the exchange rate for a given given amount to another.
func Currency(currAmount, currFrom, currTo string) (string, error) {
	baseUrl := "https://api.frankfurter.app"
	currFromUpper := strings.ToUpper(currFrom)
	currToUpper := strings.ToUpper(currTo)

	// https://api.frankfurter.app/latest?amount=10&from=gbp&to=usd
	response, err := http.Get(fmt.Sprintf("%s/latest?amount=%s&from=%s&to=%s", baseUrl, currAmount, currFromUpper, currToUpper))
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject currencyResponse
	json.Unmarshal(responseData, &responseObject)

	// log.Info(responseObject.Amount)
	// log.Info(responseObject.Base)
	// log.Info(responseObject.Rates[strings.ToUpper(currTo)])

	if responseObject.Rates[currToUpper] != 0 {
		return fmt.Sprintf("%s %s = %.2f %s", currAmount, currFromUpper, responseObject.Rates[currToUpper], currToUpper), nil
	}

	return "Something went wrong FeelsBadMan", nil
}
