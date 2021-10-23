nourybot commands

##### 8ball:
	Ask the magic 8ball for guidance

	
#### Bot/Botinfo/Nourybot/Help: 
	Returns information about the bot
  
#### Botstatus:
	Returns the status of a verified/known/normal bots and its ratelimits
	API used: https://customapi.aidenwallis.co.uk/docs/twitch/bot-status
  
#### Bttv:
	Returns a bttv search link for a given emote name
	
#### Bttvemotes:
	Returns all Bttv emotes from the current channel
	API used: https://customapi.aidenwallis.co.uk/docs/emotes/bttv

#### Commands:
	Returns this page 4Head
  
#### Coinflip: 
	Head or Tails
  
#### Ffz:
  	Return sa ffz search link for a given emote name
      Usage: ()ffz [emotename]
	
#### Ffzemotes:
	Returns all Ffz emotes from the current channel
	API used: https://customapi.aidenwallis.co.uk/docs/emotes/ffz

#### Fill:
	Repeats an emote for the maximum size of a single message and "fills" the message with it. 
	Only usable for vips/moderators/broadcaster.

#### Firstline/fl:
	Returns the first line a user has written in a channel.
	Usage: ()fl [channel] [user]
	API used: https://api.ivr.fi/logs/firstmessage/

#### Followage:
	Returns the date a user has been following a channel.
	Usage: ()followage [channel] [user]
	API used: https://api.ivr.fi/twitch/subage/
	
#### Game:
	Returns the game of the current channel if no parameters are given, otherwise for the given channel
    Usage: ()game [channel]

#### Number/num:
	Returns a fact about a given number, or if none is given a random number.
    API used: http://numbersapi.com/#42
	
#### Godoc/Godocs:
	Returns the godocs.io search for a given word
	
#### Ping:
	Returns a Pong

#### Profilepicture/pfp:
	Returns a link to a users profile picture.
	API used: https://api.ivr.fi/twitch/resolve/
	
#### Pingme:
	Pings you in chat

#### Pyramid:
	Builds a pyramid of a given emote, only usable by vips/moderators/broadcaster in the channel. Max size 20

#### RandomCat/cat: 
	Returns a random cat image
	API used: https://aws.random.cat/meow
	
#### RandomCat/dog: 
	Returns a random dog image
	API used: https://random.dog/woof.json
	
#### Randomduck/duck: 
	Returns a random duck image
	API used: https://random-d.uk/

#### RandomFox/fox: 
	Returns a random fox image
	API used: https://randomfox.ca/floof
	
#### Randomxkcd/rxkcd:
	Returns a random Xkcd comic

#### Randomquote/rq:
	Returns a random quote from a user in a channel.
	Usage: ()rq [channel] [user]
	API used: https://api.ivr.fi/logs/rq/

### Robohash/robo:
    Returns a link to the robohash image of your Twitch Message Id.
    API Used: https://robohash.org
	
#### Subage:
	Returns the months someone has been subscribed to a channel.
    Usage: ()subage [channel] [user]
	API used: https://api.ivr.fi/twitch/subage/

### Thumb/Preview:
    Returns a screenshot of a given live channel.
    Usage: ()thumb [channel]

#### Title:
	Returns the title of the current channel if no parameters are given, otherwise for the given channel

#### Mycolor/color:
	Returns the hexcode of your Twitch color
	
#### Uptime:
	Returns the uptime of the current channel if no parameters are given, otherwise for the given channel
    Usage: ()uptime [channel]

#### userid/uid:
 	Returns the Twitch User ID of a given user, otherwise the senders userid.
    Usage: ()uid [name]
	
#### Uid:
	Returns the Twitch userid of a given username
	API used: https://api.ivr.fi/twitch/resolve/

#### Weather:
	Returns the weather for a given location
	API used: https://customapi.aidenwallis.co.uk/api/v1/misc/weather/
    Usage: ()weather [location]
	
#### Xkcd:
	Returns a link to the current Xkcd comic

