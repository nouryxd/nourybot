package personal

import "github.com/lyx0/nourybot/cmd/bot"

// Farm responds with osrs wiki farming guide
func Farm(target string, nb *bot.Bot) {
	nb.Send(target, "Trees: https://oldschool.runescape.wiki/w/Crop_running#Example_tree_run_sequence")
	nb.Send(target, "Herbs: https://oldschool.runescape.wiki/w/Crop_running#Example_allotment,_flower_and_herb_run_sequence")
}
