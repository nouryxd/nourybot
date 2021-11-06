package personal

import "github.com/lyx0/nourybot/cmd/bot"

func Arch(target string, nb *bot.Bot) {
	reply := "Your friend isn't wrong. Being on the actual latest up to date software, having a single unified community repository for out of repo software (AUR) instead of a bunch of scattered broken PPAs for extra software, not having so many hard dependencies that removing GNOME removes basic system utilities, broader customization support and other things is indeed, pretty nice."

	nb.Send(target, reply)
}

func ArchTwo(target string, nb *bot.Bot) {
	reply := "One time I was ordering coffee and suddenly realised the barista didn't know I use Arch. Needless to say, I stopped mid-order to inform her that I do indeed use Arch. I must have spoken louder than I intended because the whole café instantly erupted into a prolonged applause. I walked outside with my head held high. I never did finish my order that day, but just knowing that everyone around me was aware that I use Arch was more energising than a simple cup of coffee could ever be."

	nb.Send(target, reply)
}
