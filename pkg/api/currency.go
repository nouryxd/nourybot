package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// reply, err := api.Currency(currAmount, currFrom, currTo)

type currencyResponse struct {
	Amount int                `json:"amount"`
	Base   string             `json:"base"`
	Rates  map[string]float32 `json:"rates"`
}

// type rates map[string]interface{}
// type rates struct {
// 	AUD float32 `json:"AUD"`
// 	BGN float32 `json:"BGN"`
// 	BRL float32 `json:"BRL"`
// 	CAD float32 `json:"CAD"`
// 	CHF float32 `json:"CHF"`
// 	CNY float32 `json:"CNY"`
// 	CZK float32 `json:"CZK"`
// 	DKK float32 `json:"DKK"`
// 	GBP float32 `json:"GBP"`
// 	HKD float32 `json:"HKD"`
// 	HRK float32 `json:"HRK"`
// 	HUF float32 `json:"HUF"`
// 	IDR float32 `json:"IDR"`
// 	ILS float32 `json:"ILS"`
// 	// TODO: ADD MORE
// 	USD float32 `json:"USD"`
// }

// https://api.frankfurter.app/latest?amount=10&from=gbp&to=usd

// Really messy code but works right now
func Currency(currAmount string, currFrom string, currTo string) (string, error) {
	baseUrl := "https://api.frankfurter.app"
	response, err := http.Get(fmt.Sprintf("%s/latest?amount=%s&from=%s&to=%s", baseUrl, currAmount, strings.ToUpper(currFrom), strings.ToUpper(currTo)))
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
	if responseObject.Rates[strings.ToUpper(currTo)] != 0 {
		return fmt.Sprintf("%s%s are %v%s", currAmount, currFrom, responseObject.Rates[strings.ToUpper(currTo)], currTo), nil
	}
	return "Something went wrong FeelsBadMan", nil
}
