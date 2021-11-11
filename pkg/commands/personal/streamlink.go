package personal

import "github.com/lyx0/nourybot/cmd/bot"

// Streamlink returns the streamlink config.
func Streamlink(target string, nb *bot.Bot) {
	reply := `https://haste.zneix.eu/udajirixep put this in ~/.config/streamlink/config on Linux (or %appdata%\streamlink\streamlinkrc on Windows)`

	nb.Send(target, reply)
}
