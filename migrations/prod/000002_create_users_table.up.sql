CREATE TABLE IF NOT EXISTS users (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	login text UNIQUE NOT NULL,
	twitchid text NOT NULL,
	level integer,
	location text,
	lastfm_username text
);

INSERT INTO users (added_at,login,twitchid,"level",location,lastfm_username) VALUES
	 (NOW(),'nouryxd','31437432',1000,'vilnius','nouryqt');
