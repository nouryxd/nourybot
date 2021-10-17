package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func RandomNumber() string {
	response, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(responseData)
}

func Number(num string) string {
	number, err := strconv.Atoi(num)
	if err != nil {
		return "Given value is not a number."
	} else {
		response, err := http.Get(fmt.Sprint("http://numbersapi.com/" + string(number)))
		if err != nil {
			log.Fatalln(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		return string(responseData)
	}

}
