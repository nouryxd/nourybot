CREATE TABLE IF NOT EXISTS commands_logs (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	twitch_login text NOT NULL,
	twitch_id text NOT NULL,
	twitch_channel text NOT NULL,
	twitch_message text NOT NULL,
	command_name text NOT NULL,
	user_level integer NOT NULL,
	identifier text NOT NULL,
	raw_message text NOT NULL
);

INSERT INTO commands_logs (added_at,twitch_login,twitch_id,twitch_channel,twitch_message,command_name,user_level,identifier,raw_message) VALUES
	 (NOW(),'nouryxd','31437432','nourybot','()weather Vilnius','weather',1000,'8441e97b-f622-4c42-b9b1-9bf22ba0d0bd','@badge-info=;badges=moderator/1,game-developer/1;color=#00F2FB;display-name=nouryxd;emotes=;first-msg=0;flags=;id=87d40f5c-8c7c-4105-9f57-b1a953bb42d0;mod=1;returning-chatter=0;room-id=596581605;subscriber=0;tmi-sent-ts=1696945359165;turbo=0;user-id=31437432;user-type=mod :nouryxd!nouryxd@nouryxd.tmi.twitch.tv PRIVMSG #nourybot :()weather Vilnius');

