package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// reply, err := api.Currency(currAmount, currFrom, currTo)

type currencyResponse struct {
	Amount int    `json:"amount"`
	Base   string `json:"base"`
	Rates  rates
}

type rates map[string]interface{}

func Currency(currAmount string, currFrom string, currTo string) (string, error) {
	//
	baseUrl := "https://api.frankfurter.app"
	response, err := http.Get(fmt.Sprintf("%s/latest?amount=%s&from=%s&to=%s", baseUrl, currAmount, currFrom, currTo))
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject currencyResponse
	json.Unmarshal(responseData, &responseObject)

	log.Info(responseObject.Amount)
	log.Info(responseObject.Base)
	log.Info(responseObject.Rates)
	return "", nil
}
