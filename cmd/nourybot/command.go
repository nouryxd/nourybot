package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	"github.com/lyx0/nourybot/internal/data"
)

// AddCommand splits a message into two parts and passes on the
// name and text to the database handler.
func (app *application) AddCommand(name string, message twitch.PrivateMessage) {
	// snipLength is the length we need to "snip" off of the start of `message`.
	//  `()add command` = +12
	//  trailing space  =  +1
	//      zero-based  =  +1
	//                  =  15
	snipLength := 15

	// Split the twitch message at `snipLength` plus length of the name of the
	// command that we want to add.
	// The part of the message we are left over with is then passed on to the database
	// handlers as the `text` part of the command.
	//
	// e.g. `()addcommand CoolSponsors Check out CoolSponsor.com they are the coolest sponsors!
	//       | <- snipLength + name -> |  <--- command text with however many characters ---> |
	//       | <----- 14 + 12  ------> |
	text := message.Message[snipLength+len(name) : len(message.Message)]
	command := &data.Command{
		Name:        name,
		Channel:     message.Channel,
		Text:        text,
		Level:       0,
		Description: "",
	}
	app.Log.Info(command)
	err := app.Models.Commands.Insert(command)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply, message)
		return
	} else {
		reply := fmt.Sprintf("Successfully added command: %s", name)
		app.Send(message.Channel, reply, message)
		return
	}
}

// GetCommand queries the database for a name. If an entry exists it checks
// if the Command.Level is 0, if it is the command.Text value is returned.
//
// If the Command.Level is not 0 it queries the database for the level of the
// user who sent the message. If the users level is equal or higher
// the command.Text field is returned.
func (app *application) GetCommand(target, commandName string, userLevel int) (string, error) {
	app.Log.Infow("command",
		"target", target,
		"commandname", commandName,
	)
	// Fetch the command from the database if it exists.
	command, err := app.Models.Commands.Get(commandName, target)
	if err != nil {
		// It probably did not exist
		return "", err
	}

	app.Log.Info("command", command)

	if command.Level == 0 {
		return command.Text, nil
	} else if userLevel >= command.Level {
		// Userlevel is sufficient so return the command.Text
		return command.Text, nil

		// If the command has no level set just return the text.
		// Otherwise check if the level is high enough.
	}
	// Userlevel was not enough so return an empty string and error.
	return "", ErrUserInsufficientLevel
}

// GetCommand queries the database for a name. If an entry exists it checks
// if the Command.Level is 0, if it is the command.Text value is returned.
//
// If the Command.Level is not 0 it queries the database for the level of the
// user who sent the message. If the users level is equal or higher
// the command.Text field is returned.
func (app *application) GetCommandDescription(name, channel, username string) (string, error) {
	// Fetch the command from the database if it exists.
	command, err := app.Models.Commands.Get(name, channel)
	if err != nil {
		// It probably did not exist
		return "", err
	}

	// If the command has no level set just return the text.
	// Otherwise check if the level is high enough.
	if command.Level == 0 {
		return command.Description, nil
	} else {
		// Get the user from the database to check if the userlevel is equal
		// or higher than the command.Level.
		user, err := app.Models.Users.Get(username)
		if err != nil {
			return "", err
		}
		if user.Level >= command.Level {
			// Userlevel is sufficient so return the command.Text
			return command.Description, nil
		}
	}

	// Userlevel was not enough so return an empty string and error.
	return "", ErrUserInsufficientLevel
}

// EditCommandLevel takes in a name and level string and updates the entry with name
// to the supplied level value.
func (app *application) EditCommandLevel(name, lvl string, message twitch.PrivateMessage) {
	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Log.Error(err)
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrCommandLevelNotInteger), message)
		return
	}

	err = app.Models.Commands.SetLevel(name, message.Channel, level)

	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), message)
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated command %s to level %v", name, level)
		app.Send(message.Channel, reply, message)
		return
	}
}

// DebugCommand checks if a command with the provided name exists in the database
// and outputs information about it in the chat.
func (app *application) DebugCommand(name string, message twitch.PrivateMessage) {
	// Query the database for a command with the provided name
	cmd, err := app.Models.Commands.Get(name, message.Channel)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply, message)
		return
	} else {
		reply := fmt.Sprintf("id=%v\nname=%v\nchannel=%v\nlevel=%v\ntext=%v\ndescription=%v\n",
			cmd.ID,
			cmd.Name,
			cmd.Channel,
			cmd.Level,
			cmd.Text,
			cmd.Description,
		)

		//app.Send(message.Channel, reply)
		resp, err := app.uploadPaste(reply)
		if err != nil {
			app.Log.Errorln("Could not upload paste:", err)
			app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %v", ErrDuringPasteUpload), message)
			return
		}
		app.Send(message.Channel, resp, message)
		//app.SendEmail(fmt.Sprintf("DEBUG for command %s", name), reply)
		return
	}
}

// SetCommandHelp updates the `help` column of a given commands name in the
// database to the provided new help text.
func (app *application) EditCommandHelp(name string, message twitch.PrivateMessage) {
	// snipLength is the length we need to "snip" off of the start of `message`.
	// `()edit command` = +13
	//  trailing space =  +1
	//      zero-based =  +1
	//          `help` =  +4
	//                 =  19
	snipLength := 19

	// Split the twitch message at `snipLength` plus length of the name of the
	// command that we want to set the help text for so that we get the
	// actual help message left over and then assign this string to `text`.
	//
	// e.g. `()editcommand help FeelsDankMan Returns a FeelsDankMan ascii art. Requires user level 500.`
	//       | <---- snipLength + name ----> | <------  help text with however many characters. ----> |
	//       | <--------- 19 + 12  --------> |
	text := message.Message[snipLength+len(name) : len(message.Message)]
	err := app.Models.Commands.SetDescription(name, message.Channel, text)

	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), message)
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated help text for command %s to: %v", name, text)
		app.Send(message.Channel, reply, message)
		return
	}
}

// DeleteCommand takes in a name value and deletes the command from the database if it exists.
func (app *application) DeleteCommand(name string, message twitch.PrivateMessage) {
	err := app.Models.Commands.Delete(name, message.Channel)
	if err != nil {
		app.Send(message.Channel, "Something went wrong FeelsBadMan", message)
		app.Log.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted command %s", name)
	app.Send(message.Channel, reply, message)
}

func (app *application) LogCommand(msg twitch.PrivateMessage, commandName string, userLevel int) {
	twitchLogin := msg.User.Name
	twitchID := msg.User.ID
	twitchMessage := msg.Message
	twitchChannel := msg.Channel
	identifier := uuid.NewString()
	rawMsg := msg.Raw

	go app.Models.CommandsLogs.Insert(twitchLogin, twitchID, twitchChannel, twitchMessage, commandName, userLevel, identifier, rawMsg)
}

// InitialTimers is called on startup and queries the database for a list of
// timers and then adds each onto the scheduler.
func (app *application) ListCommands() string {
	command, err := app.Models.Commands.GetAll()
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var cs []string

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range command {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i

		// cronName is the internal, unique tag/name for the timer.
		// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"
		c := fmt.Sprintf(
			"ID: \t\t%v\n"+
				"Name: \t\t%v\n"+
				"Channel: \t%v\n"+
				"Text: \t\t%v\n"+
				"Level: \t\t%v\n"+
				"Description: %v\n"+
				"\n\n",
			v.ID, v.Name, v.Channel, v.Text, v.Level, v.Description,
		)

		// Add new value to the slice
		cs = append(cs, c)

		//app.Scheduler.AddFunc(repeating, func() { app.newPrivateMessageTimer(v.Channel, v.Text) }, cronName)
	}

	reply, err := app.uploadPaste(strings.Join(cs, ""))
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}

	return reply
}

// InitialTimers is called on startup and queries the database for a list of
// timers and then adds each onto the scheduler.
func (app *application) ListChannelCommands(channel string) string {
	channelUrl := fmt.Sprintf("https://bot.noury.is/commands/%s", channel)
	commandUrl := "https://bot.noury.is/commands"
	command, err := app.Models.Commands.GetAllChannel(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var cs []string

	allHelpText := app.GetAllHelpText()
	app.Log.Info(allHelpText)
	cs = append(cs, fmt.Sprintf("General commands: \n\n%s\nChannel commands:\n\n", allHelpText))

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range command {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c string

		c = fmt.Sprintf(
			"Name: \t%v\n"+
				"Description: %v\n"+
				"Level: \t%v\n"+
				"Text: \t%v\n"+
				"\n",
			v.Name, v.Description, v.Level, v.Text,
		)

		// Add new value to the slice
		cs = append(cs, c)

	}

	reply := fmt.Sprintf("Channel commands: %s | General commands: %s", channelUrl, commandUrl)
	return reply
}
