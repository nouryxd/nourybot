package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// https://api.ivr.fi
type whoisApiResponse struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Login       string `json:"login"`
	Bio         string `json:"bio"`
	ChatColor   string `json:"chatColor"`
	Partner     bool   `json:"partner"`
	// Affiliate   bool   `json:"affiliate"`
	Bot       bool   `json:"bot"`
	CreatedAt string `json:"createdAt"`
	Roles     roles
	Error     string `json:"error"`
}

type roles struct {
	IsAffiliate bool `json:"isAffiliate"`
	IsPartner   bool `json:"isPartner"`
	IsSiteAdmin bool `json:"isSiteAdmin"`
	IsStaff     bool `json:"isStaff"`
}

// Userid returns the userID of a given user
func Whois(username string) string {
	baseUrl := "https://api.ivr.fi/twitch/resolve"

	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, username))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject whoisApiResponse
	json.Unmarshal(body, &responseObject)

	reply := fmt.Sprintf("User: %s, ID: %s, Color: %s, Partner: %v, Affiliate: %v, Staff: %v, Admin: %v, Bot: %v, Bio: %v", responseObject.DisplayName, responseObject.Id, responseObject.ChatColor, responseObject.Roles.IsPartner, responseObject.Roles.IsAffiliate, responseObject.Roles.IsStaff, responseObject.Roles.IsSiteAdmin, responseObject.Bot, responseObject.Bio)

	// User not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan")
	} else {
		return reply
	}
}
