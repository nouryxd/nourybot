package main

func (nb *nourybot) connect() error {
	nb.twitchClient.Join("nourybot")

	err := nb.twitchClient.Connect()
	if err != nil {
		panic(err)
	}

	nb.twitchClient.Say("nourybot", "xd")
	return nil
}

func (nb *nourybot) onConnect() {
	nb.twitchClient.Say("nourybot", "xd")
}
