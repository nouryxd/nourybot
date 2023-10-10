CREATE TABLE IF NOT EXISTS sent_messages_logs (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	twitch_channel text NOT NULL,
	twitch_message text NOT NULL,
	identifier text NOT NULL
);

INSERT INTO sent_messages_logs (added_at,twitch_channel,twitch_message,identifier) VALUES
	 (NOW(),'nourybot','Weather for Vilnius, LT: Feels like: 9.3°C. Currently 10.85°C with a high of 12.07°C and a low of 10.49°C, humidity: 50%, wind: 2.57m/s.','04fbd9c0-47da-466f-b966-44d1d04de11c');
