package personal

import "github.com/lyx0/nourybot/cmd/bot"

func Rave(target string, nb *bot.Bot) {
	reply := `https://www.youtube.com/playlist?list=PLY9LTYa8xnQKrug3VvgkPWqmpmXSKAYPe`
	reply2 := `https://www.youtube.com/playlist?list=PLWwCeXamjNuZ2ZNJiNwvdmVT9TtULsHc6`

	nb.Send(target, reply)
	nb.Send(target, reply2)

}
