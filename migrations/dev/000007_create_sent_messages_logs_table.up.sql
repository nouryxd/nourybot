CREATE TABLE IF NOT EXISTS sent_messages_logs (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	twitch_channel text NOT NULL,
	twitch_message text NOT NULL,
	context_command_name text,
	context_username text,
	context_message text,
	context_user_id text,
	identifier text,
	context_raw text
);

INSERT INTO sent_messages_logs (added_at,twitch_channel,twitch_message,context_command_name,context_username,context_message,context_user_id,identifier,context_raw) VALUES
	 (NOW(),'nourybot','Weather for Vilnius, LT: Feels like: 8.07°C. Currently 8.65°C with a high of 9.29°C and a low of 8.49°C, humidity: 66%, wind: 1.54m/s.','weather','noury','()weather Vilnius','31437432','654f9761-b2d4-4975-a4fd-84c6ec7f2eb8','@badge-info=;badges=moderator/1,game-developer/1;color=#00F2FB;display-name=noury;emotes=;first-msg=0;flags=;id=357d94a4-024e-49ea-ab3d-d97286cd0492;mod=1;returning-chatter=0;room-id=596581605;subscriber=0;tmi-sent-ts=1696952295788;turbo=0;user-id=31437432;user-type=mod :nouryxd!nouryxd@nouryxd.tmi.twitch.tv PRIVMSG #nourybot :()weather Vilnius');
