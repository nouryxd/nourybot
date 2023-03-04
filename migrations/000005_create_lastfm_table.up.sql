CREATE TABLE IF NOT EXISTS lastfm_users (
	id bigserial PRIMARY KEY,
	twitch_login text NOT NULL,
	twitch_id text UNIQUE NOT NULL,
	lastfm_username text NOT NULL 
);

INSERT INTO lastfm_users (twitch_login,twitch_id,lastfm_username) VALUES
	 ('nourylul','31437432','nouryqt');

