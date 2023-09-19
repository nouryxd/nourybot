package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v4"
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
	snipLength := 14

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
		Name:     name,
		Text:     text,
		Category: "uncategorized",
		Level:    0,
		Help:     "",
	}
	app.Log.Info(command)
	err := app.Models.Commands.Insert(command)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply)
		return
	} else {
		reply := fmt.Sprintf("Successfully added command: %s", name)
		app.Send(message.Channel, reply)
		return
	}
}

// GetCommand queries the database for a name. If an entry exists it checks
// if the Command.Level is 0, if it is the command.Text value is returned.
//
// If the Command.Level is not 0 it queries the database for the level of the
// user who sent the message. If the users level is equal or higher
// the command.Text field is returned.
func (app *application) GetCommand(name, username string) (string, error) {
	// Fetch the command from the database if it exists.
	command, err := app.Models.Commands.Get(name)
	if err != nil {
		// It probably did not exist
		return "", err
	}

	// If the command has no level set just return the text.
	// Otherwise check if the level is high enough.
	if command.Level == 0 {
		return command.Text, nil
	} else {
		// Get the user from the database to check if the userlevel is equal
		// or higher than the command.Level.
		user, err := app.Models.Users.Get(username)
		if err != nil {
			return "", err
		}
		if user.Level >= command.Level {
			// Userlevel is sufficient so return the command.Text
			return command.Text, nil
		}
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
func (app *application) GetCommandHelp(name, username string) (string, error) {
	// Fetch the command from the database if it exists.
	command, err := app.Models.Commands.Get(name)
	if err != nil {
		// It probably did not exist
		return "", err
	}

	// If the command has no level set just return the text.
	// Otherwise check if the level is high enough.
	if command.Level == 0 {
		return command.Help, nil
	} else {
		// Get the user from the database to check if the userlevel is equal
		// or higher than the command.Level.
		user, err := app.Models.Users.Get(username)
		if err != nil {
			return "", err
		}
		if user.Level >= command.Level {
			// Userlevel is sufficient so return the command.Text
			return command.Help, nil
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
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrCommandLevelNotInteger))
		return
	}

	err = app.Models.Commands.SetLevel(name, level)

	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound))
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated command %s to level %v", name, level)
		app.Send(message.Channel, reply)
		return
	}
}

// EditCommandCategory takes in a name and category string and updates the command
// in the databse with the passed in new category.
func (app *application) EditCommandCategory(name, category string, message twitch.PrivateMessage) {
	err := app.Models.Commands.SetCategory(name, category)

	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound))
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated command %s to category %v", name, category)
		app.Send(message.Channel, reply)
		return
	}
}

// DebugCommand checks if a command with the provided name exists in the database
// and outputs information about it in the chat.
func (app *application) DebugCommand(name string, message twitch.PrivateMessage) {
	// Query the database for a command with the provided name
	cmd, err := app.Models.Commands.Get(name)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply)
		return
	} else {
		reply := fmt.Sprintf("id=%v\nname=%v\nlevel=%v\ncategory=%v\ntext=%v\nhelp=%v\n",
			cmd.ID,
			cmd.Name,
			cmd.Level,
			cmd.Category,
			cmd.Text,
			cmd.Help,
		)

		//app.Send(message.Channel, reply)
		resp, err := app.uploadPaste(reply)
		if err != nil {
			app.Log.Errorln("Could not upload paste:", err)
			app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %v", ErrDuringPasteUpload))
			return
		}
		app.Send(message.Channel, resp)
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
	err := app.Models.Commands.SetHelp(name, text)

	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound))
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated help text for command %s to: %v", name, text)
		app.Send(message.Channel, reply)
		return
	}
}

// DeleteCommand takes in a name value and deletes the command from the database if it exists.
func (app *application) DeleteCommand(name string, message twitch.PrivateMessage) {
	err := app.Models.Commands.Delete(name)
	if err != nil {
		app.Send(message.Channel, "Something went wrong FeelsBadMan")
		app.Log.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted command %s", name)
	app.Send(message.Channel, reply)
}
