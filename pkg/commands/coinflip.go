package commands

import (
	"github.com/nouryxd/nourybot/pkg/common"
)

// Coinflip returns either "Heads!" or "Tails!"
func Coinflip() string {
	flip := common.GenerateRandomNumber(2)
	var reply string

	switch flip {
	case 0:
		reply = "Heads!"
	case 1:
		reply = "Tails!"
	default:
		reply = "Heads!"
	}

	return reply
}
